package main

import(
    "math/rand"
)

func randomNeighbor(solution []int, m int) []int {
	newSolution := make([]int, len(solution))
	copy(newSolution, solution)

	randIndex := rand.Intn(len(solution))
	newSolution[randIndex] = rand.Intn(m)

	return newSolution
}

func monotoneRandomSearch(tasks []int, m int, alpha float64) Solution {
	n := len(tasks)
	solution := make([]int, n)

	for i := range solution {
		solution[i] = rand.Intn(m)
	}

	best := Solution{allocation: solution, makespan: evaluate(solution, tasks, m)}
	iterationsWithoutImprovement := 0

	for iterationsWithoutImprovement < 1000 {
		if rand.Float64() < alpha {
			newSolution := randomNeighbor(best.allocation, m)
			newMakespan := evaluate(newSolution, tasks, m)

			if newMakespan < best.makespan {
				best = Solution{allocation: newSolution, makespan: newMakespan}
				iterationsWithoutImprovement = 0
			} else {
				iterationsWithoutImprovement++
			}
		} else {
			iterationsWithoutImprovement++
		}
	}

	return best
}