package main

import (
	"fmt"
	"github.com/kompiuter/secretary"
)

func main() {
	numApplicants := 100
	numTrials := 10000
	e := secretary.Simulate1e(numApplicants, numTrials)
	nth := secretary.SimulateNth(numApplicants, numTrials, 4)
	kth := secretary.SimulateKth(numApplicants, numTrials, int(numApplicants/2)) // successive non-candidate rule generally performs well with k = n / 2

	fmt.Printf("%-25.25s %.3f\n", "1/e rule:", e)
	fmt.Printf("%-25.25s %.3f\n", "candidate count rule:", nth)
	fmt.Printf("%-25.25s %.3f", "non-candidate rule:", kth)
}
