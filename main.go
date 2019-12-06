package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tjandrayana/play-error/cerr"
	"log"
	"time"
)

type Person struct {
	Name string `json:"name"`
}

func main() {

	a := `{"name":"lala}`

	te := Do(a)
	if te != nil {
		log.Printf("\nhei : %v\n", te)
	}

	fmt.Println(("====Done====="))

	err := Do1()
	if err != nil {
		log.Println("log println -> ", err)
		log.Printf("log printf : \n%+v\n", err)
	}
}

func Do(a string) error {
	var p Person
	if err := json.Unmarshal([]byte(a), &p); err != nil {
		return cerr.NewCustomErr(err)
	}
	return nil
}

func Do1() error {
	return Do2()
}

func Do2() error {
	return Do3()
}

func Do3() error {
	return Do4()
}

func Do4() error {
	t := time.Now().Unix()
	switch t % 3 {
	case 0:
		return errors.New("Case 0")
	case 1:
		return errors.New("Case 1")
	case 2:
		return errors.New("Case 2")
	}
	return nil
}
