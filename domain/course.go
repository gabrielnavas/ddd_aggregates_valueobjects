package domain

import "fmt"

/*
	Course struct is aggregate root
	Reference struct is value object
*/
type Course struct {
	Id      string
	Name    string
	TutorId string
}

type Video struct {
	Id        string
	Course    Course
	URL       string
	Reference Reference
}

type Reference struct {
	author   string
	bookName string
	page     int
}

func NewReference(author string, book_name string, page int) Reference {
	return Reference{
		author:   author,
		bookName: book_name,
		page:     page,
	}
}

func (r Reference) ToString() string {
	formatStr := "%s, page: %d, by %s"
	return fmt.Sprintf(formatStr, r.bookName, r.page, r.author)
}
