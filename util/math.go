package util

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SignInt(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}
