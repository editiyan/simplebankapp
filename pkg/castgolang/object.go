package castgolang

import (
	"fmt"
)

// func test() {
// 	// var data map[string]interface{}
// 	var data map[string]any

// 	jsonStr := `{"name": "Jojo", "age":17, "verified": true}`
// 	json.Unmarshal([]byte(jsonStr), &data)

// 	fmt.Println(data["name"], data["age"], data["verified"])

// 	//Type assertion
// 	var x interface{} = "hello"

// 	str, ok := x.(string)
// 	if ok {
// 		fmt.Println("Ini string:", str)
// 	} else {
// 		fmt.Println("Ini bukan string")
// 	}
// }

// Type switch
func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float64:", v)
	}
}
