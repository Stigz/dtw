package main

import (
	//"fmt"
	"math"
	//"github.com/chewxy/tightywhities"
)

// The cost function finds the absolute distance between values in each time series
func cost(s , c float64) float64 {
	return math.Abs(s-c)
}

// Return the minimum of the three neighboring indices in matrix
func min(a, b, c float64) float64 {
	var values = []float64{a,b,c}
	var min = values[0]

	for _, v:= range values{
		if v < min {
			min = v
		}
	}

	return min
}

// Creates a sequence of x values for a time series
func newSequence(start, end float64, numSteps int) []float64 {
	var step = (end-start)/float64(numSteps)

    if numSteps <= 0 || end < start {
        return []float64{}
    }
    s := make([]float64, 0, numSteps)
    for start <= end {
        s = append(s, start)
        start += step
    }
    return s
}

// The DTW algorithm creates a matrix of the accumulated distance between neighboring indices
func DTW(s, t []float64) ([][]float64, []int) {

	// Iniitalize matrix with row length s
	DTW := make([][]float64, len(s))

	// Give matrix column length t
	for i := range DTW {
		DTW[i] = make([]float64, len(t))
	}

	//Apply the DTW algorithm
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(t); j++ {
			var dist = cost(s[i],t[j])
			DTW[i][j] = dist + min(DTW[i-1][j], DTW[i][j-1],DTW[i-1][j-1]) //Notice ordering: diagonal movement is preferred
		}
	}

	//Store DTW row and column length
	var lengths = []int{len(s), len(t)}
	
	return DTW, lengths
}

// Finds the optimal path in DTW matrix
func Path(dtw [][]float64, lengths []int) ([]int64, []int64) {
	var rowLength = lengths[0]
	var colLength = lengths[1]

	var i, j = (rowLength-1), (colLength-1) //Start 

	//Values will store matrix values from dtw algorithm
	var values = make([]float64, 0)
	values = append(values, dtw[i][j])

	//Indices when matched are the indices corresponding to 
	//the path in the DTW distance matrix
	var index1 = make([]int64, 0)
	index1 = append([]int64{int64(i)}, index1...)
	var index2 = make([]int64, 0)
	index2 = append([]int64{int64(j)}, index2...)
	

	for i != 0 || j != 0 {
		var nextMin float64
		
		if i == 0 {
			nextMin = dtw[i][j-1]
		} else if j == 0 {
			nextMin = dtw[i-1][j]
		} else {
			nextMin = min(dtw[i-1][j], dtw[i][j-1], dtw[i-1][j-1])
		}

		if  i != 0 && j != 0 && dtw[i-1][j-1] == nextMin {
			i = i-1
			j = j-1
		} else if i != 0 && dtw[i-1][j] == nextMin {
			i = i-1
			j = j
		} else if j !=0 {
			i = i
			j = j-1
		}

		values = append(values, nextMin)
		index1 = append([]int64{int64(i)}, index1...)
		index2 = append([]int64{int64(j)}, index2...)
		
	}
	return index1, index2
}

func main() {
	/*
	var sequence = newSequence(0,6.28,100)
	var sequence2 = newSequence(0,3.14,50)
	var cosine = make([]float64,0)
	var sine = make([]float64,0)

	for _, x := range sequence {
		cosine = append(cosine, math.Cos(x))
		sine = append(sine, math.Cos(2*x))
	}
	
	var theDTW, lengths = DTW(sine, cosine)
	var indices1, indices2 = Path(theDTW, lengths)
	fmt.Println(indices1, indices2)
	*/

	//Print graph of indices in terminal
	//l := tightywhities.NewLine(xValues, yValues)
	//l.Plot(50, 25, os.Stdout)

}