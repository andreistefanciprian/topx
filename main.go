package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sort"
	"time"
)

type TopNumbers struct {
	numbersFile string
	highestNumberCount int
	allNumbers []int
	largestNumbers []int
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s ran in %s", name, elapsed)
}

func (n *TopNumbers) generateList() {

	defer timeTrack(time.Now(), "generateList")
	var numbers []int
	f, err := os.Open(n.numbersFile)
	defer f.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		number := scanner.Text()
		if n, err := strconv.Atoi(number); err == nil {
			numbers = append(numbers, n)
		}	
	}
	n.allNumbers = numbers
}

func (n *TopNumbers) getLargestNumbers() {

	defer timeTrack(time.Now(), "getLargestNumbers")
	sort.Ints(n.allNumbers)
	n.largestNumbers = n.allNumbers[len(n.allNumbers)-n.highestNumberCount:]
}


func main() {

	n1 := TopNumbers{
		numbersFile: "numbers",
		highestNumberCount: 5,
	}
	n1.generateList()
	n1.getLargestNumbers()
	fmt.Println("\nLargest numbers are:")
	for _, each_ln := range n1.largestNumbers {
		fmt.Println(each_ln)
	}
}
