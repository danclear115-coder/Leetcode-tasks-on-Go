package main

import "fmt"

func getSortSlice(slice []int) (sortSlice []int) {

	for len(slice) > 0 {

		minElement, minIndex := slice[0], 0

		for i := 0; i < len(slice); i++ {
			if slice[i] < minElement {
				minElement, minIndex = slice[i], i
			}
		}

		sortSlice = append(sortSlice, minElement)
		slice = append(slice[:minIndex], slice[minIndex+1:]...)

	}

	return

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	numsAll := getSortSlice(append(nums1, nums2...))
	fmt.Println(numsAll)

	if len(numsAll)%2 == 0 {

		element1, element2 := numsAll[len(numsAll)/2-1], numsAll[len(numsAll)/2]
		return float64(element1 + element2) / 2.0

	} else {
		return float64(numsAll[(len(numsAll) - 1) / 2])
	}

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3, 5}, []int{2, 4}))
}
