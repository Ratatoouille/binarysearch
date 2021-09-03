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
	err := fmt.Errorf("add file as arg")
	log.Fatalln(err)
}

//run program
func run(filepath string) {
	data, err := readFile(filepath)
	if err != nil {
		log.Fatalf("can't read file: %v", err)
	}

	fmt.Print("Enter the target: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	target := sc.Text()

	if sorted := isSorted(data); sorted {
		fmt.Printf("data: %v\n", data)

		index := binarySearchString(data, target)
		fmt.Printf("target is: %s, index is: %d", data[index], index)
		return
	}
	data = sortSlice(data)

	fmt.Printf("data: %v\n", data)
	index := binarySearchString(data, target)

	if index == -1 {
		fmt.Printf("target: %s not found in file\n", target)
		return
	}
	fmt.Printf("target is: %s, index is: %d\n", data[index], index)
}

//check slice is sorted
func isSorted(slice []string) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

//sort slice
func sortSlice(slice []string) []string {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return slice
}

//read file to string slice
func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	sc := bufio.NewScanner(file)
	var slice []string
	for sc.Scan() {
		slice = append(slice, sc.Text())
	}

	return slice, err
}

//binary search in string slice
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
