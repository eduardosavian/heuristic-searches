package main

import (
	"math/rand"
)

type Instance struct {
	Machines int
	Tasks    int
	JobTimes []int
}

type Solution struct {
	Allocation []int
	Makespan   int
}

// Algoritmo de Primeira Melhora
func blPrimeiraMelhora(inst Instance, sol Solution) Solution {
	improved := true

	for improved {
		improved = false
		for i := 0; i < inst.Machines; i++ {
			for j := 0; j < len(sol.Allocation); j++ {
				job := sol.Allocation[j]
				for k := 0; k < inst.Machines; k++ {
					if k != i {
						newLoad := make([]int, len(sol.Allocation))
						copy(newLoad, sol.Allocation)

						// Tentando mover o trabalho para outra máquina
						newLoad[i] -= job
						newLoad[k] += job

						if max(newLoad) < sol.Makespan {
							sol.Allocation = append(sol.Allocation[:j], sol.Allocation[j+1:]...)
							sol.Allocation = append(sol.Allocation, job)
							sol.Makespan = max(newLoad)
							improved = true
							break
						}
					}
				}
				if improved {
					break
				}
			}
			if improved {
				break
			}
		}
	}
	return sol
}

// Algoritmo de Melhor Melhora
func blMelhorMelhora(inst Instance, sol Solution) Solution {
	improved := true

	for improved {
		improved = false
		bestSolution := sol

		// Tentando melhorar a solução
		for i := 0; i < inst.Machines; i++ {
			for j := 0; j < len(sol.Allocation); j++ {
				job := sol.Allocation[j]
				for k := 0; k < inst.Machines; k++ {
					if k != i {
						newLoad := make([]int, len(sol.Allocation))
						copy(newLoad, sol.Allocation)

						// Tentando mover o trabalho para outra máquina
						newLoad[i] -= job
						newLoad[k] += job

						if max(newLoad) < bestSolution.Makespan {
							bestSolution = Solution{
								Allocation: append([]int(nil), sol.Allocation...),
								Makespan:  max(newLoad),
							}
							bestSolution.Allocation = append(sol.Allocation[:j], sol.Allocation[j+1:]...)
							bestSolution.Allocation = append(bestSolution.Allocation, job)
							improved = true
						}
					}
				}
			}
		}
		sol = bestSolution
	}
	return sol
}

// Algoritmo de Random Search Monotônico
func monotoneRandomSearch(tasks []int, m int, alpha float64) Solution {
	n := len(tasks)
	solution := make([]int, n)

	// Inicializando a solução com alocação aleatória
	for i := range solution {
		solution[i] = rand.Intn(m)
	}

	best := Solution{Allocation: solution, Makespan: evaluate(solution, tasks, m)}
	iterationsWithoutImprovement := 0

	// Executando o Random Search até atingir o número máximo de iterações
	for iterationsWithoutImprovement < 1000 {
		if rand.Float64() < alpha {
			newSolution := randomNeighbor(best.Allocation, m)
			newMakespan := evaluate(newSolution, tasks, m)

			// Se a nova solução for melhor, atualize a melhor solução
			if newMakespan < best.Makespan {
				best = Solution{Allocation: newSolution, Makespan: newMakespan}
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


func initialSolution(inst Instance) Solution {
	assignment := make([]int, len(inst.JobTimes)) // Atribuir aleatoriamente as tarefas
	for i := 0; i < len(inst.JobTimes); i++ {
		assignment[i] = rand.Intn(inst.Machines)
	}

	machineLoad := make([]int, inst.Machines)
	for i, machine := range assignment {
		machineLoad[machine] += inst.JobTimes[i]
	}

	makespan := max(machineLoad)

	return Solution{Allocation: assignment, Makespan: makespan}
}

