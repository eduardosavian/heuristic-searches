package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	alphas := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	ms := []int{10, 20, 50}
	rValues := []float64{1.5, 2.0}
	file, _ := os.Create("hs.txt")
	defer file.Close()

	fmt.Fprintf(file, "Heuristica,N,M,Replicacao,Tempo,Iteracoes,Valor,Parametro\n")

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
