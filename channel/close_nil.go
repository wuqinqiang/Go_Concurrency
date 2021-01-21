package main
func main() {
	var ch chan struct{}
	close(ch)
}
