package main

// closures:
// counter function which return function which return int value
func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	println(increment())
	println(increment())
	println(increment())
	println(increment())
}