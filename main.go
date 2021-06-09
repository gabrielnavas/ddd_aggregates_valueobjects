package main

import (
	"ddd_aggregates_valueobjects/domain"
	"fmt"
)

func main() {
	tutor := domain.Tutor{
		Id:   "456",
		Name: "John Lennon",
	}

	course := domain.Course{
		Id:      "123",
		Name:    "Learning about DDD",
		TutorId: tutor.Id,
	}

	reference := domain.NewReference("Mary", "The Ultimate DDD", 400)

	video1 := domain.Video{
		Id:        "11",
		Course:    course,
		URL:       "http://course.ddd.com.br",
		Reference: reference,
	}

	fmt.Println(video1)
	fmt.Println(video1.Reference.ToString())
}
