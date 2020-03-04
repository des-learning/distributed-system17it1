package main

func IntRange(start int, end int, delta int) <-chan int {
	out := make(chan int, 1)
	switch {
	case delta >= 0:
		go func() {
			for i := start; i < end; i += delta {
				out <- i
			}
			close(out)
		}()
	case delta < 0:
		go func() {
			for i := start; i > end; i += delta {
				out <- i
			}
			close(out)
		}()
	default:
		go func() {
			close(out)
		}()
	}
	return out
}

func IntSlice(slice []int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for _, i := range slice {
			out <- i
		}
		close(out)
	}()
	return out
}

func IntFilter(numbers <-chan int, fun func(int) bool) <-chan int {
	out := make(chan int, 1)
	go func() {
		for i := range numbers {
			if fun(i) {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func main() {
	/*for i := range IntRange(1, 100, 1) {
		fmt.Println(i)
	}
	for i := range IntRange(100, 1, -1) {
		fmt.Println(i)
	}*/
	/*for i := range IntSlice([]int{1, 2, 3, 4, 5, 6, 7}) {
		fmt.Println(i)
	}*/
	/*for i := range IntFilter(IntSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		func(x int) bool { return x%2 == 0 }) {
		fmt.Println(i)
	}*/
}
