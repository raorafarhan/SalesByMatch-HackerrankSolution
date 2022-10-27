func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	socks := make(map[string]int32)
	for _, sock := range ar {
		socks[strconv.FormatInt(int64(sock), 10)] += 1
	}

	var result int32
	for _, v := range socks {
		result += v / 2
	}

	return result

}