package main

import (
   engine "quoridor/q_engine"
)

func main() {
	var e engine.Engine
	e= engine.New()
	e.Spiel()
}

