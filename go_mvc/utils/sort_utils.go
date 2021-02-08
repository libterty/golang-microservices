package utils

func BubbleSort(elements []int)  []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements) - 1; i++ {
			if elements[i] > elements[i+1] {
				temp := elements[i]
				elements[i] = elements[i+1]
				elements[i+1] = temp
				keepRunning = true
			}
		}
	}
	return  elements
}