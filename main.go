package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) == 2 {
		run(os.Args[1])
		return
	}
	log.Fatalln("add file as arg")
}

// run start program
func run(filepath string) {
	data, err := readFile(filepath)
	if err != nil {
		log.Fatalf("can't read file: %v", err)
	}

	fmt.Print("Enter the target: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	target := sc.Text()

	if !isSorted(data) {
		sortSlice(data)
	}

	fmt.Printf("data: %v\n", data)
	//index := binarySearchRecursively(data, target, data, len(data))
	index := binarySearchString(data, target)

	if index == -1 {
		fmt.Printf("target: %s not found in file", target)
		return
	}
	fmt.Printf("target is: %s, index is: %d", data[index], index)
	return
}

// isSorted check slice of strings is sorted
func isSorted(slice []string) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// sortSlice sort slice of strings
func sortSlice(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// readFile read file by path into slice of strings
func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	sc := bufio.NewScanner(file)
	var slice []string
	for sc.Scan() {
		slice = append(slice, sc.Text())
	}

	return slice, sc.Err()
}

// binarySearchString binary search in slice of strings
func binarySearchString(slice []string, target string) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return mid
		}
		if slice[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func binarySearchRecursively(slice []string, target string, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	mid := (lowIndex + highIndex) / 2

	if slice[mid] > target {
		return binarySearchRecursively(slice, target, lowIndex, mid)
	} else if slice[mid] < target {
		return binarySearchRecursively(slice, target, mid+1, highIndex)
	} else {
		return mid
	}
}
