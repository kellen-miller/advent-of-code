package internal

import (
	"fmt"
)

type Station struct {
	Name       string
	Passengers []*Passenger
	Routes     []*Route
}

func NewStation(name string) *Station {
	return &Station{
		Name: name,
	}
}

func (s *Station) checkIn(p *Passenger, t int) error {
	if err := p.checkIn(s, t); err != nil {
		return err
	}

	s.Passengers = append(s.Passengers, p)
	return nil
}

func (s *Station) checkOut(p *Passenger, t int) error {
	for i, passenger := range s.Passengers {
		if passenger.ID != p.ID {
			continue
		}

		s.updateAverageTime(p, t)
		s.Passengers = append(s.Passengers[:i], s.Passengers[i+1:]...)

		return p.checkOut()
	}

	return fmt.Errorf("%s is not checked in at %s", p, s.Name)
}

func (s *Station) updateAverageTime(p *Passenger, t int) {
	var r *Route
	for _, route := range s.Routes {
		if route.From.Name != p.Station.Name || route.To.Name != s.Name {
			continue
		}

		r = route
		break
	}

	if r == nil {
		r = &Route{
			From: p.Station,
			To:   s,
		}
	}

	r.updateAverageTime(t - p.CheckedInAt)
}
