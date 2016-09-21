package secretary

import (
	"math"
	"math/rand"
	"time"
)

// Simulate1e represents the 1/e rule by skipping n / math.E applicants
// and then picks the first candidate that appears. A candidate is an applicant
// with the best score up to that point
func Simulate1e(numApplicants int, numTrials int) float64 {
	ratings := make([]float64, numApplicants)
	for i := 0; i < len(ratings); i++ {
		ratings[i] = float64(i)
	}
	optimal := float64(ratings[len(ratings)-1])
	numBest := 0

	for i := 0; i < numTrials; i++ {
		shuffle(&ratings)
		selected := selectCandidate1e(ratings)
		if selected == -1 {
			continue
		}
		if ratings[selected] == optimal {
			numBest++
		}
	}
	return float64(numBest) / float64(numTrials)
}

// SimulateNth represents the candidate count rule by selecting the nth
// Candidate. A Candidate is an applicant with the best score up to that point
func SimulateNth(numApplicants int, numTrials int, toPick int) float64 {
	ratings := make([]float64, numApplicants)
	for i := 0; i < len(ratings); i++ {
		ratings[i] = float64(i)
	}
	optimal := float64(ratings[len(ratings)-1])
	numBest := 0

	for i := 0; i < numTrials; i++ {
		shuffle(&ratings)
		selected := selectNthCandidate(ratings, toPick)
		if selected == -1 {
			continue
		}
		if ratings[selected] == optimal {
			numBest++
		}
	}
	return float64(numBest) / float64(numTrials)
}

// SimulateKth simulates the successive non-candidate rule by skipping k non-candidates
// and then picking the first Candidate. A Candidate is an applicant
// with the best score up to that point
func SimulateKth(numApplicants int, numTrials int, toSkip int) float64 {
	ratings := make([]float64, numApplicants)
	for i := 0; i < len(ratings); i++ {
		ratings[i] = float64(i)
	}
	optimal := float64(ratings[len(ratings)-1])
	numBest := 0

	for i := 0; i < numTrials; i++ {
		shuffle(&ratings)
		selected := selectKthCandidate(ratings, toSkip)
		if selected == -1 {
			continue
		}
		if ratings[selected] == optimal {
			numBest++
		}
	}
	return float64(numBest) / float64(numTrials)
}

func selectNthCandidate(ratings []float64, toPick int) int {
	len := len(ratings)
	skipped := 0
	best := ratings[0]
	for cur := 0; cur < len; cur++ {
		rating := ratings[cur]
		if rating > best {
			skipped++
			best = rating
			if skipped == toPick {
				return cur
			}
		}
	}
	return -1
}

func selectKthCandidate(ratings []float64, nonToSkip int) int {
	len := len(ratings)
	skipped := 0
	best := ratings[0]
	for cur := 0; cur < len; cur++ {
		rating := ratings[cur]
		if rating > best {
			best = rating
			if skipped > nonToSkip {
				return cur
			}
		} else {
			skipped++
		}
	}
	return -1
}

func selectCandidate1e(ratings []float64) int {
	len := len(ratings)
	skip := round(float64(len) / math.E)
	best := ratings[0]
	hire := false
	for cur := 0; cur < len; cur++ {
		rating := ratings[cur]
		if cur > skip {
			hire = true
		}
		if rating > best {
			best = rating
			if hire {
				return cur
			}
		}
	}
	return -1
}

func shuffle(values *[]float64) {
	if len(*values) == 0 {
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(*values)-1; i++ {
		ri := randInt(i, len(*values))
		(*values)[i], (*values)[ri] = (*values)[ri], (*values)[i]
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}
