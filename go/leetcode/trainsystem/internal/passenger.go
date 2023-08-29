package internal

import (
	"fmt"
)

type Passenger struct {
	ID          int
	Name        string
	Station     *Station
	CheckedInAt int
}

func NewPassenger(id int) *Passenger {
	return &Passenger{
		ID: id,
	}
}

func (p *Passenger) checkIn(s *Station, t int) error {
	if p.Station != nil {
		return fmt.Errorf("%s is already checked in at %s", p, p.Station.Name)
	}

	p.Station = s
	p.CheckedInAt = t
	return nil
}

func (p *Passenger) checkOut() error {
	if p.Station == nil {
		return fmt.Errorf("%s is not checked in", p)
	}

	p.Station = nil
	p.CheckedInAt = -1
	return nil
}

func (p *Passenger) String() string {
	return fmt.Sprintf("Passenger %d: %s", p.ID, p.Name)
}
