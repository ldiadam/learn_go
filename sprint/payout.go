package sprint

func Payout(amount int, denominations []int) (payout []int) {
	for i := 0; i < len(denominations); i++ {
		for j := 0; j < len(denominations)-i-1; j++ {
			if denominations[j] < denominations[j+1] {
				denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
			}
		}
	}

	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}
	if amount > 0 {
		return []int{}
	}

	return payout
}
