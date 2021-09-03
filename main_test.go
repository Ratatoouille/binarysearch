package main

import (
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

var sortSliceTests = []struct {
	in  []string
	out []string
}{
	{
		[]string{
			"bob",
			"alice",
			"john",
		},
		[]string{
			"alice",
			"bob",
			"john",
		},
	},
	{
		[]string{
			"bob",
			"alice",
			"john",
			"bob",
			"alice",
			"john",
		},
		[]string{
			"alice",
			"alice",
			"bob",
			"bob",
			"john",
			"john",
		},
	},
}

var searchTests = []struct {
	slice       []string
	target      string
	targetIndex int
}{
	{
		[]string{
			"alice",
			"bob",
			"john",
		},
		"alice",
		0,
	},
	{
		[]string{
			"alice",
			"bob",
			"john",
		},
		"john",
		2,
	},
	{
		[]string{
			"alice",
			"bob",
			"john",
		},
		"kate",
		-1,
	},
}

func TestIsSorted(t *testing.T) {
	for i, tt := range sortSliceTests {
		t.Run("isSorted"+string(rune(i)), func(t *testing.T) {
			if !isSorted(tt.in) {
			} else {
				t.Errorf("Test isSorted failed")
			}
			if isSorted(tt.out) {
			} else {
				t.Errorf("Test isSorted failed")
			}
		})
	}
}

func TestSortSlice(t *testing.T) {
	for i, tt := range sortSliceTests {
		t.Run("sort"+string(rune(i)), func(t *testing.T) {
			sortSlice(tt.in)
			if !equal(tt.in, tt.out) {
				t.Errorf("Test sortSlice failed")
			}
		})
	}
}

func TestReadFile(t *testing.T) {
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

func TestBinarySearchString(t *testing.T) {
	for i, tt := range searchTests {
		t.Run("sort"+string(rune(i)), func(t *testing.T) {
			index := binarySearchString(tt.slice, tt.target)
			if index != tt.targetIndex {
				t.Errorf("Test BinarySearchString failed")
			}
		})
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
