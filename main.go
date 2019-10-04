/*
This repo contains my solution to the exercise
from https://gophercises.com
@Author DerAlexmeister
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
	"time"
)

/*
CSVPair represents a pair of question and answer
*/
type CSVPair struct {
	Question string
	Answer   int64
}

/*
Game represents generalinformation about your game
*/
type Game struct {
	CorrectAnswers int
	TimeOutAfter   int
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
timer is a timer that will wait a given time in Milliseconds
and in case the goroutine is not suspended it will exit
*/
func timer(timeout int) {
	time.Sleep(time.Millisecond * time.Duration(timeout))
	fmt.Printf("❔ \n >> You run out of time 😨 << ")
	os.Exit(0)
}

/*
AskAQuestion will print a Mathproblem
Then what for the User to answer the Question
After the Users answers it will return true or false in case the given
Answer was correct or not
*/
func askAQuestion(questionNumber int, pair CSVPair, time int) bool {
	var UserAnswer int64
	go timer(time)
	fmt.Printf(" ❔ Question #%d: %s = ", questionNumber, pair.Question)
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

/*
doesFileExist will return true in case a given file exists
if not it will return false
*/
func doesFileExist(path string) bool {
	state, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(state)
	}
	return true
}

func main() {
	CVSFile := flag.String("csv", "problems.csv", "Enter a the CSV-File with Questions like 5+5 = 10")
	TimeOut := flag.Int("timeout", 7000, "Enter the time every player get for a question")
	flag.Parse()
	if !doesFileExist(*CVSFile) {
		fmt.Printf("[-] An error occured while trying to read the File: %s \n", "File does not exist")
		os.Exit(0)
	}
	_Game := Game{
		CorrectAnswers: 0,
		TimeOutAfter:   *TimeOut,
	}
	fmt.Println("------------------------------------------")
	fmt.Printf(" 🎮 Gameconfig 🎮 : Timeout = %d Sec. ⏱️ \n", (_Game.TimeOutAfter / 1000))
	fmt.Println("------------------------------------------")
	flag.Parse()
	if !checkForCSVEnding(*CVSFile) {
		*CVSFile = (*CVSFile) + ".csv"
	}
	filePointer, err := Open(*CVSFile)
	if err != nil {
		os.Exit(1)
	} else {
		for number, pair := range readCSV(filePointer) {
			if askAQuestion(number+1, pair, _Game.TimeOutAfter) {
				fmt.Println(">> Correct Answer 👍")
			} else {
				fmt.Println(">> Wrong Answer 👎")
				fmt.Println("\n>> Boy oh boy 💩 💩")
				os.Exit(0)
			}
		}
		fmt.Println("\n>> Good job you completed the game! 👍 👍")
	}
	os.Exit(0)
}
