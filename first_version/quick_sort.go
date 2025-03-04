package first_version

func quickSort(arr *[]int64, l, r uint) {
	if l >= r {
		return
	}
	x := l + 1
	y := r
	z := l
	for {
		if x >= y {
			break
		}
		if (*arr)[x] > (*arr)[z] && (*arr)[y] < (*arr)[z] {
			temp := (*arr)[z]
			(*arr)[z] = (*arr)[y]
			(*arr)[y] = (*arr)[x]
			(*arr)[x] = temp
			z = x
		}
		if (*arr)[y] >= (*arr)[z] {
			y--
		}
		if (*arr)[x] <= (*arr)[z] {
			x++
		}
	}
	quickSort(arr, l, z-1)
	quickSort(arr, z+1, r)
}
