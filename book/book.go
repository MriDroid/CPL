package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Book struct {
	Id              string
	Title           string
	PublicationDate string
	Author          string
	Genre           string
	Publisher       string
	Language        string
}

func (book Book) AddBook() {
	data, err := ioutil.ReadFile("./database/book.txt")
	if err != nil {
		fmt.Println(err)
	}
	var bookList []Book
	err = json.Unmarshal(data, &bookList)
	bookList = append(bookList, book)

	data, _ = json.Marshal(bookList)
	err = ioutil.WriteFile("./database/book.txt", []byte(string(data)+"\n"), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (book Book) SearchByTitle() (b Book, isFounded bool) {
	data, err := ioutil.ReadFile("./database/book.txt")
	if err != nil {
		fmt.Println(err)
	}
	var bookList []Book
	err = json.Unmarshal(data, &bookList)

	for i := range bookList {
		if bookList[i].Title == book.Title {
			return bookList[i], true
		}
	}
	return Book{}, false
}

func (book Book) SearchById() (b Book, isFounded bool) {
	data, err := ioutil.ReadFile("./database/book.txt")
	if err != nil {
		fmt.Println(err)
	}
	var bookList []Book
	err = json.Unmarshal(data, &bookList)

	for i := range bookList {
		if bookList[i].Id == book.Id {
			return bookList[i], true
		}
	}
	return Book{}, false
}

func (book Book) SaveBookInfo() (b Book, isUpdated bool) {
	data, err := ioutil.ReadFile("./database/book.txt")
	if err != nil {
		fmt.Println(err)
	}
	var bookList []Book
	err = json.Unmarshal(data, &bookList)

	for i := range bookList {
		if bookList[i].Id == book.Id {
			bookList[i].Title = book.Title
			bookList[i].PublicationDate = book.PublicationDate
			bookList[i].Author = book.Author
			bookList[i].Genre = book.Genre
			bookList[i].Publisher = book.Publisher
			bookList[i].Language = book.Language
			data, _ = json.Marshal(bookList)
			err = ioutil.WriteFile("./database/book.txt", []byte(string(data)+"\n"), 0777)
			if err != nil {
				fmt.Println(err)
			}
			return bookList[i], true
		}
	}
	return Book{}, false
}

func SortByTitle() []Book {
	data, err := ioutil.ReadFile("./database/book.txt")
	if err != nil {
		fmt.Println(err)
	}
	var bookList []Book
	err = json.Unmarshal(data, &bookList)

	sort.SliceStable(bookList, func(i, j int) bool {
		return bookList[i].Title < bookList[j].Title
	})

	return bookList
}

func SortByTime() []Book {
	layout := "2/1/2006"

	data, err := ioutil.ReadFile("./database/book.txt")
	if err != nil {
		fmt.Println(err)
	}
	var bookList []Book
	err = json.Unmarshal(data, &bookList)

	sort.SliceStable(bookList, func(i, j int) bool {
		t1, _ := time.Parse(layout, bookList[i].PublicationDate)
		t2, _ := time.Parse(layout, bookList[j].PublicationDate)

		return t1.Before(t2)
	})

	return bookList
}
