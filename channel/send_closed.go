package main

func main() {
	ch := make(chan struct{})
	close(ch)
	ch <- struct{}{}
}
