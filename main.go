package main

import "ManifacturingShopSimulation/simulation"

func main() {
	simulation := simulation.NewSimulation()
	simulation.Start(100,2,10,10)
}
