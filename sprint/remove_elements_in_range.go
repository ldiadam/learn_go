package sprint

func RemoveElementsInRange(arr []float64, lowerIndex int, upperIndex int) []float64 {
	//Revese if lower and upper not correct
	if lowerIndex > upperIndex {
		lowerIndex, upperIndex = upperIndex, lowerIndex
	}
	result := []float64{}

	//loop through range
	for index, i := range arr {
		//if index is lower that lowerIndex (not higher) OR if index is higher and same as upperIndex (lower)
		//its basically in between
		if index < lowerIndex || index >= upperIndex {
			// add to result
			result = append(result, i)
		}
	}

	return result
}
