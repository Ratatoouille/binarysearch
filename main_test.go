package main

import (
	"fmt"
	"testing"
)

var testSlice = []string{
	"bob",
	"alice",
	"john",
	"kate",
	"sophia",
	"emma",
	"james",
	"robert",
}

var notSortedSlice = []string{
	"bob",
	"alice",
	"john",
}

var sortedSlice = []string{
	"alice",
	"bob",
	"john",
}

func TestOkIsSorted(t *testing.T) {
	if isSorted(sortedSlice) {
		fmt.Println("Test isSort OK: slice is sorted")
		return
	}
	t.Errorf("Test isSort failed")
}

func TestFailIsSorted(t *testing.T) {
	if !isSorted(notSortedSlice) {
		fmt.Println("Test isSort OK: slice not sorted")
		return
	}
	t.Errorf("Test isSort failed")
}

func TestSortSlice(t *testing.T) {
	slice := sortSlice(notSortedSlice)
	if !equal(slice, sortedSlice) {
		t.Errorf("Test sortSlice failed")
	}
}

func TestOkReadFile(t *testing.T) {
	slice, err := readFile("./test_data.txt")
	if err != nil {
		t.Errorf("Test readFile failed: %s", err)
		return
	}
	if !equal(slice, testSlice) {
		t.Errorf("Test readFile failed")
		return
	}
}

func TestFindHightBinarySearchString(t *testing.T) {
	i := binarySearchString(sortedSlice, "alice")
	if i != 0 {
		t.Errorf("Test binarySearchString failed: target wasn't found")
	}
}

func TestFindLowBinarySearchString(t *testing.T) {
	i := binarySearchString(sortedSlice, "john")
	if i != 2 {
		t.Errorf("Test binarySearchString failed: target wasn't found")
	}
}

func TestNotFindBinarySearchString(t *testing.T) {
	i := binarySearchString(sortedSlice, "kate")
	if i != -1 {
		t.Errorf("Test binarySearchString failed: target was found")
	}
}

func TestRun(t *testing.T) {
	run("./test_data.txt")
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
