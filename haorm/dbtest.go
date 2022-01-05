package hao

import (
	haolog "gochen/haorm/log"
	haosession "gochen/haorm/session"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Test1 struct {
	Name string `haorm:"PRIMARY KEY"`
	Age  int
}

type Account struct {
	ID          int `haorm:"PRIMARY KEY"`
	UserName    string
	UserAccount string
	Password    string
	CreateDate  time.Time
	UpdateDate  time.Time
}

func (ac *Account) AfterQuery(s *haosession.Session) error {
	haolog.Info("after query", ac)
	ac.Password = "********"
	return nil
}

func (ac *Account) BeforeInsert(s *haosession.Session) error {
	haolog.Info("before insert step 1", ac)
	ac.Password = "default"
	haolog.Info("before insert step 2", ac)
	return nil
}

var (
	user1 = &Test1{"tom", 18}
	user2 = &Test1{"alan", 19}
)

var (
	account1 = &Account{1, "chen", "111", "1111", time.Now(), time.Now()}
)

func DbTest() {
	engine, err := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local", 1, 10, 10, 100, "blog")
	if err != nil {
		log.Fatal("start dataBase error", err)
	}
	defer engine.Close()
	s := engine.NewSession()
	//day 1 test
	// result1, err1 := s.Raw("DROP Table if EXISTS Test;").Exec()
	// if err1 != nil {
	// 	log.Println("drop table error", err1)
	// } else {
	// 	log.Println(result1)
	// }
	// result2, err2 := s.Raw("Create table test(Name text);").Exec()
	// if err2 != nil {
	// 	log.Println("drop table error", err2)
	// } else {
	// 	log.Println(result2)
	// }
	// result3, err3 := s.Raw("Create table test(Name text);").Exec()
	// if err3 != nil {
	// 	log.Println("drop table error", err3)
	// } else {
	// 	log.Println(result3)
	// }
	// result4, err6 := s.Raw("INSERT INTO test(`Name`) values (?), (?)", "Tom1", "Sam2").Exec()
	// if err6 != nil {
	// 	log.Fatal("some bad things happend", err6)
	// }
	// count, err5 := result4.RowsAffected()
	// if err5 != nil {
	// 	log.Println("some bad things happend", err5)
	// } else {
	// 	fmt.Printf("Exec success, %d affected\n", count)
	// }

	//day 2 test
	s.Model(&Account{})
	err2 := s.DropTable()
	err1 := s.CreateTable()
	_, err3 := s.Insert(account1)
	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatal("insert data error")
	}

	//limit test
	var users []Account
	err4 := s.Limit(1).Select(&users)
	if err4 != nil {
		log.Fatal("limit test error")
	}
	log.Println(users)
	//update test
	affected, _ := s.Where("username = ?", "chen").Update("UserAccount", "6666")
	u := &Account{}
	_ = s.OrderBy("ID Desc").SelectOne(u)

	if affected != 1 {
		log.Fatal("update test error")
	}
	log.Println(u)
	//delete and count test
	count, _ := s.Count()
	if count != 1 {
		log.Fatal("delet and count error")
	}
}
