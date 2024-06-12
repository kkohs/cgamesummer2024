package main

type HurdleGame struct {
	track       string
	playerIndex int
	agents      []*Agent
}

func (g *HurdleGame) applyMove(move string) ActionResult {

	if g.track == "" {
		return ActionResult{0, false}
	}

	score := 0

	if g.agents[g.playerIndex].stunTime > 0 {
		g.agents[g.playerIndex].stunTime--
		return ActionResult{0, true}
	}

	switch move {
	case UP:
		score += 2
		nextHurdleIndex := g.getFirstHurdleIndex(g.agents[g.playerIndex].position, g.agents[g.playerIndex].position+2)
		if nextHurdleIndex == g.agents[g.playerIndex].position+1 {
			score += 3
		} else if nextHurdleIndex == g.agents[g.playerIndex].position+2 {
			score = -3
		}
		break
	case DOWN:
		score += 2
		if g.isHurdleInPath(g.agents[g.playerIndex].position, g.agents[g.playerIndex].position+2) {
			score = -3
		}
		break
	case LEFT:
		score = 1
		if g.isHurdleInPath(g.agents[g.playerIndex].position, g.agents[g.playerIndex].position+1) {
			score = -3
		}
		break
	case RIGHT:
		score = 3
		if g.isHurdleInPath(g.agents[g.playerIndex].position, g.agents[g.playerIndex].position+3) {
			score = -3
		}
		break
	}

	if score > 0 {
		return ActionResult{score, false}
	}
	return ActionResult{score: score, stunned: true}

}

func (g *HurdleGame) updateAgentScore(agentIndex int, scoreInfo string) {
	g.agents[agentIndex].score = int(scoreInfo[0])
	g.agents[agentIndex].bronze = int(scoreInfo[1])
	g.agents[agentIndex].silver = int(scoreInfo[2])
	g.agents[agentIndex].gold = int(scoreInfo[3])
}

func (g *HurdleGame) updateAgent(reg0, reg1, reg2, reg3, reg4, reg5, reg6 int) {
	g.agents[0].position = reg0
	g.agents[0].stunTime = reg3
	g.agents[1].position = reg1
	g.agents[1].stunTime = reg4
	g.agents[2].position = reg2
	g.agents[2].stunTime = reg5

	log("Agent 0: ", g.agents[0].position, g.agents[0].stunTime)
}

func (g *HurdleGame) updateField(gpu string) {
	log(gpu)
	if gpu == "GAME_OVER" {
		g.track = ""
	} else {
		g.track = gpu
	}
}

func (g *HurdleGame) isHurdleInPath(startPos, endPos int) bool {
	for i := startPos; i <= endPos; i++ {
		if i >= len(g.track) {
			return false
		}
		if g.track[i] == Hurdle {
			return true
		}
	}
	return false
}

func (g *HurdleGame) getFirstHurdleIndex(start, end int) int {
	for i := start; i <= end; i++ {
		if i >= len(g.track) {
			return -1
		}
		if g.track[i] == Hurdle {
			return i
		}
	}
	return -1
}

func (g *HurdleGame) Copy() Game {
	newGame := HurdleGame{}
	newGame.track = g.track
	newGame.playerIndex = g.playerIndex
	newGame.agents = make([]*Agent, len(g.agents))
	for i, agent := range g.agents {
		newGame.agents[i] = agent.copy()
	}
	return &newGame
}
