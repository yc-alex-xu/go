package main

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	l, la, lb := len(b)+1, len(a), len(b)
	s := make([]byte, l)
	var c byte = 0
	var i int
	for i = 0; i < len(a); i++ {
		s[l-i-1], c = add(a[la-i-1]-'0', b[lb-i-1]-'0', c)
	}
	for ; i < len(b); i++ {
		s[l-i-1], c = add(0, b[lb-i-1]-'0', c)
	}
	if c != 0 {
		s[0] = '1'
		return string(s)
	}
	return string(s[1:])

}

func add(a, b, c byte) (r, rc byte) {
	switch a + b + c {
	case 3:
		r, rc = '1', 1
	case 2:
		r, rc = '0', 1
	case 1:
		r, rc = '1', 0
	case 0:
		r, rc = '0', 0
	}
	return r, rc
}
