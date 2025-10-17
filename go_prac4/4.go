package main

import "fmt"

func main() {
	total := 0

	for i := 0; i < 5; i++ { // 반복문 조건식 안에서 i와 같은 변수를 선언 할때는 자동변수를 사용하자.
		fmt.Println(i + 1)
	}

	fmt.Println("***************************")

	for i := 0; i < 5; i++ {
		total = total + i

	}
	fmt.Printf("총합계는 : %d\n", total) //Println은 %d 이런것을 해석하는 기능이 없다 따라서 total 변수 값을 출력하고 싶다면 Println("총합계는 : ",total) 이렇게 써야한다
}
