package main

func main() {
	done := make(chan struct{})
	c := make(chan string)
	println("1")
	go func() {
		println("5")
		s := <-c

		println(s)

		defer close(done)
		println("4")
	}()
	println("2")
	c <- "hi!"
	println("3")
	<-done
}
