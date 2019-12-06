package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var globalcounter int

func main() {
	startTime := time.Now()
	globalcounter = 0
	csvfile := flag.String("csv", "/home/manikanta/go_projects/src/myAlgorithms/text.csv", "a csv file for answers and question for the quiz")
	flag.Parse()

	readFile, err := os.OpenFile(*csvfile, os.O_RDWR, 0776)

	if err != nil {
		fmt.Errorf("the file you are trying to open is not present %v\n", *csvfile)

	}
	fileinfo, _ := readFile.Stat()
	filesize := int(fileinfo.Size())
	data := make([]byte, filesize)
	count, _ := readFile.Read(data)
	fmt.Println(count)
	displayer(data)
	fmt.Println("the time taken by you to complete the quizz is ", time.Now().Sub(startTime))
	fmt.Println("your score is ", globalcounter)

}
func displayer(data []byte) {
	for i := 0; i < len(data); i++ {
		fmt.Print(string(data[i]))
		if rune(data[i]) == ' ' {

			val := i
			val++
			v := answerparser(data[val:])

			val--
			i = val + len(v) + 1
			intval, _ := strconv.Atoi(v)

			if inputter(intval) == true {
				globalcounter++
			}

		}

	}

}
func answerparser(data []byte) string {

	preindexptr := 0
	postindexptr := 0
	for i := 0; i < len(data); i++ {
		preindexptr++
		if rune(data[i]) == ',' {
			preindexptr--
			s := (string(data[postindexptr:preindexptr]))

			postindexptr = preindexptr
			return s

		}
	}
	return ""
}
func inputter(val int) bool {
	var n int
	fmt.Scan(&n)
	if n == val {
		fmt.Println("correct!")
		return true
	}
	fmt.Println("not correct!!")
	return false
}
