package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Quiz struct {
	Results   map[int]bool
	Questions []*Question
}

func NewQuiz() *Quiz {
	return &Quiz{Results: make(map[int]bool)}
}

func (q *Quiz) InsertQuestion(ques *Question) {
	q.Questions = append(q.Questions, ques)
}

type Question struct {
	q string
	a string
}

func NewQuestion(r []string) *Question {
	return &Question{q: r[0], a: r[1]}
}

func (q *Question) String() string {
	return fmt.Sprintf("<%T %s: %s>", *q, q.q, q.a)
}

func (q *Question) matchAnswer(answer string) bool {
	if q.a == answer {
		return true
	} else {
		return false
	}
}

func main() {
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("cannot open file. ", err)
	}
	// file.Read(b)

	// if err != nil {
	// 	fmt.Println("err: ", err)
	// }

	quiz := NewQuiz()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("err: ", err)
		}
		quiz.InsertQuestion(NewQuestion(record))

	}

	fmt.Println(quiz)
}
