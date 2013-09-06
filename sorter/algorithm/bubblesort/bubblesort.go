package bubblesort

import "fmt"

func BubbleSort(values []int) {
	valuesLen := len(values)
	for i := 0; i < valuesLen-1; i++ {
		for j := 0; j < valuesLen-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	fmt.Println("The bubblesort complete!")
}
