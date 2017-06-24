package main

/** GameState State of the game **/
type GameState struct {
	Started    bool
	Generation uint
	Current    [NumRows][NumColumns]bool
}

func step(state GameState) GameState {
	previous := state.Current
	for i, row := range previous {
		for j, cell := range row {
			aliveNeighbors := countAliveNeighbors(i, j, previous)
			if cell {
				// Cell is alive
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					// Dies or under/overpopulation
					state.Current[i][j] = false
				}
			} else if aliveNeighbors == 3 {
				// Cell was dead but is now alive through reproduction
				state.Current[i][j] = true
			}
		}
	}
	state.Generation++
	return state
}

func countAliveNeighbors(i int, j int, state [NumRows][NumColumns]bool) int {
	total := 0
	if i >= 1 {
		// Neighbors above
		if j >= 1 && state[i-1][j-1] {
			total++
		}
		if state[i-1][j] {
			total++
		}
		if j < NumColumns-1 && state[i-1][j+1] {
			total++
		}
	}
	// Neighbors on the same row
	if j >= 1 && state[i][j-1] {
		total++
	}
	if j < NumColumns-1 && state[i][j+1] {
		total++
	}
	// Neighbors below
	if i < NumRows-1 {

		if j >= 1 && state[i+1][j-1] {
			total++
		}
		if state[i+1][j] {
			total++
		}
		if j < NumColumns-1 && state[i+1][j+1] {
			total++
		}
	}
	return total
}
