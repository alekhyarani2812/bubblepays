package learning

import (
	"fmt"
	"testing"
)

func TestChangeCoin(t *testing.T) {
	change := []int{2, 5, 3, 6}
	target := 10
	ChangeCoin(change, target)
}

func TestCoinChange(t *testing.T) {
	target := 10
	change := []int{2, 5, 3, 6}
	result := CoinChange(target, change)
	fmt.Printf("Number of ways to make change for %d using coins %v: %d\n",
		target, change, result)
}

func TestCountCombinations(t *testing.T) {
	target := 10
	coins := []int{2, 5, 3, 6}
	numCombinations := CountCombinations(coins, len(coins), target)
	fmt.Printf("Number of combinations: %d\n", numCombinations)
}
