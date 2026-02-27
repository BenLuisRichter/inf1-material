package seats

type Reservation struct {
	id          string
	origin      string
	destination string
}

func (r Reservation) Occupied(current_station int, route []string) bool {
	// Hinweis:
	// Bestimme, ob die aktuelle Station
	// zwischen r.origin und r.destination liegt.
	// TODO
	return false
}

func (r Reservation) Text(station int, route []string) string {
	// TODO
	return "ggf. freigeben"
}
