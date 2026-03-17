package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
)

// суть игры:
// - поле 100 клеток
// - два игрока - поочереди кидают рандом на от 1 до 6
// - если они находятся на соседних клетках - выбор "драться или нет"
// - если на одной клетке - драка не избежна
// - режим драки - по очереди бросают кубик и если выйград - на 1 клетку вперед. Проиграл - на две назад
// - каждая кратная 10 клетка - плюс 2 вперед. Каждая кратная 5, но не 10 клетка - 1 назад
// - Победил тот, кто дошел до конца первый.

// - запускает генерацию поля
func Pole(p1, p2 int) {
	for i := 0; i < 101; i++ {

		if i == p1 && i == p2 {
			fmt.Print("|p1/2")
		} else if i == p1 {
			fmt.Print("|-p1", "-")
		} else if i == p2 {
			fmt.Print("|-p2", "-")
		} else {
			fmt.Print("| ", i, " ")
		}
		if i%10 == 0 {
			fmt.Println("")
		}
		if i > 1 && i < 10 {
			fmt.Print(" ")
		}
		if i == 100 {
			fmt.Print("\t\t      Finish \n")
		}
	}
}

// - проверка на кратность 5 и 10
func CheckFiveTen(pole int) int {
	var dop int
	if pole%10 == 0 {
		dop = 2
	} else if pole%5 == 0 {
		dop = -1
	} else {
		dop = 0
	}
	return dop
}

// - проверка на победу
func СheckWin(p1, p2 int) bool {
	if p1 > 100 || p2 > 100 {
		return true
	}
	return false
}

func PrintResultForCheckFiveSeven(dop int) {
	if dop == 2 {
		fmt.Print("Ты попал в Десяточку: Плюс 2 хода\n")
	} else if dop == -1 {
		fmt.Print("Ты попал в Пятерку: Минус 1 ход\n")
	}
}

func PrintWiner(count int) {
	if count%2 == 0 {
		fmt.Print("Победил Player1\nПОЗДРАВЛЯЮ!")
	} else {
		fmt.Print("Победил Player2\nПОЗДРАВЛЯЮ!")
	}
}

// - рандомим число от 1 до 6
// - рекурсим если 0
func NewCryptoRand() int {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(7))
	if err != nil {
		panic(err)
	}
	if safeNum.Int64() != 0 {
		num := int(safeNum.Int64())
		return num
	}
	return NewCryptoRand()
}

func Throw(cube int) int {
	var tap string
	fmt.Println("Бросай кубик!\nPress Enter")
	fmt.Scanln(&tap)
	cube = NewCryptoRand()
	fmt.Print("Выпало - ", cube, "\n")
	return cube
}

func CheckForFight(p1, p2 int) bool {
	count := p1 - p2
	return count == 0 || count == 1 || count == -1
}

func Fight(p1, p2 int) (int, int) {
	cube1, cube2 := 0, 0
	cube1 = NewCryptoRand()
	cube2 = NewCryptoRand()
	fmt.Println("Первый - ", cube1, " Второй - ", cube2)

	if cube1 > cube2 {
		fmt.Print("\nPlayer1 победил!\n")
		p1 += 2
		p2 -= 1
	} else if cube2 > cube1 {
		fmt.Print("\nPlayer2 победил!\n")
		p2 += 2
		p1 -= 1
	} else {
		fmt.Print("\nНичья!\n")
	}

	fmt.Print("\nРасклад после драки:\n")
	showPole(p1, p2)

	return p1, p2
}

// - режим драки
func checkFight(p1, p2 int) bool {
	count := p1 - p2
	chouse := ""

	if count == 1 || count == -1 {
		fmt.Print("\nХотите драться? Y - да\n")
		fmt.Scanln(&chouse)
		if chouse == "y" || chouse == "н" || chouse == "Н" || chouse == "Y" {
			return true
		}
	} else if count == 0 {
		fmt.Print("\nДрака неизбежна\n")
		return true
	}
	return false
}
