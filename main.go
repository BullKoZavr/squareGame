package main

import (
	"fmt"
)

func player1Throu(player1 int) int {
	cube := 0
	fmt.Println("\n\n-- Ход Player1 --")
	cube = Throw(cube)
	player1 += cube
	dop := CheckFiveTen(player1)
	PrintResultForCheckFiveSeven(dop)
	player1 += dop
	fmt.Print("Player1 на клетке - ", player1)

	return player1
}

func player2Throw(player2 int) int {
	cube := 0
	fmt.Println("\n\n-- Ход Player2 --")
	cube = Throw(cube)
	player2 += cube
	dop := CheckFiveTen(player2)
	player2 += dop
	fmt.Print("Player2 на клетке - ", player2)

	return player2
}
func showPole(p1, p2 int) {
	var tap string
	fmt.Scanln(&tap)
	Pole(p1, p2)
}

func main() {

	player1, player2 := 0, 0
	count := 1
	for !СheckWin(player1, player2) {
		if count%2 == 0 {
			player2 = player2Throw(player2)
			showPole(player1, player2)
			count++
		} else {
			player1 = player1Throu(player1)
			count++
		}
		if checkFight(player1, player2) {
			player1, player2 = Fight(player1, player2)
		}
	}
	PrintWiner(count)

}
