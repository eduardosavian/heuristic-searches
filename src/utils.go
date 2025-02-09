package main

import (
	"math/rand"
	"time"
)

func generateInstance(m int, r float64) ([]int, int) {
	n := int(float64(m) * r)
	tasks := make([]int, n)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range tasks {
		tasks[i] = int(rng.Intn(100) + 1)
	}

	return tasks, n
}

type Solution struct {
	allocation []int
	makespan     int;
}