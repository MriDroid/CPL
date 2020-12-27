package main

import (
	"ZeinabProject/book"
	"ZeinabProject/reader"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Run Client --help to get more information\n")
}

func addBook(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var b book.Book
		err := json.Unmarshal(body, &b)
		if err != nil {
			panic(err)
		}
		b.AddBook()
		fmt.Fprintln(response, "Done")
	}
}

func saveBookInfo(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var b book.Book
		err := json.Unmarshal(body, &b)
		if err != nil {
			panic(err)
		}
		b, isUpdated := b.SaveBookInfo()
		if isUpdated {
			data, _ := json.Marshal(b)
			fmt.Fprintln(response, string(data))
		} else {
			fmt.Fprintln(response, "Book is not Founded")
		}
	}
}

func bookSearchByTitle(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var b book.Book
		err := json.Unmarshal(body, &b)
		if err != nil {
			panic(err)
		}
		b, isFounded := b.SearchByTitle()
		if isFounded {
			data, _ := json.Marshal(b)
			fmt.Fprintln(response, string(data))
		} else {
			fmt.Fprintln(response, "Book is not Founded")
		}
	}
}

func bookSearchByID(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var b book.Book
		err := json.Unmarshal(body, &b)
		if err != nil {
			panic(err)
		}
		b, isFounded := b.SearchById()
		if isFounded {
			data, _ := json.Marshal(b)
			fmt.Fprintln(response, string(data))
		} else {
			fmt.Fprintln(response, "Book is not Founded")
		}
	}
}

func sortByTitle(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var b []book.Book
		b = book.SortByTitle()
		data, _ := json.Marshal(b)
		fmt.Fprintln(response, string(data))
	}
}

func sortByTime(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var b []book.Book
		b = book.SortByTime()
		data, _ := json.Marshal(b)
		fmt.Fprintln(response, string(data))
	}
}

func addReader(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var r reader.Reader
		err := json.Unmarshal(body, &r)
		if err != nil {
			panic(err)
		}
		r.AddReader()
		fmt.Fprintln(response, "Done")
	}
}

func removeReader(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var r reader.Reader
		err := json.Unmarshal(body, &r)
		if err != nil {
			panic(err)
		}
		isRemoved := r.RemomveReader()
		if isRemoved {
			fmt.Fprintln(response, "Done")
		} else {
			fmt.Fprintln(response, "Reader is not Founded")
		}
	}
}

func readerSearchByID(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var r reader.Reader
		err := json.Unmarshal(body, &r)
		if err != nil {
			panic(err)
		}
		b, isFounded := r.SearchById()
		if isFounded {
			data, _ := json.Marshal(b)
			fmt.Fprintln(response, string(data))
		} else {
			fmt.Fprintln(response, "Reader is not Founded")
		}
	}
}

func readerSearchByName(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var r reader.Reader
		err := json.Unmarshal(body, &r)
		if err != nil {
			panic(err)
		}
		b, isFounded := r.SearchByName()
		if isFounded {
			data, _ := json.Marshal(b)
			fmt.Fprintln(response, string(data))
		} else {
			fmt.Fprintln(response, "Reader is not Founded")
		}
	}
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/addbook", addBook)
	http.HandleFunc("/savebookinfo", saveBookInfo)
	http.HandleFunc("/booksearchbytitle", bookSearchByTitle)
	http.HandleFunc("/booksearchbyid", bookSearchByID)
	http.HandleFunc("/sortbytitle", sortByTitle)
	http.HandleFunc("/sortbytime", sortByTime)
	http.HandleFunc("/addreader", addReader)
	http.HandleFunc("/removereader", removeReader)
	http.HandleFunc("/readersearchbyid", readerSearchByID)
	http.HandleFunc("/readersearchbyname", readerSearchByName)

	err := http.ListenAndServe(":8050", nil)
	if err != nil {
		panic(err)
	}
}
