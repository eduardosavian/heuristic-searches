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
	file1, _ := os.Create("hsmrs.txt")
    file2, _ := os.Create("hsm.txt")
    file3, _ := os.Create("hsmm.txt")
	defer file1.Close()

	fmt.Fprintf(file1, "Heuristica,N,M,Replicacao,Tempo,Iteracoes,Valor,Parametro\n")
    fmt.Fprintf(file2, "Heuristica,N,M,Replicacao,Tempo,Iteracoes,Valor,Parametro\n")
    fmt.Fprintf(file3, "Heuristica,N,M,Replicacao,Tempo,Iteracoes,Valor,Parametro\n")

	for _, m := range ms {
		for _, r := range rValues {
			tasks, n := generateInstance(m, r)

			for _, alpha := range alphas {
				for rep := 1; rep <= 10; rep++ {
					start := time.Now()
					bestSolution := monotoneRandomSearch(tasks, m, alpha)
					duration := time.Since(start).Seconds()

					fmt.Fprintf(file1, "monotona,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, duration, 1000, bestSolution.makespan, alpha)
				}

                for rep := 1; rep <= 10; rep++ {
					start := time.Now()
					bestSolution := blmMelhorDeterministico(tasks, m)
					duration := time.Since(start).Seconds()

					fmt.Fprintf(file2, "melhor,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, duration, 1000, bestSolution.makespan, alpha)
				}

                for rep := 1; rep <= 10; rep++ {
					start := time.Now()
					bestSolution := blmMelhorMelhorDeterministico(tasks, m, 10)
					duration := time.Since(start).Seconds()

					fmt.Fprintf(file3, "melhormelhor,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, duration, 1000, bestSolution.makespan, alpha)
				}
			}
		}
	}
}
