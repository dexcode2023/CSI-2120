//Name: Dexter Day
//Student number: 300192263

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Intn(6) + 5

	treasureRow, treasureCol := r.Intn(n), r.Intn(n)

	player1Row, player1Col := 0, 0
	player1Moves := 0
	player1Found := false

	player2Row, player2Col := n-1, n-1
	player2Moves := 0
	player2Found := false

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for {

		if !player1Found {
			dir := directions[r.Intn(4)]
			newRow, newCol := player1Row+dir[0], player1Col+dir[1]
			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < n {
				player1Row, player1Col = newRow, newCol
				player1Moves++
			}
			if player1Row == treasureRow && player1Col == treasureCol {
				fmt.Printf("Player 1 found the treasure after %d moves.\n", player1Moves)
				player1Found = true
			}
		}

		if !player2Found {
			dir := directions[r.Intn(4)]
			newRow, newCol := player2Row+dir[0], player2Col+dir[1]
			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < n {
				player2Row, player2Col = newRow, newCol
				player2Moves++
			}
			if player2Row == treasureRow && player2Col == treasureCol {
				fmt.Printf("Player 2 found the treasure after %d moves.\n", player2Moves)
				player2Found = true
			}
		}
		if player1Found && player2Found {
			break
		}

	}
}
