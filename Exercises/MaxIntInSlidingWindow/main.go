import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	result := []int{}

	max := findMax(arr[:k])

	result = append(result, max)

	for i := 1; i <= len(arr)-k; i++ {

		if arr[i+2] > max {

			max = arr[i+2]

		}

		result = append(result, max)

	}

	fmt.Println(result)
}

func findMax(arr []int) int {

	var m int = math.MinInt

	for _, val := range arr {
		if m < val {
			m = val
		}
	}

	return m

}
