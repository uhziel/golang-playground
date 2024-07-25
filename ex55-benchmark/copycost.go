package copycost

var arr1 = make([][10]int, 1000)
var arr2 = make([][10]int, 1000)
var arr3 = make([][10]int, 1000)

var sum1, sum2, sum3 = 0, 0, 0

func copyForI() int {
	for i := 0; i < len(arr1); i++ {
		sum1 += arr1[i][0]
	}
	return sum1
}

func copyForR1() int {
	for i := range arr2 {
		sum2 += arr2[i][0]
	}
	return sum2
}

func copyForR2() int {
	for _, v := range arr3 {
		sum3 += v[0]
	}
	return sum3
}
