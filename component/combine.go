package component

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return a.firstName + " " + a.lastName
}

type post struct {
	title   string
	content string
	author
}

func (p post) detail() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("start create website")
	for _, value := range w.posts {
		value.detail()
		fmt.Println()
	}
}

func CombineTest() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	author2 := author{
		"chen",
		"hao",
		"java to go",
	}

	post1 := post{
		"how to learn go",
		"*********gtttt**********",
		author1,
	}

	post2 := post{
		"how to learn java",
		"*********java**********",
		author1,
	}

	post3 := post{
		"just learning go",
		"*********golang**********",
		author2,
	}

	post4 := post{
		"just learning java",
		"*********java**********",
		author2,
	}
	website1 := website{
		posts: []post{post1, post2, post3, post4},
	}
	website1.contents()
}
