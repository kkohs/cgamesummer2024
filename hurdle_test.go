package main

import (
	"testing"
)

func TestHurdleMove(t *testing.T) {

	field := "....#....#....#....#....#....."

	game := &HurdleGame{playerIndex: 0}
	game.agents = append(game.agents, &Agent{index: 0, position: 0, stunTime: 0, score: 0, bronze: 0, silver: 0, gold: 0})
	game.updateField(field)

	// Test Right
	result := game.applyMove(RIGHT)
	if result.score != 3 {
		t.Errorf("Expected score 3, got %d", result.score)
	}
	// Update position
	game.agents[0].position += 3

	// Test Right
	result = game.applyMove(RIGHT)
	if result.score != -3 {
		t.Errorf("Expected score -3, got %d", result.score)
	}

	// Test out of bounds
	game.agents[0].position = len(field) - 1
	result = game.applyMove(RIGHT)
	if result.score != 3 {
		t.Errorf("Expected score 3, got %d", result.score)
	}

	// test empty track
	game.track = ""
	result = game.applyMove(RIGHT)
	if result.score != 0 {
		t.Errorf("Expected score 0, got %d", result.score)
	}
}
