package main

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range ints {
		if num == 0 {
			println(num, ":Neither Even nor Odd")
		}
		if num%2 == 0 {
			println(num, ":Even")
		} else {
			println(num, ":Odd")
		}
	}
}
