package main

import (
	"math"
)

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

func simulate(games []Game) (string, int) {
	moves := []string{UP, DOWN, LEFT, RIGHT}

	var bestMove string
	bestScore := math.MinInt
	for i := range moves {
		move := moves[i]

		gameCopies := make([]Game, len(games))
		for i, game := range games {
			gameCopies[i] = game.Copy()
		}
		var results = make([]ActionResult, len(gameCopies))
		for _, game := range gameCopies {
			results = append(results, game.applyMove(move))
		}
		score := 0
		for _, result := range results {
			score += result.score
		}
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove, bestScore
}
