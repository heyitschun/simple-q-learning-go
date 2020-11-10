package agent

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
