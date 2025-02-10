package sieve

import (
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

type primeFinder struct{}

func NewSieve() Sieve {
	return &primeFinder{}
}

func (s *primeFinder) NthPrime(n int64) int64 {
	if n == 0 {
		return 2
	}

	limit := getUpperBound(n)
	sieve := make([]bool, limit)
	sieve[0], sieve[1] = true, true

	var p int64 = 2
	for p*p < limit {
		if !sieve[p] {
			for i := p * p; i < limit; i += p {
				sieve[i] = true
			}
		}
		p++
	}

	var count int64 = 0
	for i := int64(2); i < limit; i++ {
		if !sieve[i] {
			if count == n {
				return i
			}
			count++
		}
	}

	return -1
}

func getUpperBound(n int64) int64 {
	if n < 6 {
		return 13
	}

	nFloat := float64(n)
	return int64(nFloat * (math.Log(nFloat) + math.Log(math.Log(nFloat))) * 1.1)
}
