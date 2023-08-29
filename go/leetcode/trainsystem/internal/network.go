package internal

type Network struct {
	Stations []*Station
}

func NewNetwork() *Network {
	return &Network{}
}

func (n *Network) CheckIn(id int, stationName string, t int) error {
	var s *Station
	for _, station := range n.Stations {
		if station.Name != stationName {
			continue
		}

		s = station
		break
	}

	if s == nil {
		s = NewStation(stationName)
		n.Stations = append(n.Stations, s)
	}

	return s.checkIn(NewPassenger(id), t)
}

func (n *Network) CheckOut(id int, stationName string, t int) error {
	var s *Station
	for _, station := range n.Stations {
		if station.Name != stationName {
			continue
		}

		s = station
		break
	}

	if s == nil {
		s = NewStation(stationName)
		n.Stations = append(n.Stations, s)
	}

	return s.checkOut(NewPassenger(id), t)
}

func (n *Network) GetAverageTime(from Station, to Station) float64 {
	for _, station := range n.Stations {
		if station.Name != from.Name {
			continue
		}

		for _, route := range station.Routes {
			if route.To.Name != to.Name {
				continue
			}

			return route.AverageTime
		}
	}

	return -1
}
