package persona_entities

type Status uint

const (
	AB Status = iota
	NU
	RS
	WK
	RP
)

func (s Status) String() string {
	switch s {
	case AB:
		return "ab"
	case NU:
		return "nu"
	case RS:
		return "rs"
	case WK:
		return "wk"
	case RP:
		return "rp"
	}
	return "unknown"
}

type Elem struct {
	ID int

	PersonaID int
	Persona   Persona
	ElementID int
	Element   Element

	Status Status
}

func (elem *Elem) StatusName() string {
	return elem.Status.String()
}

func (elem *Elem) IsAB() bool {
	return elem.Status == AB
}

func (elem *Elem) IsNU() bool {
	return elem.Status == NU
}

func (elem *Elem) IsRS() bool {
	return elem.Status == RS
}

func (elem *Elem) IsWK() bool {
	return elem.Status == WK
}

func (elem *Elem) IsRP() bool {
	return elem.Status == WK
}
