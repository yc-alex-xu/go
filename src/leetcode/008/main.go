package main

func myAtoi(str string) int {
	s := []byte(str)
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return 0
	}

	valAbs, signed := 0, 1

	if v := s[i]; v == '+' {
		//do nothing
	} else if v == '-' {
		signed = -1
	} else if v >= '0' && v <= '9' {
		valAbs = int(v - '0')
	} else {
		return 0
	}

	for _, v := range s[i+1:] {
		if v >= '0' && v <= '9' {
			valAbs = valAbs*10 + int(v-'0')
			if valAbs >= 1<<31 {
				break
			}
		} else {
			break
		}
	}

	if signed == 1 && valAbs > 1<<31-1 {
		return 1<<31 - 1
	}
	if signed == -1 && valAbs > 1<<31 {
		return -(1 << 31)
	}
	return signed * valAbs
}
