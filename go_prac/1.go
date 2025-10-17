package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var f64 float64
	var t bool
	var s string
	var i int
	var i16 int16
	name := "jellyam"
	age := 32
	fmt.Println(math.Round(2.21))               //소수를 가장 가까운 정수로 반올림함
	fmt.Println(math.Ceil(2.21))                //소수를 무조건 올림함
	fmt.Println(strings.Title("go developer!")) //문자열내 첫글자를 대문자로 만든다
	fmt.Println("Kim\nInha\t\"\\")
	fmt.Println('2', '가') //go 언어에서 작은따움표'' 로 감싸는건 문자열이 아닌 유니코드 포인트를 의미하는 룬 타입이다.
	fmt.Println("********************************************")
	fmt.Println("이름은 :", name, ",", reflect.TypeOf(name), ", 나이는 :", age, ",", reflect.TypeOf(age))
	fmt.Println("********************************************")
	fmt.Println(reflect.TypeOf(2.31))       //float64
	fmt.Println(reflect.TypeOf("Kim Inha")) //string
	fmt.Println(reflect.TypeOf(true))       //bool
	fmt.Println(reflect.TypeOf('A'))        //원래라면 string이 맞지만 '' 두개쓰면 유니코드를 나타내는것 유니코드는 int32규격을 사용용
	fmt.Println(reflect.TypeOf(19))         //int
	fmt.Println("********************************************")
	fmt.Println(f64, reflect.TypeOf(f64)) //변수 선언하고 아무것도 안넣어서 0 float 64
	fmt.Println(t, reflect.TypeOf(t))     // 변수 선언하고 아무것도 안넣어서 0인 false bool
	fmt.Println(s, reflect.TypeOf(s))     //변수 선언하고 아무것도 안넣어서   string
	fmt.Println(i, reflect.TypeOf(i))     // 변수 선언하고 아무것도 안넣어서 0 int
	fmt.Println(i16, reflect.TypeOf(i16)) //변수 선언하고 아무것도 안넣어서 0 int16
	// var id int16
	// var name string
	// var gpa float32

	// id = 999
	// name = "Kim Inha"
	// gpa = 3.99

	// var id int16 = 999
	// var name string = "Kim Inha"
	// var gpa float32 = 3.99

	// var id = 999
	// var name = "Kim Inha"
	// var gpa = 3.99

	// id := 999
	// name := "Kim Inha"
	// gpa := 3.99
}
