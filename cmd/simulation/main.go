package main

import (
	"github.com/eneskzlcn/manifacturing-shop-simulation/internal/simulation"
)

func main() {
	simulation := simulation.NewSimulation()
	simulation.Start(100, 2, 10, 10)
}
