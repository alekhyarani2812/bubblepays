package learning

import "log"

func ChangeCoin(numbers []int, target int) {
	//int{2, 5, 3, 6} 
	for i, num := range numbers {
		//num := i * num
		num = num + num
		if num == target{
			var arr []int
			arr = append(arr, numbers[i])
			log.Println("arr :", arr)
		}
	}
}



