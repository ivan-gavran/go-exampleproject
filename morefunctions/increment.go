package morefunctions

func Increment(x int) int {
	if x%2 == 0 {
		return AddOne(x)
	} else {
		return HopeForTheBest(x)
	}
}

func AddOne(x int) int {
	return x + 1
}

func HopeForTheBest(x int) int {
	if x%10 == 0 {
		panic("x is too round")
	}
	return x + 1
}
