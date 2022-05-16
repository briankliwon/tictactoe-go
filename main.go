package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
)

type Point struct {
	player    string
	indicator [9]string
	// playingCap int
}

func main() {
	var game Point
	game.player = "0"
	gameOver := false

	// var winner string

	for !gameOver {
		generateBoard(game.indicator)
		inputIdx := askforplay()
		game.setInput(inputIdx)
		win, val := game.checkWinner()
		log.Println(win, "a")
		log.Println(val, "b")
	}
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func generateBoard(idx [9]string) {
	// clearScreen()
	for i, v := range idx {
		if v == "" {
			fmt.Print("* ")
		} else {
			fmt.Print(v + " ")
		}

		if (i+1)%3 == 0 {
			fmt.Print("\n")
		}
	}
}

func askforplay() (moveInt int) {
	fmt.Println("Enter Position to play: ")
	fmt.Scan(&moveInt)
	return moveInt
}

func (p *Point) setInput(input int) {
	if !(input >= 9) {
		if p.indicator[input] == "" {
			p.indicator[input-1] = "O"
		}
	}
}

func (p *Point) checkWinner() (winner bool, key string) {
	var total int
	var value string

	for i := 0; i < 9; i++ {
		if !winner && p.indicator[i] != "" {
			total = 0
			value = ""
			lastValueH := i - 3
			lastValueV := i - 1

			for l := 0; l < 3; l++ {
				if lastValueH >= 0 && p.indicator[i] == p.indicator[lastValueH] {
					total++
					lastValueH -= 3
					value = p.indicator[i]
				}
				if lastValueV >= 0 && p.indicator[i] == p.indicator[lastValueV] {
					total++
					lastValueV -= 1
					value = p.indicator[i]
				}

			}

			if total == 2 {
				winner = true
				continue
			}

		}
	}
	p.getCrossLineCheck()

	return winner, value
}

func (p *Point) getCrossLineCheck() {
	log.Println("hI JON")
	lenIndicator := len(p.indicator)
	square := math.Sqrt(float64(lenIndicator))
	var lr []int
	var rl []int

	if lenIndicator%2 != 0 {
		var list []int
		for i := 0; i < lenIndicator; i++ {
			if (i+1)%int(square) != 0 {
				list = append(list, i)
			} else {
				log.Println(list, "IKI KANG")
				if list[0] == 0 {
					lr = append(lr, list[0])
					rl = append(rl, i)
					list = make([]int, 0)
				} else {
					lr = append(lr, len(list))
					rl = append(rl, i)
					list = make([]int, 0)
				}
			}
			log.Println(list)
			log.Println(lr, "lr")
			log.Println(rl)
		}
	}
}
