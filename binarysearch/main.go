package main

import(
 "fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6,7}
	var target int = 4
	for true{
		var high int = nums[len(nums)-1]
		var mid int = len(nums)/2
		if target != high{
			if target < nums[mid]{
				nums = nums[0:mid]

			}else{
				nums = nums[mid:]
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