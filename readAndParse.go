package main

import (
	"encoding/csv"
	"bufio"
	"os"
	"strconv"
	"log"
)

func Read(aFile string) [][]string {
	file, err := os.Open(aFile) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(file))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func Parse(records [][]string) []int64 {
	var parsedArray = make([]int64,0)
	
	for i := 1; i < len(records); i++ {
		f, err := strconv.ParseInt(records[i][1], 10, 64)
		
		if err != nil {
			log.Fatal(err)
		}
		parsedArray = append(parsedArray, (f-1)) //Adjust for R's stupid indexing
	}
	return parsedArray
}

