package main

import (
	"math"
	"math/rand"
	"time"
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

