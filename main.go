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
	numbersFile    string
	count          int
	largestNumbers []int
}

// get minimum number and its index from slice of int
func min(s []int) (min [2]int) {
	min[1] = s[0]
	for i, v := range s {
		if min[1] > v {
			min[1] = v
			min[0] = i
		}
	}
	return
}

// calculates how long it takes to execute a function
func timeTrack(start time.Time, name string) {

	elapsed := time.Since(start)
	fmt.Printf("%v ran in %v \n", name, elapsed)
}

// generate []int slice from file
// []int slice contains largest numbers
func (s *TopNumbers) generateSlice() {

	defer timeTrack(time.Now(), "generateSlice")
	f, err := os.Open(s.numbersFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()
	// load file into memory, line by line instead of all the file at once
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var numbers []int
	for scanner.Scan() {
		number := scanner.Text()
		if n, err := strconv.Atoi(number); err == nil {
			if len(numbers) < s.count {
				numbers = append(numbers, n)
			} else if m := min(numbers); n > m[1] {
				numbers[m[0]] = n
			}
		}
	}
	s.largestNumbers = numbers
}

// print largest numbers
func (n *TopNumbers) printNumbers() {

	fmt.Printf("Largest %v numbers are: \n", n.count)
	sort.Sort(sort.Reverse(sort.IntSlice(n.largestNumbers)))
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
		numbersFile: file,
		count:       count,
	}

	n.generateSlice() // generate []int slice from file with largest X numbers
	n.printNumbers()  // print largest numbers
}
