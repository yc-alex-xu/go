package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l, l1, l2 := len(nums1)+len(nums2), len(nums1), len(nums2)
	i1, i2 := 0, 0

	var v int
	for cnt := 0; cnt < (l+1)/2; cnt++ {
		v = next(&i1, nums1, l1, &i2, nums2, l2)
	}
	if l%2 != 0 {
		return float64(v)
	}
	v2 := next(&i1, nums1, l1, &i2, nums2, l2)
	return float64(v+v2) / 2

}

//谁是归并排序中的一下
func next(i1 *int, nums1 []int, l1 int, i2 *int, nums2 []int, l2 int) (v int) {
	if *i1 < l1 && *i2 < l2 {
		if nums1[*i1] < nums2[*i2] {
			v = nums1[*i1]
			(*i1)++
		} else {
			v = nums2[*i2]
			(*i2)++
		}
		return v
	}
	if *i1 == l1 {
		v = nums2[*i2]
		(*i2)++
		return v
	}
	v = nums1[*i1]
	(*i1)++
	return v
}
