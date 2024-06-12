package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var miniGame []Game

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var playerIdx int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &playerIdx)
	var nbGames int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &nbGames)
	miniGame = make([]Game, 0)
	for i := 0; i < nbGames; i++ {
		switch i {
		case 0:
			game := &HurdleGame{playerIndex: playerIdx}
			game.agents = make([]*Agent, 3)
			for i := 0; i < 3; i++ {
				game.agents[i] = &Agent{index: i, position: 0}
			}
			miniGame = append(miniGame, game)
			break
		}
	}
	log("Starting game", len(miniGame))
	for {
		for i := 0; i < 3; i++ {
			scanner.Scan()
			scoreInfo := scanner.Text()
			_ = scoreInfo // to avoid unused error
			for _, game := range miniGame {
				game.updateAgentScore(i, scoreInfo)
			}
		}
		for i := 0; i < nbGames; i++ {
			var gpu string
			var reg0, reg1, reg2, reg3, reg4, reg5, reg6 int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &gpu, &reg0, &reg1, &reg2, &reg3, &reg4, &reg5, &reg6)
			for idx := range miniGame {
				if idx == i {
					game := miniGame[idx]
					game.updateAgent(reg0, reg1, reg2, reg3, reg4, reg5, reg6)
					game.updateField(gpu)
				}
			}
		}

		s, i := simulate(miniGame)
		fmt.Println(fmt.Sprintf("%s score=%d", s, i)) // Write action to stdout
	}
}
