package component

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Id         int
	EmployeeNo string
	Name       string
}

func New(name string, employeeNo string, id int) Employee {
	employee := Employee{id, employeeNo, name}
	return employee
}

func (employee Employee) ToString() string {
	fmt.Printf("%s %s has %d leaves remaining", employee.Name, employee.EmployeeNo, employee.Id)
	return employee.Name
}

type add func(a int, b int) int

func AAndB() int {
	var a add = func(a, b int) int {
		return a + b
	}
	result := a(1, 2)
	fmt.Println("result is :", result)
	return result
}

func Simple(a func(a, b int) int) {
	fmt.Println("结果是：", a(1, 3))
}

func Simple1() func(a, b int) int {
	f := func(a, b int) int {
		return a + b + 1
	}
	return f
}

func AppenStr() func(a string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

/**
  计算斐波那契数列
**/
func MakeFibGen(f1 *int, f2 *int) func() int {
	f := func() int {
		*f1, *f2 = *f2, *f1+*f2
		return *f1
	}
	return f
}

var f1 int
var f2 int

//普通的斐波那契数列构造方式
func MakeFibGen1(p int, p1 int) int {
	f3 := 0
	if p1 == 0 {
		f1 = 0
		f2 = p
		f3 = p
		return f3
	}
	f3 = f1 + f2
	f1 = f2
	f2 = f3
	return f3

}

func Filter(e []Employee, f func(Employee) bool) []Employee {
	var r []Employee
	for _, value := range e {
		if f(value) {
			r = append(r, value)
		}
	}
	return r
}

func CreateQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("type", t)
	fmt.Println("value", v)
}

//insert create
func CompleteCreateQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")

}
