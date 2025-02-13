package main

import (
    "math/rand"
)


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

func monotoneSearch(tasks []int, m int) Solution {
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

func blmMelhorDeterministico(tasks []int, m int) Solution {
    return monotoneSearch(tasks, m)
}

func blmMelhorMelhorDeterministico(tasks []int, m int, iterations int) Solution {
    bestOfBest := blmMelhorDeterministico(tasks, m)

    for i := 1; i < iterations; i++ {
        candidate := blmMelhorDeterministico(tasks, m)
        if candidate.makespan < bestOfBest.makespan {
            bestOfBest = candidate
        }
    }

    return bestOfBest
}
