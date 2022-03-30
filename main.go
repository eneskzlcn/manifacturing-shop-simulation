package main

import (
	"ManifacturingShopSimulation/manifacturing-shop-simulation"
)

func main() {
	simulation := manifacturing_shop_simulation.NewSimulation()
	simulation.Start(100,2,10,10)
}
