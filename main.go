package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	playground := [4][4]int{}
	playground = addRandomNumber(playground)
	playground = addRandomNumber(playground)
	render(playground)
	haveMoves := true
	for haveMoves {
		key, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(key))
		switch string(key) {
		case "ц", "w":
			playground = mvUp(playground)
		case "ф", "a":
			playground = mvLeft(playground)
		case "ы", "s":
			playground = mvDown(playground)
		case "в", "d":
			playground = mvRight(playground)
		case "з", "p":
			fmt.Print("\033[H\033[2J")
			fmt.Println("Quit.")
			return
		}
	}
}

func render(playground [4][4]int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Render:")
	fmt.Println(playground[0])
	fmt.Println(playground[1])
	fmt.Println(playground[2])
	fmt.Println(playground[3])
}

func mvUp(pg [4][4]int) [4][4]int {
	for col := 0; col < 4; col++ {
		prevT := 0
		for i := 0; i < 4; i++ {
			for row := 1; row <= 3; row++ {
				t := pg[row][col]
				if t == 0 {
					continue
				}
				if pg[row-1][col] == t && prevT*4 != t*2 {
					pg[row-1][col] *= 2
					pg[row][col] = 0
					prevT = t

				} else if pg[row-1][col] == 0 {
					pg[row-1][col] = t
					pg[row][col] = 0
				}
			}
		}
	}
	res := addRandomNumber(pg)
	render(res)
	fmt.Println("mv up")
	return res
}

func mvDown(pg [4][4]int) [4][4]int {
	for col := 0; col < 4; col++ {
		prevT := 0

		for i := 0; i < 4; i++ {
			for row := 2; row >= 0; row-- {
				t := pg[row][col]
				if t == 0 {
					continue
				}
				if pg[row+1][col] == t && prevT*4 != t*2 {
					pg[row+1][col] *= 2
					pg[row][col] = 0
					prevT = t
				} else if pg[row+1][col] == 0 {
					pg[row+1][col] = t
					pg[row][col] = 0
				}
			}
		}
	}
	res := addRandomNumber(pg)
	render(res)
	fmt.Println("mv down")
	return res
}

func mvRight(pg [4][4]int) [4][4]int {
	for row := 0; row < 4; row++ {
		prevT := 0
		for i := 0; i < 4; i++ {
			a
			for col := 2; col >= 0; col-- {
				t := pg[row][col]
				if t == 0 {
					continue
				}
				if pg[row][col+1] == t && prevT*4 != t*2 {
					pg[row][col+1] *= 2
					pg[row][col] = 0
					prevT = t
				} else if pg[row][col+1] == 0 {
					pg[row][col+1] = t
					pg[row][col] = 0
				}
			}
		}
	}
	res := addRandomNumber(pg)
	render(res)
	fmt.Println("mv right")
	return res
}

func mvLeft(pg [4][4]int) [4][4]int {
	for row := 0; row < 4; row++ {
		prevT := 0
		for i := 0; i < 4; i++ {
			for col := 1; col <= 3; col++ {
				a
				t := pg[row][col]
				if t == 0 {
					continue
				}
				if pg[row][col-1] == t && prevT*4 != t*2 {
					pg[row][col-1] *= 2
					pg[row][col] = 0
					prevT = t
				} else if pg[row][col-1] == 0 {
					pg[row][col-1] = t
					pg[row][col] = 0
				}
			}
		}
	}
	res := addRandomNumber(pg)
	render(res)
	fmt.Println("mv left")
	return res
}

func addRandomNumber(pg [4][4]int) [4][4]int {
	rand.Seed(time.Now().UnixMicro())
	rnd := rand.Intn(16)
	nb := 0
	for nb != 2 && nb != 4 {
		nb = rand.Intn(5)
	}
	counter := 0
	for i := 0; i < len(pg); i++ {
		for k := 0; k < len(pg[i]); k++ {
			counter++
			if counter == rnd && pg[i][k] == 0 {
				pg[i][k] = nb
				return pg
			}
		}
	}
	return addRandomNumber(pg)
}

// func checkMoves(s []byte) bool {
// 	var pg [4][4]int
// 	err := json.Unmarshal(s, &pg)
// 	if err != nil {
// 		return true
// 	}
// 	if pg == mvUp(pg) && pg == mvLeft(pg) && pg == mvRight(pg) && pg == mvDown(pg) {
// 		return false
// 	}
// 	return true
// }
