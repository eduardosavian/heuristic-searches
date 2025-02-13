package main



func blPrimeiraMelhora(inst Instance, sol Solution) Solution {
    improved := true

    for improved {
        improved = false
        for i := 0; i < inst.Machines; i++ {
            for j := 0; j < len(sol.Assignment[i]); j++ {
                job := sol.Assignment[i][j]
                for k := 0; k < inst.Machines; k++ {
                    if k != i {
                        newLoad := make([]int, len(sol.MachineLoad))
                        copy(newLoad, sol.MachineLoad)

                        newLoad[i] -= job
                        newLoad[k] += job

                        if max(newLoad) < sol.Makespan {
                            sol.Assignment[i] = append(sol.Assignment[i][:j], sol.Assignment[i][j+1:]...)
                            sol.Assignment[k] = append(sol.Assignment[k], job)
                            sol.MachineLoad = newLoad
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

func blMelhorMelhora(inst Instance, sol Solution) Solution {
    improved := true

    for improved {
        improved = false
        bestSolution := sol

        for i := 0; i < inst.Machines; i++ {
            for j := 0; j < len(sol.Assignment[i]); j++ {
                job := sol.Assignment[i][j]
                for k := 0; k < inst.Machines; k++ {
                    if k != i {
                        newLoad := make([]int, len(sol.MachineLoad))
                        copy(newLoad, sol.MachineLoad)

                        newLoad[i] -= job
                        newLoad[k] += job

                        if max(newLoad) < bestSolution.Makespan {
                            bestSolution = Solution{
                                Assignment: append([][]int(nil), sol.Assignment...),
                                MachineLoad: newLoad,
                                Makespan: max(newLoad),
                            }
                            bestSolution.Assignment[i] = append(sol.Assignment[i][:j], sol.Assignment[i][j+1:]...)
                            bestSolution.Assignment[k] = append(sol.Assignment[k], job)
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