package learning

import (
	"fmt"
	"log"
)

func ChangeCoin(numbers []int, target int) {
	//  int{2, 5, 3, 6}   target =10
	sum := 0
	var arr []int
	for _, num := range numbers {
		sum = sum + num
		arr = append(arr, num)
		if sum == target {
			log.Println("arr :", arr)
			return
		}
	}
}

func CoinChange(target int, coins []int) int {
	// Create a dp array to store the number of ways to make change for each amount
	dp := make([]int, target+1)

	// Base case: There is one way to make change for 0 amount
	dp[0] = 1

	// Iterate over each coin
	for _, coin := range coins {
		// Update the dp array for each amount from coin to n
		for amount := coin; amount <= target; amount++ {
			dp[amount] = dp[amount] + dp[amount-coin]
			//dp[amount] += dp[amount-coin]
			log.Println("dp[amount]: ", dp[amount], " amount: ", amount,
				"dp[amount-coin]: ", dp[amount-coin], "amount-coin: ", amount-coin, 
				"coin:", coin)
		}
		log.Println("coins : ", coins)
	}
	// The answer is the number of ways to make change for amount n
	return dp[target]
}

func CountCombinations(coins []int, n, sum int) int{
    if sum == 0 {
        return 1
    }
    if sum < 0 || n <= 0 {
        return 0
    }
	fmt.Println("CountCombinations(coins, n, sum-coins[n-1]) :",CountCombinations(coins, n, sum-coins[n-1]))
	fmt.Println("CountCombinations(coins, n-1, sum) :",CountCombinations(coins, n-1, sum))
    // Include the current coin or exclude it
    return CountCombinations(coins, n, sum-coins[n-1]) + CountCombinations(coins, n-1, sum)
}

