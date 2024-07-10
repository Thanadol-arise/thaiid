package thaiid

func IsValid(id string) bool {
	validId := ""
	for i := 0; i < len(id); i++ {
		if int(id[i]-'0') >= 0 && int(id[i]-'0') <= 9 {
			validId = validId + string(id[i])
		}
	}
	println(validId)

	if len(validId) != 13 {
		return false
	}

	sum := 0
	for i := 0; i < len(validId)-1; i++ {
		sum += int(validId[i]-'0') * (len(validId) - i)
	}
	println(sum)

	remain := sum % 11
	println(remain)

	diff := (11 - remain) % 10

	if diff == int(validId[12]-'0') {
		return true
	}

	return false
}
