package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a score: ")
	r := bufio.NewReader(os.Stdin)
	// i, _ := r.ReadString('\n') // ignore error
	i, err := r.ReadString('\n')
	// fmt.Println(err)
	if err != nil {
		log.Fatal(err) // report the error and exit the program
		log.Fatal(err)
	}
	fmt.Println(i)

	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공란 및 탭 키 등 제거
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	if score >= 60 {
		fmt.Println(score, "Pass")
	} else {
		fmt.Println(score, "Fail")
	}

	// shadowing

	// var float64 float64 = 2.71
	// var f float64 = 3.991
	// fmt.Println(float64)
	// fmt.Println(f)

	// var fmt float64 = 2.71
	// fmt.Println(fmt)
}
