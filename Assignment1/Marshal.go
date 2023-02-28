//package
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//)
//
//type Book struct {
//	Title  string `json:"title"`
//	Author Author `json:"author"`
//}
//
//type Author struct {
//	Sales     int  `json:"book_sales"`
//	Age       int  `json:"age"`
//	Developer bool `json:"is_developer"`
//}
//type SensorReading struct {
//	Name     string `json:"name"`
//	Capacity int    `json:"capacity"`
//	Time     string `json:"time"`
//}
//
//func main() {
//
//	//jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
//
//	//var reading SensorReading
//	//err := json.Unmarshal([]byte(jsonString), &reading)
//	//fmt.Printf("%+v\n", reading)
//
//	author := Author{Sales: 3, Age: 25, Developer: true}
//	book := Book{Title: "Cacl2", Author: author}
//	//
//	byteArray, err := json.MarshalIndent(book, "", " ")
//	if err != nil {
//		fmt.Println("sdfs")
//	}
//	err_ := ioutil.WriteFile("test.json", byteArray, 0644)
//	if err_ != nil {
//		return
//	}
//	//fmt.Println(string(byteArray))
//}
