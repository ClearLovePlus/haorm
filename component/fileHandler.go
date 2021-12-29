package component

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func HandlerTxt() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("file reading error", err)
		return
	}
	fmt.Println("contents of file is", string(data))
}

func BfReader() {
	//分区缓存读取文件
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}
