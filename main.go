/**
* This repo contains my solution to the exercise
* from https://gophercises.com
 */

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
CSVPair represents a pair of question and answer
*/
type CSVPair struct {
	Question string
	Answer   int64
}

/*
Open Function to open a given File
will return a Filedescriptor
*/
func Open(name string) (*os.File, error) {
	return os.Open(name)
}

/*
readCSV Is a Function to read a csv
add all Lines in a slide an return
*/
func readCSV(file *os.File) []CSVPair {
	reader := csv.NewReader(bufio.NewReader(file))
	pairs := []CSVPair{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Printf("[-] An error occured while trying to read the File: %s \n", error)
			break
		} else {
			a, err := strconv.ParseInt(line[1], 10, 8)
			if err != nil {
				fmt.Printf("[-] An error occured while trying cast from String to Int: %s \n", error)
				continue
			} else {
				pairs = append(pairs, CSVPair{
					Question: line[0],
					Answer:   a,
				})
			}
		}
	}
	return pairs
}

/*
AskAQuestion will print a Mathproblem
Then what for the User to answer the Question
After the Users answers it will return true or false in case the given
Answer was correct or not
*/
func askAQuestion(questionNumber int, pair CSVPair) bool {
	var UserAnswer int64
	fmt.Printf("Question #%d: %s = ", questionNumber, pair.Question)
	fmt.Scanf("%d", &UserAnswer)
	if UserAnswer == pair.Answer {
		return true
	}
	return false
}

/*
checkForCSVEnding check whether a given filename
ends with .csv or not. Returned will be either true or false
*/
func checkForCSVEnding(csvFile string) bool {
	if strings.HasSuffix(csvFile, ".csv") {
		return true
	}
	return false
}

func main() {
	CVSFile := flag.String("csv", "problems.csv", "Enter as a Parameter the CSV-File")
	flag.Parse()
	if !checkForCSVEnding(*CVSFile) {
		*CVSFile = (*CVSFile) + ".csv"
	}
	filePointer, err := Open(*CVSFile)
	if err != nil {
		os.Exit(1)
	} else {
		for number, pair := range readCSV(filePointer) {
			if askAQuestion(number, pair) {
				fmt.Println(">> Correct Answer")
			} else {
				fmt.Println(">> Wrong Answer")
				os.Exit(0)
			}
		}
	}
	os.Exit(0)
}
