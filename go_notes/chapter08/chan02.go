package main

func main() {

	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-c
			if !ok {
				return
			}
			println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done
}
