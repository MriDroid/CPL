package main

import (
	"ZeinabProject/book"
	"ZeinabProject/reader"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func addBook(b book.Book) {
	data, _ := json.Marshal(b)
	response, err := http.Post("http://localhost:8050/addbook", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func bookSearchByID(id string) {
	data := fmt.Sprintf(`{"Id":"%s"}`, id)
	response, err := http.Post("http://localhost:8050/booksearchbyid", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func bookSearchByTitle(title string) {
	data := fmt.Sprintf(`{"Title":"%s"}`, title)
	response, err := http.Post("http://localhost:8050/booksearchbytitle", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func sortByTitle() {
	response, err := http.Get("http://localhost:8050/sortbytitle")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func sortByTime() {
	response, err := http.Get("http://localhost:8050/sortbytime")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func saveBookInfo(b book.Book) {
	data, _ := json.Marshal(b)
	response, err := http.Post("http://localhost:8050/savebookinfo", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func addReader(r reader.Reader) {
	data, _ := json.Marshal(r)
	response, err := http.Post("http://localhost:8050/addreader", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func removeReader(id string) {
	data := fmt.Sprintf(`{"Id":"%s"}`, id)
	response, err := http.Post("http://localhost:8050/removereader", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func readerSearchByID(id string) {
	data := fmt.Sprintf(`{"Id":"%s"}`, id)
	response, err := http.Post("http://localhost:8050/readersearchbyid", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func readerSearchByName(name string) {
	data := fmt.Sprintf(`{"Name":"%s"}`, name)
	response, err := http.Post("http://localhost:8050/readersearchbyname", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}

func main() {
	args := os.Args

	if len(args) > 1 {
		if args[1] == "--help" {
			fmt.Println(`
		to add book: -ab BookData(Id, Title, PublicationDate, Author, Genre, Publisher, Language)
		Example: -ab 1 CPL 1/1/2020 OmarSaad Sciantific OmarSaad English
		to update book data: -ub BookId BookData(Title, PublicationDate, Author, Genre, Publisher, Language)
		Example: -ub 1 CPL 1/1/2020 OmarSaad Sciantific OmarSaad English
		to search book by id: -Sb BookId
		Example: -Sb 1
		to search book by title: -sb BookTitle
		Example: -sb CPL
		to sort by title: -otb
		to sort by time: -odb
		to add reader: -ar ReaderData(Id, Name, Gender, BirthDay, Height, Weight)
		Example: -ar 1 OmarSaad male 28/9/1999 170 75
		to remove reader: -rr ReaderId
		Example: -rr 1
		to search reader by id: -Sr ReaderId
		Example: -Sr 1
		to search reader by name: -sr ReaderName
		Example: -sr OmarSaad
		`)
		} else if args[1] == "-ab" {
			if len(args) == 9 {
				_, err := time.Parse("2/1/2006", args[4])
				if err != nil {
					panic(err)
				}
				addBook(book.Book{
					Id:              args[2],
					Title:           args[3],
					PublicationDate: args[4],
					Author:          args[5],
					Genre:           args[6],
					Publisher:       args[7],
					Language:        args[8],
				})
			} else {
				fmt.Println("Please Make Sure From Inputs Data")
			}
		} else if args[1] == "-ub" {
			if len(args) == 9 {
				saveBookInfo(book.Book{
					Id:              args[2],
					Title:           args[3],
					PublicationDate: args[4],
					Author:          args[5],
					Genre:           args[6],
					Publisher:       args[7],
					Language:        args[8],
				})
			} else {
				fmt.Println("Please Make Sure From Inputs Data")
			}
		} else if args[1] == "-Sb" {
			if len(args) == 3 {
				bookSearchByID(args[2])
			} else {
				fmt.Println("Please Make Sure From Inputs Data")
			}
		} else if args[1] == "-sb" {
			if len(args) == 3 {
				bookSearchByTitle(args[2])
			} else {
				fmt.Println("Please Make Sure From Inputs Data")
			}
		} else if args[1] == "-otb" {
			sortByTitle()
		} else if args[1] == "-odb" {
			sortByTime()
		} else if args[1] == "-ar" {
			if len(args) == 8 {
				_, err := time.Parse("2/1/2006", args[5])
				if err != nil {
					panic(err)
				}
				addReader(reader.Reader{
					Id:       args[2],
					Name:     args[3],
					Gender:   args[4],
					Birthday: args[5],
					Height:   args[6],
					Weight:   args[7],
				})
			} else {
				fmt.Println("Please Make Sure From Inputs Data")
			}
		} else if args[1] == "-rr" {
			removeReader(args[2])
		} else if args[1] == "-Sr" {
			readerSearchByID(args[2])
		} else if args[1] == "-sr" {
			readerSearchByName(args[2])
		} else {
			fmt.Println("Run Client --help to get more information")
		}
	} else {
		fmt.Println("Run Client --help to get more information")
	}
}
