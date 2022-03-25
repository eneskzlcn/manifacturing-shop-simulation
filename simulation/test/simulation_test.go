package test

import (
	"ManifacturingShopSimulation/simulation"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRandomFunction(t *testing.T) {
	for i:= 0 ; i < 200; i++ {
		random := simulation.Random(0, 100)
		assert.GreaterOrEqual(t, random,0)
		assert.LessOrEqual(t, random,100)
		log.Printf("\t%d", random)
	}
}
func TestSimulation(t *testing.T) {
	simulation := simulation.NewSimulation()
	simulation.Start(100,2,10,10)
}