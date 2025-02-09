package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)


func evaluate(solution []int, tasks []int, m int) int {
	load := make([]int, m)

	for i, machine := range solution {
		load[machine] += tasks[i]
	}

	maxLoad := 0
	for _, l := range load {
		if l > maxLoad {
			maxLoad = l
		}
	}

	return maxLoad
}

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

func main() {
	alphas := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	ms := []int{10, 20, 50}
	rValues := []float64{1.5, 2.0}
	file, _ := os.Create("resultados_monotona.txt")
	defer file.Close()

	fmt.Fprintf(file, "heuristica,n,m,replicacao,tempo,iteracoes,valor,parametro\n")

	for _, m := range ms {
		for _, r := range rValues {
			tasks, n := generateInstance(m, r)

			for _, alpha := range alphas {
				for rep := 1; rep <= 10; rep++ {
					start := time.Now()
					bestSolution := monotoneRandomSearch(tasks, m, alpha)
					duration := time.Since(start).Seconds()

					fmt.Fprintf(file, "monotona,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, duration, 1000, bestSolution.makespan, alpha)
				}
			}
		}
	}
}
