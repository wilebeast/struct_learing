package max_length_str_20260316

func maxLengthStr(input string) string {
	maxStart, maxEnd := 0, 0
	start := 0
	existByte := make(map[byte]int)
	for index := 0; index < len(input); index++ {
		if value, ok := existByte[input[index]]; ok && value >= start {
			if index-start > maxEnd-maxStart {
				maxStart = start
				maxEnd = index
			}
			start = existByte[input[index]] + 1
		}
		existByte[input[index]] = index
	}
	return input[maxStart:maxEnd]
}
