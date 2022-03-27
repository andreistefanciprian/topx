package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type TopNumbers struct {
	numbersFile        string
	highestNumberCount int
	allNumbers         []int
	largestNumbers     []int
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%v ran in %v \n", name, elapsed)
}

func (n *TopNumbers) generateList() {

	defer timeTrack(time.Now(), "generateList")

	f, err := os.Open(n.numbersFile)
	defer f.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var numbers []int
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

func (n *TopNumbers) printNumbers() {

	fmt.Printf("Largest %v numbers are: \n", n.highestNumberCount)
	for _, each_ln := range n.largestNumbers {
		fmt.Println(each_ln)
	}
}

func main() {

	// define cli flags
	var file string
	var count int
	flag.StringVar(&file, "file", "file.txt", "Specify file path.")
	flag.IntVar(&count, "count", 5, "Specify how many of the largest numbers you want to extract from file.")
	flag.Parse()

	// create struct with params
	n := TopNumbers{
		numbersFile:        file,
		highestNumberCount: count,
	}

	n.generateList()      // generate []int slice from file
	n.getLargestNumbers() // extract largest X numbers from []int slice
	n.printNumbers()      // print largest numbers
}
