package hao

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DbTest() {
	engine, err := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/blog", 1, 10, 10, 100)
	if err != nil {
		log.Fatal("start dataBase error", err)
	}
	defer engine.Close()
	s := engine.NewSession()
	result1, err1 := s.Raw("DROP Table if EXISTS Test;").Exec()
	if err1 != nil {
		log.Println("drop table error", err1)
	} else {
		log.Println(result1)
	}
	result2, err2 := s.Raw("Create table test(Name text);").Exec()
	if err2 != nil {
		log.Println("drop table error", err2)
	} else {
		log.Println(result2)
	}
	result3, err3 := s.Raw("Create table test(Name text);").Exec()
	if err3 != nil {
		log.Println("drop table error", err3)
	} else {
		log.Println(result3)
	}
	result4, err6 := s.Raw("INSERT INTO test(`Name`) values (?), (?)", "Tom1", "Sam2").Exec()
	if err6 != nil {
		log.Fatal("some bad things happend", err6)
	}
	count, err5 := result4.RowsAffected()
	if err5 != nil {
		log.Println("some bad things happend", err5)
	} else {
		fmt.Printf("Exec success, %d affected\n", count)
	}

}
