package main

import (
	"fmt"
)

// func main() {

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	ch := make(chan struct{}, 1)

// 	go func() {

// 		for i := 0; i < 10; i++ {
// 			ch <- struct{}{}
// 			if i%2 == 0 {
// 				fmt.Println("even:", i)
// 			}

// 			<-ch
// 		}

// 		wg.Done()

// 	}()

// 	go func() {

// 		for i := 0; i < 10; i++ {
// 			<-ch
// 			if i%2 != 0 {
// 				fmt.Println("odd:", i)
// 			}

// 			ch <- struct{}{}
// 		}

// 		wg.Done()

// 	}()

// 	wg.Wait()

// }

// type Stack struct {
// 	Data []interface{}
// }

// func (s *Stack) enqueue(val interface{}) {
// 	s.Data = append(s.Data, val)
// }

// func (s *Stack) IsEmpty() bool {
// 	return len(s.Data) == 0
// }

// func (s *Stack) dequeue() {
// 	if s.IsEmpty() {

// 		fmt.Println("Stack is Empty")

// 		return

// 	}

// 	last := len(s.Data) - 1

// 	fmt.Println(s.Data[last])

// 	s.Data = s.Data[:last]

// }

func main() {

	// var stk Stack

	// stk.enqueue(1)
	// stk.enqueue(2)
	// stk.enqueue(3)

	// stk.dequeue()
	// stk.dequeue()

	arr := make([]int, 5, 6)

	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	//0,1,2,3,4

	arr1 := arr[:3]

	//0,1,2,6

	arr1 = append(arr1, 6)

	fmt.Println(arr, arr1)

	arr = append(arr, 7)

	fmt.Println(arr, arr1)

	// 0,1,2,6,4,7 0,1,2,6

	//0,1,2,3,4,6 arr

	//0,1,2,

}
