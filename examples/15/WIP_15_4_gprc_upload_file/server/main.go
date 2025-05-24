package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "test/proto"

	"google.golang.org/grpc"
)

const (
	uploadDir    = "./uploads"
	maxChunkSize = 1024 * 1024 // 1MB chunks
)

type server struct{}

// SimpleUpload handles small file uploads in a single request
func (s *server) SimpleUpload(stream grpc.ServerStream, req *pb.SimpleUploadRequest) (*pb.SimpleUploadResponse, error) {
	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return &pb.SimpleUploadResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to create upload directory: %v", err),
		}, nil
	}

	// Create the file
	filePath := filepath.Join(uploadDir, req.Filename)
	file, err := os.Create(filePath)
	if err != nil {
		return &pb.SimpleUploadResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to create file: %v", err),
		}, nil
	}
	defer file.Close()

	// Write content to file
	if _, err := file.Write(req.Content); err != nil {
		return &pb.SimpleUploadResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to write file: %v", err),
		}, nil
	}

	log.Printf("Successfully uploaded file: %s (%d bytes)", req.Filename, len(req.Content))

	return &pb.SimpleUploadResponse{
		Success: true,
		Message: "File uploaded successfully",
		Size:    int64(len(req.Content)),
	}, nil
}

// ChunkUpload handles large file uploads using client streaming
func (s *server) ChunkUpload(stream pb.FileUploadService_ChunkUploadServer) error {
	var file *os.File
	var filename string
	var totalSize int64
	var uploadedSize int64

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return stream.SendAndClose(&pb.ChunkUploadResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to create upload directory: %v", err),
		})
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Client finished sending
			if file != nil {
				file.Close()
				log.Printf("Successfully uploaded file: %s (%d bytes)", filename, uploadedSize)
				return stream.SendAndClose(&pb.ChunkUploadResponse{
					Success:      true,
					Message:      "File uploaded successfully",
					UploadedSize: uploadedSize,
				})
			}
			return stream.SendAndClose(&pb.ChunkUploadResponse{
				Success: false,
				Message: "No file data received",
			})
		}
		if err != nil {
			if file != nil {
				file.Close()
				os.Remove(filepath.Join(uploadDir, filename)) // Clean up partial file
			}
			return err
		}

		switch data := req.Data.(type) {
		case *pb.ChunkUploadRequest_Info:
			// First message with file info
			filename = data.Info.Filename
			totalSize = data.Info.TotalSize

			filePath := filepath.Join(uploadDir, filename)
			file, err = os.Create(filePath)
			if err != nil {
				return stream.SendAndClose(&pb.ChunkUploadResponse{
					Success: false,
					Message: fmt.Sprintf("Failed to create file: %v", err),
				})
			}
			log.Printf("Starting upload: %s (%d bytes)", filename, totalSize)

		case *pb.ChunkUploadRequest_Chunk:
			// File chunk data
			if file == nil {
				return stream.SendAndClose(&pb.ChunkUploadResponse{
					Success: false,
					Message: "File info not received first",
				})
			}

			n, err := file.Write(data.Chunk)
			if err != nil {
				file.Close()
				os.Remove(filepath.Join(uploadDir, filename))
				return stream.SendAndClose(&pb.ChunkUploadResponse{
					Success: false,
					Message: fmt.Sprintf("Failed to write chunk: %v", err),
				})
			}
			uploadedSize += int64(n)
		}
	}
}

// StreamUpload handles bidirectional streaming with progress updates
func (s *server) StreamUpload(stream pb.FileUploadService_StreamUploadServer) error {
	var file *os.File
	var filename string
	var totalSize int64
	var uploadedSize int64

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return stream.Send(&pb.StreamUploadResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to create upload directory: %v", err),
		})
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if file != nil {
				file.Close()
				log.Printf("Successfully uploaded file: %s (%d bytes)", filename, uploadedSize)
			}
			return nil
		}
		if err != nil {
			if file != nil {
				file.Close()
				os.Remove(filepath.Join(uploadDir, filename))
			}
			return err
		}

		switch data := req.Data.(type) {
		case *pb.StreamUploadRequest_Info:
			filename = data.Info.Filename
			totalSize = data.Info.TotalSize

			filePath := filepath.Join(uploadDir, filename)
			file, err = os.Create(filePath)
			if err != nil {
				return stream.Send(&pb.StreamUploadResponse{
					Success: false,
					Message: fmt.Sprintf("Failed to create file: %v", err),
				})
			}
			log.Printf("Starting streaming upload: %s (%d bytes)", filename, totalSize)

			// Send acknowledgment
			stream.Send(&pb.StreamUploadResponse{
				Success: true,
				Message: "Ready to receive file chunks",
			})

		case *pb.StreamUploadRequest_Chunk:
			if file == nil {
				return stream.Send(&pb.StreamUploadResponse{
					Success: false,
					Message: "File info not received first",
				})
			}

			n, err := file.Write(data.Chunk)
			if err != nil {
				file.Close()
				os.Remove(filepath.Join(uploadDir, filename))
				return stream.Send(&pb.StreamUploadResponse{
					Success: false,
					Message: fmt.Sprintf("Failed to write chunk: %v", err),
				})
			}
			uploadedSize += int64(n)

			// Send progress update
			progress := float32(uploadedSize) / float32(totalSize) * 100
			stream.Send(&pb.StreamUploadResponse{
				Success:         true,
				Message:         "Chunk received",
				UploadedSize:    uploadedSize,
				ProgressPercent: progress,
			})
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024 * 1024 * 50), // 50MB max message size
	)

	pb.RegisterFileUploadServiceServer(s, &server{})

	log.Printf("File upload server listening on port 50051...")
	log.Printf("Upload directory: %s", uploadDir)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
