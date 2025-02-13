package main

import (
    "math/rand"
)

//  Non-Monotone Local Random Search
func NMLRS(tasks []int, m int, alpha float64) Solution {
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

// Monotone Local Search
func MLS(tasks []int, m int) Solution {
    n := len(tasks)
    solution := make([]int, n)

    best := Solution{allocation: solution, makespan: evaluate(solution, tasks, m)}
    iterationsWithoutImprovement := 0

    for iterationsWithoutImprovement < 1000 {
        newSolution := nextNeighbor(best.allocation, m)
        newMakespan := evaluate(newSolution, tasks, m)

        if newMakespan < best.makespan {
            best = Solution{allocation: newSolution, makespan: newMakespan}
            iterationsWithoutImprovement = 0
        } else {
            iterationsWithoutImprovement++
        }
    }

    return best
}

// Monotone Local Search Best
func MLSB(tasks []int, m int) Solution {
    return MLS(tasks, m)
}

// Monotone Local Search Best of Best
func MLSBB(tasks []int, m int, iterations int) Solution {
    bestOfBest := MLSB(tasks, m)

    for i := 1; i < iterations; i++ {
        candidate := MLSB(tasks, m)
        if candidate.makespan < bestOfBest.makespan {
            bestOfBest = candidate
        }
    }

    return bestOfBest
}
