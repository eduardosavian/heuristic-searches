package main

import (
	"fmt"
	"time"
)

func main() {
	file, err := createFile()
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

	header := "Heuristica,N,M,Replicacao,Tempo,Iteracoes,Valor,Parametro\n"

	fmt.Fprint(file, header)

	alphas := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	ms := []int{10, 20, 50}
	rValues := []float64{1.5, 2.0}

	for _, m := range ms {
		for _, r := range rValues {
			tasks, n := generateInstance(m, r)

			for _, alpha := range alphas {
				for rep := 1; rep <= 10; rep++ {
					start := time.Now()
					bestSolution := NMLRS(tasks, m, alpha)
					duration := time.Since(start).Seconds()
					fmt.Fprintf(file, "nmlrs,%d,%d,%d,%.9f,%d,%d,%.1f\n", n, m, rep, duration, 1000, bestSolution.makespan, alpha)

					start = time.Now()
					bestSolution = MLSB(tasks, m)
					duration = time.Since(start).Seconds()
					fmt.Fprintf(file, "mlsb,%d,%d,%d,%.9f,%d,%d,NA\n", n, m, rep, duration, 1000, bestSolution.makespan)

					start = time.Now()
					bestSolution = MLSBB(tasks, m, 10)
					duration = time.Since(start).Seconds()
					fmt.Fprintf(file, "mlsbb,%d,%d,%d,%.9f,%d,%d,NA\n", n, m, rep, duration, 1000, bestSolution.makespan)
				}
			}
		}
	}
}
