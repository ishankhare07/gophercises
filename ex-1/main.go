package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Quiz struct {
	Results   map[bool]int
	Questions []*Question
}

func NewQuiz() *Quiz {
	return &Quiz{Results: make(map[bool]int)}
}

func (q *Quiz) InsertQuestion(ques *Question) {
	q.Questions = append(q.Questions, ques)
}

func (q *Quiz) RegisterResponse(ques *Question, ans string) {
	if ok := ques.matchAnswer(ans); ok {
		q.Results[true] += 1
	} else {
		q.Results[false] += 1
	}
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

	for _, q := range quiz.Questions {
		var ans string
		fmt.Printf("What is %s? :", q.q)
		fmt.Scanf("%s", &ans)
		quiz.RegisterResponse(q, ans)
	}

	fmt.Printf("You answered %d correct out of %d\n", quiz.Results[true], quiz.Results[true]+quiz.Results[false])
}
