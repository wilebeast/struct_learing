package first_version

func MaxLenStr(str []byte) {
	indexs := make([][]int, 0, 256)
	byte2indexs := make(map[byte]int)
	for i, v := range str {
		if index, ok := byte2indexs[v]; ok {
			indexs[index] = append(indexs[index], i)
		} else {
			indexs = append(indexs, []int{i})
			byte2indexs[v] = len(indexs) - 1
		}
	}
	maxLen := 0
	start := 0
	end := 0
	for _, v := range indexs {
		if len(v) <= 1 {
			continue
		}
		for i := range v {
			if i == 0 {
				//if v[i] != 0 {
				if v[i+1]-0 > maxLen {
					maxLen = v[i+1] - 0
					start = 0
					end = v[i+1]
				}
				//}
			} else if i == len(v)-1 {
				//if v[i] != len(str)-1 {
				if len(str)-1-v[i-1] > maxLen {
					maxLen = len(str) - 1 - v[i-1]
					start = v[i-1]
					end = len(str)
				}
				//}
			} else {
				if v[i+1]-v[i-1] > maxLen {
					maxLen = v[i+1] - v[i-1]
					start = v[i-1]
					end = v[i+1]
				}
			}
		}
	}
	println("start:%v, end:%v", start, end)
	println("result:%v", string(str[start:end]))
}

func MaxLenStrFinal(str []byte) {
	maxLen := 0
	start := 0
	end := 0
	nextStart := 0
	byte2index := make(map[byte]int)
	for index, char := range str {
		if _, ok := byte2index[char]; ok {
			if index-nextStart > maxLen {
				maxLen = index - nextStart
				start = nextStart
				end = index
				nextStart = byte2index[char] + 1
			}
		}
		byte2index[char] = index
	}
	println("start:%v, end:%v", start, end)
	println("result:%v", string(str[start:end]))
}

func MaxLenStrForce(str []byte) {
	var result []byte
	for index, _ := range str {
		mapExist := make(map[byte]bool)
		temp := make([]byte, 0, 256)
		for i := index; i < len(str); i++ {
			if mapExist[str[i]] {
				if len(result) < len(temp) {
					result = temp
					break
				}
			}
			temp = append(temp, str[i])
			mapExist[str[i]] = true
		}
	}
	println("result:%v", string(result))
}
