package internal

type Route struct {
	From        *Station
	To          *Station
	AverageTime float64
}

func (r *Route) updateAverageTime(t int) {
	r.AverageTime = (r.AverageTime + float64(t)) / 2
}
