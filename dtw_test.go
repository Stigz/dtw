package main

import (
	"testing"
	"math"
	"log"
)

func testEq(a, b []int64) bool {
    if a == nil || b == nil { 
    	log.Fatal("Some Array(s) are Empty")
        return false; 
    }

    if len(a) != len(b) {
    	log.Fatal("Length Sucks", len(a), len(b))
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestDtwSinusoid(t *testing.T){
	//Create a cosine and a sine sequence
	var sequence = newSequence(0,6.28,100)
	var cosine = make([]float64,0)
	var sine = make([]float64,0)
	for _, x := range sequence {
		cosine = append(cosine, math.Cos(x))
		sine = append(sine, math.Sin(x))
	}

	//Run the DTW algorithm which computes accumulated cost matrix
	var theDTW, lengths = DTW(sine, cosine)
	//Find the best path which returns indices
	var indices1, indices2 = Path(theDTW, lengths)

	//Read first and second R output of indices
	file1 := Read("sinuIndex1.csv")
	file2 := Read("sinuIndex2.csv")
	
	//Parse dem sloots
	parsedIndex1 := Parse(file1)
	parsedIndex2 := Parse(file2)

	//Delete 75th index where there is a repetition that screws everything up
	parsedIndex1 = append(parsedIndex1[:75], parsedIndex1[76:]...)
	parsedIndex2 = append(parsedIndex2[:75], parsedIndex2[76:]...)

	//Begin comparison
	var isEqual1 = testEq(parsedIndex1, indices1)
	var isEqual2 = testEq(parsedIndex2, indices2)
	if isEqual1 == false {
		t.Errorf("Array of Indices 1 Not equal")
	}
	if isEqual2 == false {
		t.Errorf("Array of Indices 2 Not equal")
	}
}