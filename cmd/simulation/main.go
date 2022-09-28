package main

import (
	"github.com/eneskzlcn/manifacturing-shop-simulation/internal/simulation"
)

func main() {
	manifacturingShopSimulation := simulation.New()
	manifacturingShopSimulation.Start(simulation.Properties{
		MinExamineTime:               2,
		MaxExamineTime:               10,
		TerminateCounter:             100,
		FailurePossibilityPercentage: 10,
		PartTurnOutRate:              5,
	})
}
