package Array

func ArraySum(numberArray [5]int)int{
	sum := 0
	for _,number := range numberArray{
		sum += number
	}
	return sum

}

func SliceSum(numerArray []int)int{
	sum := 0
	for _, number := range numerArray{
		sum += number
	}
	return sum
}

func sumAll(numbersToSum ...[]int)[]int{
	var returnSlice []int
	for _,numbers := range numbersToSum{
		returnSlice = append(returnSlice, SliceSum(numbers))
	}
	
	return returnSlice
}
func sumAllTails(numbersToSum ...[]int)[]int{
	var returnSlice []int
	for _,numbers := range numbersToSum{
		returnSlice = append(returnSlice, SliceSum(numbers[1:]))
	}
	
	return returnSlice
}
