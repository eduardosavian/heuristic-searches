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
    Assignment [][]int
    MachineLoad []int
    Makespan   int
}

func generateInstance(machines int, taskFactor float64) Instance {
    tasks := int(float64(machines) * taskFactor)
    jobTimes := make([]int, tasks)
    for i := range jobTimes {
        jobTimes[i] = rand.Intn(100) + 1
    }
    return Instance{machines, tasks, jobTimes}
}

func initialSolution(inst Instance) Solution {
    assignment := make([][]int, inst.Machines)
    machineLoad := make([]int, inst.Machines)

    for i, jobTime := range inst.JobTimes {
        idx := i % inst.Machines
        assignment[idx] = append(assignment[idx], jobTime)
        machineLoad[idx] += jobTime
    }

    return Solution{assignment, machineLoad, max(machineLoad)}
}

func max(arr []int) int {
    maxVal := arr[0]
    for _, v := range arr {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
