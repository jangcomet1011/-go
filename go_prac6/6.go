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
	fmt.Println("점수를 입력하세요")

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i2 := strings.TrimSpace(i)
	score, err := strconv.ParseInt(i2, 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	var grade string
	if score > 100 || score < 0 {
		log.Fatal(score)
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
	fmt.Printf("당신의 점수는:%d, 당신의 학점은:%s", score, grade)
}
