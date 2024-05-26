package main

import(
 "fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6,7}
	var target int = 4
	for true{
		if target != nums[len(nums)-1]{
		if target < nums[len(nums)/2]{
			nums = nums[0:len(nums)/2]
			_ = nums
		}else{
			nums = nums[len(nums)/2:]
			_ = nums
	}
	}else{
		fmt.Println(nums[len(nums)-1],"gg")
		break
		}
	if len(nums)==1 && target != nums[len(nums)-1] {
		fmt.Println("target not in nums")
		break
	}
	}
}