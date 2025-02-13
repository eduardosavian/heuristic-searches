package main

import (
	"math"
	"math/rand"
	"time"
	"os"
	"fmt"
)

type Solution struct {
	allocation []int
	makespan     int;
}

func generateInstance(m int, r float64) ([]int, int) {
	n := int(math.Pow(float64(m), float64(r)))

	tasks := make([]int, n)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range tasks {
		tasks[i] = int(rng.Intn(100) + 1)
	}

	return tasks, n
}

func randomNeighbor(solution []int, m int) []int {
    newSolution := make([]int, len(solution))
    copy(newSolution, solution)

    randIndex := rand.Intn(len(solution))
    newSolution[randIndex] = rand.Intn(m)

    return newSolution
}

func nextNeighbor(solution []int, m int) []int {
    newSolution := make([]int, len(solution))
    copy(newSolution, solution)

    for i := range newSolution {
        newSolution[i] = (newSolution[i] + 1) % m
        if newSolution[i] != 0 {
            break
        }
    }

    return newSolution
}

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

func createFile() (*os.File, error) {
    file, err := os.Create("hs.csv")
    if err != nil {
        return nil, fmt.Errorf("erro ao criar file: %v", err)
    }

    return file, nil
}
