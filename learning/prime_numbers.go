package learning

import "log"

func Prime(){
	max := 10
	if max < 2 {
		log.Println("This is not a prime")
	}  
	isPrime := make([]bool, max+1)
    for i := 2; i <= max; i++ {
        isPrime[i] = true
    }

    // Implement the Sieve of Eratosthenes
    for p := 2; p*p <= max; p++ {
        if isPrime[p] {
            for i := p * p; i <= max; i += p {
                isPrime[i] = false
            }
        }
    }

    // Collect all prime numbers
    primes := []int{}
    for p := 2; p <= max; p++ {
        if isPrime[p] {
            primes = append(primes, p)
        }
    }
	log.Println("Prime numbers between 0 and", max, "are:", primes)
}