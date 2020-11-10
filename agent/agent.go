package agent

import "math/rand"

type Agent struct {
	Epsilon        float64
	Alpha          float64
	MinExplore     float64
	ProfitLoss     int
	ProfitLossHist []int
}

func (a *Agent) NewEpsilon() float64 {
	var alpha = 1.0 - a.Alpha
	return alpha * a.Epsilon
}

func (a *Agent) Bet(nTrial int) bool {
	if nTrial <= a.MinExplore {
		if rand.Float64() > 0.5 {
			return true
		}
		return false
	} else if rand.Float64() <= a.Epsilon {
		a.Epsilon = a.NewEpsilon()
	}
}
