package main

import (
	"testing"
)

func TestHurdleSim(t *testing.T) {

	field := "....#....#....#....#....#....."

	game := &HurdleGame{playerIndex: 0}
	game.agents = append(game.agents, &Agent{index: 0, position: 0, stunTime: 0, score: 0, bronze: 0, silver: 0, gold: 0})
	game.updateField(field)

	games := make([]Game, 0)
	games = append(games, game)

	move, score := simulate(games)
	if move != RIGHT {
		t.Errorf("Expected move RIGHT, got %s", move)
	}
	if score != 3 {
		t.Errorf("Expected score 3, got %d", score)
	}

	// Update position
	game.agents[0].position += 3

	move, score = simulate(games)

	if move != UP {
		t.Errorf("Expected move UP, got %s", move)
	}

	field = ".....#....#...#....#....#....."
	game.updateField(field)
	game.agents[0].position = 0
	move, score = simulate(games)
	if move != RIGHT {
		t.Errorf("Expected move RIGHT, got %s", move)
	}
	if score != 3 {
		t.Errorf("Expected score 3, got %d", score)
	}

	game.agents[0].position = 3
	move, score = simulate(games)
	if move != LEFT {
		t.Errorf("Expected move LEFT, got %s", move)
	}
}
