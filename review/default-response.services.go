package review

func IncrementStarInReturnData(returnData *DefaultResponse, star int) *DefaultResponse {
	if star == 1 {
		returnData.OneStar++
	}
	if star == 2 {
		returnData.TwoStar++
	}
	if star == 3 {
		returnData.ThreeStar++
	}
	if star == 4 {
		returnData.FourStar++
	}
	if star == 5 {
		returnData.FiveStar++
	}
	return returnData
}
