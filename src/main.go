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

	// Cabeçalho do arquivo de saída
	fmt.Fprintf(file, "Heuristica,N,M,Replicacao,Tempo,Iteracoes,Valor,Parametro\n")

	// Iteração para diferentes configurações
	for _, m := range ms {
		for _, r := range rValues {
			tasks, n := generateInstance(m, r)

			// Para cada valor de alpha, executar os algoritmos
			for _, alpha := range alphas {
				for rep := 1; rep <= 10; rep++ {
					// Primeira Melhora
					start := time.Now()
					initialSolutionPrimeira := initialSolution(Instance{Machines: m, Tasks: n, JobTimes: tasks})
					bestSolutionPrimeira := blPrimeiraMelhora(Instance{Machines: m, Tasks: n, JobTimes: tasks}, initialSolutionPrimeira)
					durationPrimeira := time.Since(start).Seconds()

					// Melhor Melhora
					start = time.Now()
					initialSolutionMelhor := initialSolution(Instance{Machines: m, Tasks: n, JobTimes: tasks})
					bestSolutionMelhor := blMelhorMelhora(Instance{Machines: m, Tasks: n, JobTimes: tasks}, initialSolutionMelhor)
					durationMelhor := time.Since(start).Seconds()

					// Random Search Monotônico
					start = time.Now()
					bestSolutionRandom := monotoneRandomSearch(tasks, m, alpha)
					durationRandom := time.Since(start).Seconds()

					// Escrever os resultados para o arquivo
					fmt.Fprintf(file, "primeira_melhora,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, durationPrimeira, 1000, bestSolutionPrimeira.Makespan, alpha)
					fmt.Fprintf(file, "melhor_melhora,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, durationMelhor, 1000, bestSolutionMelhor.Makespan, alpha)
					fmt.Fprintf(file, "random_search,%d,%d,%d,%.4f,%d,%d,%.1f\n", n, m, rep, durationRandom, 1000, bestSolutionRandom.Makespan, alpha)
				}
			}
		}
	}
}
