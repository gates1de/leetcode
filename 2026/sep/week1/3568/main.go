package main

import (
	"fmt"
)

func minMoves(classroom []string, energy int) int {
	m := len(classroom)
	if m == 0 {
		return -1
	}
	n := len(classroom[0])

	start := int(0)
	litterCount := int(0)
	litterIndex := make([][]int, m)
	for i := range classroom {
		litterIndex[i] = make([]int, n)

		for j := range classroom[i] {
			litterIndex[i][j] = -1

			switch classroom[i][j] {
			case 'S':
				start = i*n + j
			case 'L':
				litterIndex[i][j] = litterCount
				litterCount++
			}
		}
	}

	if litterCount == 0 {
		return 0
	}

	energyStates := energy + 1
	stateCount := m * n * (1 << litterCount) * energyStates
	visited := make([]bool, stateCount)

	type state struct {
		position int
		mask     int
		energy   int
	}

	queue := make([]state, 0)
	queue = append(queue, state{position: start, energy: energy})
	visited[start*energyStates+energy] = true

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	fullMask := (1 << litterCount) - 1
	moves := int(0)

	for head := 0; head < len(queue); {
		levelEnd := len(queue)
		for head < levelEnd {
			current := queue[head]
			head++
			if current.mask == fullMask {
				return moves
			}
			if current.energy == 0 {
				continue
			}

			r, c := current.position/n, current.position%n
			for _, direction := range directions {
				nr, nc := r+direction[0], c+direction[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n || classroom[nr][nc] == 'X' {
					continue
				}

				nextEnergy := current.energy - 1
				if classroom[nr][nc] == 'R' {
					nextEnergy = energy
				}
				nextMask := current.mask
				if index := litterIndex[nr][nc]; index >= 0 {
					nextMask |= 1 << index
				}

				nextPosition := nr*n + nc
				visitedIndex := (nextMask*m*n+nextPosition)*energyStates + nextEnergy
				if visited[visitedIndex] {
					continue
				}
				visited[visitedIndex] = true
				queue = append(queue, state{
					position: nextPosition,
					mask:     nextMask,
					energy:   nextEnergy,
				})
			}
		}
		moves++
	}

	return -1
}

func main() {
	// result: 2
	// classroom := []string{"S.", "XL"}
	// energy := int(2)

	// result: 3
	// classroom := []string{"LS", "RL"}
	// energy := int(4)

	// result: -1
	classroom := []string{"L.S", "RXL"}
	energy := int(3)

	// result:
	// classroom := []string{}
	// energy := int(0)

	result := minMoves(classroom, energy)
	fmt.Printf("result = %v\n", result)
}
