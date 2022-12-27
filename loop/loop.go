package loop


func SuperAdd(numbers ...int) int {
	// for in, for each
	total := 0
	for _, number := range numbers {
		total += number
	}

	// for i:=0; i < len(numbers); i++ {
	// 	fmt.Println(i, numbers[i])
	// }

	return total
}