package main

func main() {
	ch := make(chan struct{})
	close(ch)
	close(ch)
}
