# Heuristic Searches

Made by Eduardo Savian, Mateus Winter

## Description

This project implements two **Local Search** algorithms for solving the **Task Distribution Problem** among parallel machines. The goal is to minimize the **makespan** (total processing time of the busiest machine).

### Implemented Algorithms:

- **BL-Primeira Melhora (First Improvement Local Search)**: Iteratively swaps tasks between machines, stopping as soon as an improvement is found.
- **BL-Melhor Melhora (Best Improvement Local Search)**: Evaluates all possible swaps and applies the one that results in the best improvement.

Given:
- `m` parallel machines.
- `n = m * k` tasks, where `k ∈ {1.5, 2.0}`.
- Each task `t` has a processing time `p_t` randomly generated in the range `[1,100]`.

Find a task distribution that minimizes the **makespan**.

## How to Run

### Prerequisites

- Ensure you have [Go](https://golang.org/dl/) installed.

### Build

1. Open your terminal.
2. Navigate to the project directory.
3. Run the following command to build the executable:

#### Linux
```bash
go build -o local_search src/main.go src/methods.go src/utils.go
```

#### Windows
```ps1
go build -o local_search.exe src/main.go src/methods.go src/utils.go
```

### Execute

Run the program using:

#### Linux
```bash
./local_search 
```

#### Windows
```bash
./local_search.exe 
```

## Output

For each run, the program will display:
- Initial task distribution.
- Execution time.
- Makespan before and after optimization.
- Number of iterations performed.

## Bibliographic References



