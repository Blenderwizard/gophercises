package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func quiz(timeLimit int, problems []problem) {
	correct := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
Loop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\n")
			break Loop
		case answer := <-answerCh:
			if strings.ToLower(answer) == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You got %d out of %d correct\n", correct, len(problems))
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file with the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit in seconds for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Error opening the csv file: %s", *csvFile))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error error while parsing the csv file: %s", *csvFile))
	}

	quiz(*timeLimit, parseLines(lines))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.ToLower(strings.TrimSpace(line[1])),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Printf("%s", msg)
	os.Exit(1)
}
