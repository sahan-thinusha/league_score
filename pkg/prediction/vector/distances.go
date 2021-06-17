package vector

import (
	"fmt"
	"github.com/dgryski/go-onlinestats"
	"math"
)

type Distance int

const (
	Euclidean Distance = 0
	Cosine    Distance = 1
	Manhattan Distance = 2
	Pearson   Distance = 3
)

func EuclideanDistance(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes len(x): %v, len(y): %v", len(x), len(y))
	}

	for index, element := range x {
		sum += math.Pow(element-y[index], 2)
	}

	return math.Sqrt(sum), nil
}

func ManhattanDistance(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes len(x): %v, len(y): %v", len(x), len(y))
	}

	for index, element := range x {
		sum += math.Abs(element - y[index])
	}

	return sum, nil
}


func CosineSimilarity(x []float64, y []float64) (float64, error) {
	dot, err := dotProduct(x, y)
	if err != nil {
		return 0, fmt.Errorf("Could not calculate the cosine similarity: %v", err)
	}

	fmt.Printf("cosine: %v\n", float64((dot))/(vectorEuclideanNorm(x)*vectorEuclideanNorm(y)))
	return float64((dot)) / (vectorEuclideanNorm(x) * vectorEuclideanNorm(y)), nil
}

func PearsonCorrelation(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("different slice sizes len(a): %v, len(b): %v", len(a), len(b))
	}

	return onlinestats.Pearson(a, b), nil
}


func vectorEuclideanNorm(vec []float64) float64 {
	if len(vec) == 0 {
		return 0.0
	}

	var sum float64
	for _, val := range vec {
		sum += math.Pow(float64(val), 2)
	}

	return math.Sqrt(sum)
}

func dotProduct(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes %v, %v", len(x), len(y))
	}

	for i, element := range x {
		sum += element * y[i]
	}

	return sum, nil
}
