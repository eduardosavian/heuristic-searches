package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    machinesSet := []int{10, 20, 50}
    taskFactors := []float64{1.5, 2.0}

    for _, machines := range machinesSet {
        for _, factor := range taskFactors {
            fmt.Printf("\nSimulando para %d máquinas e fator %.1f\n", machines, factor)
            for run := 0; run < 10; run++ {
                inst := generateInstance(machines, factor)
                initSol := initialSolution(inst)
                
                start := time.Now()
                sol1 := blPrimeiraMelhora(inst, initSol)
                elapsed1 := time.Since(start)
                
                start = time.Now()
                sol2 := blMelhorMelhora(inst, initSol)
                elapsed2 := time.Since(start)
                
                fmt.Printf("Execução %d - Primeira Melhora: Makespan %d (%.10fs), Melhor Melhora: Makespan %d (%.10fs)\n",
                    run+1, sol1.Makespan, elapsed1.Seconds(), sol2.Makespan, elapsed2.Seconds())
            }
        }
    }
}
