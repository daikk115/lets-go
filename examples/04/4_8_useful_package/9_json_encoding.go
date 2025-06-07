package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"` // tag để đặt tên trường JSON
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // omit empty nếu rỗng sẽ không xuất hiện
}

func encodeJson() {
	fmt.Println("\n------------------ encodeJson ------------------")

	a := User{Name: "Alice", Age: 30}
	b := User{Name: "Alice", Age: 30, Email: "daidv@example.vn"}

	jsonData, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonData2, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
	fmt.Println(string(jsonData2))
}

func decodeJson() {
	fmt.Println("\n------------------ decodeJson ------------------")
	rawData := `{"name":"Alice","age":30}`
	rawData2 := `{"name":"Alice","age":30,"email":"daidv@example.vn"}`
	var u, u2 User
	err := json.Unmarshal([]byte(rawData), &u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = json.Unmarshal([]byte(rawData2), &u2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", u)
	fmt.Printf("%+v\n", u2)
}

func main() {
	encodeJson()
	decodeJson()
}
