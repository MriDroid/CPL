package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Reader struct {
	Id       string
	Name     string
	Gender   string
	Birthday string
	Height   string
	Weight   string
}

func removeIndex(r []Reader, index int) []Reader {
	return append(r[:index], r[index+1:]...)
}

func (reader Reader) AddReader() {
	data, err := ioutil.ReadFile("./database/reader.txt")
	if err != nil {
		fmt.Println(err)
	}
	var readerList []Reader
	err = json.Unmarshal(data, &readerList)
	readerList = append(readerList, reader)

	data, _ = json.Marshal(readerList)
	err = ioutil.WriteFile("./database/reader.txt", []byte(string(data)+"\n"), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (reader Reader) RemomveReader() bool {
	data, err := ioutil.ReadFile("./database/reader.txt")
	if err != nil {
		fmt.Println(err)
	}
	var readerList []Reader
	err = json.Unmarshal(data, &readerList)
	for i := range readerList {
		if readerList[i].Id == reader.Id {
			readerList = removeIndex(readerList, i)
			data, _ = json.Marshal(readerList)
			err = ioutil.WriteFile("./database/reader.txt", []byte(string(data)+"\n"), 0777)
			if err != nil {
				fmt.Println(err)
			}
			return true
		}
	}
	return false
}

func (reader Reader) SearchById() (r Reader, isFounded bool) {
	data, err := ioutil.ReadFile("./database/reader.txt")
	if err != nil {
		fmt.Println(err)
	}
	var readerList []Reader
	err = json.Unmarshal(data, &readerList)

	for i := range readerList {
		if readerList[i].Id == reader.Id {
			return readerList[i], true
		}
	}
	return Reader{}, false
}

func (reader Reader) SearchByName() (r Reader, isFounded bool) {
	data, err := ioutil.ReadFile("./database/reader.txt")
	if err != nil {
		fmt.Println(err)
	}
	var readerList []Reader
	err = json.Unmarshal(data, &readerList)

	for i := range readerList {
		if readerList[i].Name == reader.Name {
			return readerList[i], true
		}
	}
	return Reader{}, false
}
