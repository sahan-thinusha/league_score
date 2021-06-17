package recommender

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	vm "LOLGamePredict/pkg/vector"
	log "github.com/sirupsen/logrus"
)

type NeighborhoodBasedRecommender struct {
	data        [][]float64
	neighbors   int
	numberItems int
}

type slice struct {
	sort.Interface
	idx []int
}

func NewNeighborhoodBasedRecommender(data [][]float64, k int) *NeighborhoodBasedRecommender {
	if len(data) == 0 {
		log.Fatalf("Dataset is empty")
	}

	return &NeighborhoodBasedRecommender{
		data:        data,
		neighbors:   k,
		numberItems: len(data[0]),
	}
}

func (nbr *NeighborhoodBasedRecommender) Recommend(items []float64, distanceMeasure vm.Distance, intercept, shuffle, serendipitous bool) ([]Recommendation, error) {
	recommendations, err := nbr.findKNearestNeighbors(items, distanceMeasure, intercept, shuffle, serendipitous)
	if err != nil {
		return nil, fmt.Errorf("Error encountered while finding K nearest neighbors: %v", err)
	}

	return recommendations, nil
}

func (nbr *NeighborhoodBasedRecommender) findKNearestNeighbors(items []float64, distanceMeasure vm.Distance, intercept, shuffle, serendipitous bool) ([]Recommendation, error) {
	var (
		d                 float64
		err               error
		distancesFromUser []Recommendation
		recommendations   []Recommendation
		order             []int
	)


	order = make([]int, len(nbr.data))
	for i := range order {
		order[i] = i
	}


	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
	}


	for i, val := range order {
		user := nbr.data[val]
		if len(user) != nbr.numberItems {
			return nil, fmt.Errorf("Incorrect number of items in vector")
		}

		switch distanceMeasure {
		case vm.Euclidean:
			d, err = vm.EuclideanDistance(items, user)
		case vm.Cosine:
			d, err = vm.CosineSimilarity(items, user)
		case vm.Manhattan:
			d, err = vm.ManhattanDistance(items, user)
		case vm.Pearson:
			d, err = vm.PearsonCorrelation(items, user)

			d = 1 - math.Abs(d)
		default:
			return nil, fmt.Errorf("Invalid distance measure: %v", distanceMeasure)
		}

		if err != nil {
			return nil, fmt.Errorf("Error calculating distance: %v", err)
		}

		distancesFromUser = append(distancesFromUser, MultipleRecommendation{
			Index:           i,
			Recommendation:  user,
			Distance:        d,
			DistanceMeasure: distanceMeasure,
		})
	}

	sort.Slice(distancesFromUser, func(i, j int) bool {
		return distancesFromUser[i].GetDistance() < distancesFromUser[j].GetDistance()
	})
	recommendations = distancesFromUser[:nbr.neighbors]


	if serendipitous {
		sereOptions := make([]float64, len(recommendations[0].GetRecommendation()))
		for _, reco := range distancesFromUser[nbr.neighbors:int(math.Min(float64(nbr.neighbors*2), float64(len(nbr.data))))] {
			for j, item := range reco.GetRecommendation() {
				sereOptions[j] += item
			}
		}
		s := newFloat64Slice(sereOptions...)
		sort.Sort(sort.Reverse(s))
		serendipitousRecommendation := make([]float64, len(recommendations[0].GetRecommendation()))
		for _, val := range s.idx[0:5] {
			serendipitousRecommendation[val] = 1
		}

		recommendations = append(recommendations, SerendipitousRecommendation{
			Recommendation:  serendipitousRecommendation,
			DistanceMeasure: distanceMeasure,
		})

	}


	if intercept {
		intercepts := recommendations[0].GetRecommendation()
		for _, val := range recommendations[1:len(recommendations)] {
			intercepts, err = vm.Intercept(intercepts, val.GetRecommendation())
			if err != nil {
				return nil, fmt.Errorf("Error calculating set intercept: %v", err)
			}
		}

		recommendations = []Recommendation{
			&SingleRecommendation{
				Recommendation:  intercepts,
				DistanceMeasure: distanceMeasure,
			},
		}

	}

	return recommendations, nil
}

func newFloat64Slice(n ...float64) *slice { return newSlice(sort.Float64Slice(n)) }

func newSlice(n sort.Interface) *slice {
	s := &slice{
		Interface: n,
		idx:       make([]int, n.Len()),
	}

	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

func (s slice) swap(i, j int) {
	s.Interface.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}
