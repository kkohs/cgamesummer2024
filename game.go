package main

const (
	Empty  = '.'
	Hurdle = '#'
)

type Game interface {
	applyMove(move string) ActionResult
	updateAgentScore(agentIndex int, scoreInfo string)
	updateAgent(reg0, reg1, reg2, reg3, reg4, reg5, reg6 int)
	updateField(gpu string)
	Copy() Game
}

type ActionResult struct {
	score   int
	stunned bool
}

type Agent struct {
	index, position, stunTime, score, bronze, silver, gold int
}

// copy agent
func (a *Agent) copy() *Agent {
	newAgent := Agent{}
	newAgent.index = a.index
	newAgent.position = a.position
	newAgent.stunTime = a.stunTime
	newAgent.score = a.score
	newAgent.bronze = a.bronze
	newAgent.silver = a.silver
	newAgent.gold = a.gold
	return &newAgent

}
