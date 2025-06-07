/*
reflect package allow inspect and manipulate types and values at runtime

| Purpose                       | Function / Type                  |
| ----------------------------- | -------------------------------- |
| Get the type of a value       | `reflect.TypeOf()`               |
| Get the value of a variable   | `reflect.ValueOf()`              |
| Access struct fields          | `Type.Field()`, `Value.Field()`  |
| Read struct tags (e.g., JSON) | `StructTag.Get()`                |
| Modify values at runtime      | `Value.Set()` (requires pointer) |
*/

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name"` // tag để đặt tên trường JSON
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // omit empty nếu rỗng sẽ không xuất hiện
}

func main() {
	var x = 10
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Type:", t)        // int
	fmt.Println("Kind:", t.Kind()) // int
	fmt.Println("Value:", v.Int()) // 10

	user := User{Name: "Alice", Age: 30}
	user_type := reflect.TypeOf(user)

	fmt.Println("\n")
	for i := 0; i < user_type.NumField(); i++ {
		field := user_type.Field(i)
		fmt.Println("Field:", field.Name, "- Tag:", field.Tag.Get("json"))
	}
}
