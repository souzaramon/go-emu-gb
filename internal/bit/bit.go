package bit

func GetNth(b byte, n int) bool {
	return b&(byte(1)<<n) != 0
}

func SetNth(b byte, n int, v bool) byte {
	if v {
		return b | (byte(1) << n)
	} else {
		return b & ^(byte(1) << n)
	}
}
