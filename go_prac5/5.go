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
	fmt.Println("학점을 입력하세요")

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i2 := strings.TrimSpace(i)
	score, err := strconv.ParseInt(i2, 10, 32)

	var grade string
	if score > 100 || score < 0 {
		fmt.Println("잘못된 입력입니다")
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
	fmt.Printf("점수는 %d 이고 학점은 %s 입니다", score, grade)
}
