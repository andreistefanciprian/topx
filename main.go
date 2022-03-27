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

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s ran in %s", name, elapsed)
}

func generateList(file string) []int {

	defer timeTrack(time.Now(), "generateList")

	var numbers []int

	f, err := os.Open(file)

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

	return numbers

}

func getLargestNumbers(numbers []int, count int) []int {

	defer timeTrack(time.Now(), "getLargestNumbers")

	// slower
	// sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	// var topNumbers []int = numbers[:count]

	sort.Ints(numbers)
	var topNumbers []int = numbers[len(numbers)-count:]

	return topNumbers
}

func main() {

	// define variables
	var numbersFile string = "numbers"
	var highestNumberCount int = 5
	var allNumbers []int
	var largestNumbers []int

	// get largest X numbers from file
	allNumbers = generateList(numbersFile)
	largestNumbers = getLargestNumbers(allNumbers, highestNumberCount)

	fmt.Println("\nLargest numbers are:")
	for _, each_ln := range largestNumbers {
		fmt.Println(each_ln)
	}
}
