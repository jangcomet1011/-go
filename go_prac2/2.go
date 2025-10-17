package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	/*var length float64 = 3.2
	var width int = 2
	fmt.Println("면적은", int(length)*width)
	fmt.Println("length > width?", int(length) > width)
	fmt.Println("형변환", reflect.TypeOf(int(length)))
	fmt.Println("원본", reflect.TypeOf(length)) */

	var now time.Time = time.Now()     //2025년 10월 17일 21시 13분 43초 같은 모든 정보 time.Time은 time.Now() 가 반환하는 자료형
	var month time.Month = now.Month() // month := now.Month()  time.Month는 .Month()가 반환하는 자료형
	var day int = now.Day()            // now변수에서 '일'만 꺼내는 함수
	fmt.Println(month, day)

	//var univ string = "Go$ Inha$"
	changer := strings.NewReplacer("$", "!") //strings.NewReplacer() 함수는 어떤 글자를 어떤 글자로 바꿀지에 대한 규칙을 만들어주는 역할을 한다 "$" 는 찾을 문자열 "!" 는 바꿀 문자열을 의미 한다
	//changed1 := changer.Replace(univ)
	changed := changer.Replace("Go$ Inha$")
	fmt.Println(changed)
	//fmt.Println(changed1)

	r := bufio.NewReader(os.Stdin) //var r *bufio.Reader = bufio.NewReader(os.Stdin) 수동선언
	i, _ := r.ReadString('\n')     // 사용자가 줄바꿈 즉 엔터키를 누를때 까지 모든 문자열을 읽어 들인다 ReadString함수는 두가지의 값을 반환합니다
	// 한개는 i에 _는 값이 필요 없으니 무시하겠다는 뜻 즉 오류를 무시하겠다는 뜻이다
	fmt.Println(i)

	i, err := r.ReadString('\n') // i 에는 성공적으로 읽은 문자열이 저장된다. err는 만약 입력을 받는중에 오류가 났다면 오류(에러) 정보가 err에 저장된다
	// fmt.Println(err)
	if err != nil { // nil은 값이 없음을 뜻함
		log.Fatal(err) // report the error and exit the program 에러메시지 출력 후 프로그램 종료
	}
	fmt.Println(i)

	/* fmt.Print("점수를 입력하세요: ") 학점 판정 프로그램

	// 1. 키보드로부터 한 줄을 입력받음
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // 엔터 키를 누를 때까지 읽음
	if err != nil {
		log.Fatal(err) // 입력 에러가 나면 프로그램 중단
	}

	// 2. 입력받은 문자열에서 앞뒤 공백 및 줄바꿈 문자 제거
	input = strings.TrimSpace(input)

	// 3. 깨끗해진 문자열을 정수(int) 타입으로 변환
	score, err := strconv.Atoi(input)
	if err != nil {
		// 숫자로 변환할 수 없는 값(예: "hello")이 입력되면 에러 처리
		log.Fatal("숫자만 입력해주세요!")
	}

	// 4. 변환된 점수로 학점 계산 (이전 코드와 동일)
	var grade string
	if score > 100 || score < 0 {
		grade = "잘못된 점수"
	} else if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("입력한 점수 %d는 %s 학점입니다.\n", score, grade) */
}
