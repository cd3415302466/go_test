package main

func main() {
	queue := make(chan string, 10)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	queue <- "five"
	queue <- "six"
	close(queue)

	for elem := range queue {
		println(elem)
	}
}
