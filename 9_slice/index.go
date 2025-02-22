package main

import (
	"fmt"
	"slices"
)

//* slice -> dynamic arrays
//* most used construct in go
//* useful methods
func main(){
	//* uninitialized slice is nil
	var num []int
	fmt.Println(num)
	fmt.Println(num == nil)
	fmt.Println(len(num))

	//* capacity -> maximum numbers of elements can fit
	var nums = make([]int,2,5)
	fmt.Println(cap(nums) , len(nums))
	fmt.Println(nums)

	nums[0] = 3
	nums[1] = 4

	fmt.Println(nums)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 4)
	fmt.Println("My capacity is => ",cap(nums))
	fmt.Println(nums)
	fmt.Println(len(nums))

	arr := []int{}
	arr = append(arr, 1)
	fmt.Println(arr)

	//* copy func
	var nums1 = make([]int , 0 , 5)
	nums1 = append(nums1, 2)
	var nums2 = make([]int , len(nums1))

	copy(nums2 , nums1)

	fmt.Println(nums1 , nums2)

	//* slice operator
	a := []int{1,2,3}
	fmt.Println(a[:])

	//* slice func
	a1 := []int{1,2}
	a2 := []int{1,2,3}
	fmt.Println(slices.Equal(a1 , a2))

	D_2 := [][]int{{1,2},{2,3}}
	fmt.Println(D_2)

	v := make([]int , 10)
	fmt.Println(v)

	st := "seven"
	byteArray := []byte(st)
	byteArray[0] = 'n'
	fmt.Println(byteArray)

}

