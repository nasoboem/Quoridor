package q_spieler

import (
	"fmt"
)

type data struct {
	name string
	nr int
}

func New() *data {
		var s *data
		s = new(data)
		return s
}

func (s *data) SetzeNummer (nr int) {
	s.nr = nr
}

func (s *data) SetzeName(name string) {
	s.name = name
}

func (s *data) GebeNummer () int {
	return s.nr
}

func (s *data) GebeName () string {
	return s.name
}

func (s *data) String () string {
	var erg string
	erg = erg + fmt.Sprintln ("Name:", s.name)
	return erg
}

