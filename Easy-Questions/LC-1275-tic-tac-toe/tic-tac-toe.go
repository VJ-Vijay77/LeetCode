package main

import "fmt"

func main() {
	moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	// Output: "A"
	// Explanation: A wins, they always play first.
	fmt.Println("Result = ",tictactoe(moves))
}


func tictactoe(moves [][]int) string {
    grid := [3][3]string{}
	player := "A"

	for _, move := range moves{
		row,col := move[0],move[1]
		grid[row][col] = player
		if checkWinner(row,col,player,grid){
			return player
		}

		if player == "A" {
			player = "B"
		}else{
			player = "A"
		}
		
	}

	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}


func checkWinner(row int,col int,player string,grid [3][3]string) bool{

	//check row
	if grid[row][0] == player && grid[row][1] == player && grid[row][2] == player {
		return true
	}

	//check column
	if grid[0][col] == player && grid[1][col] == player && grid[2][col] == player {
		return true
	}

	//check first diagonal
	if col == row && grid[0][0] == player && grid[1][1] == player && grid[2][2] == player {
		return true
	}

	//check second diagonal
	if col+row == 2 && grid[0][2] == player && grid[1][1] == player && grid[2][0] == player{
		return true
	}

	return false
}