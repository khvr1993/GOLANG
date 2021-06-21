package channels

import "log"

/*
	Counter will generate numbers so that they can be sent to the routine of squarer
	c acts as send only channel
*/
func counter(c chan<- int) {
	log.Println("Begin Counter ")
	for x := 0; x < 5; x++ {
		log.Println("x in Counter => ", x)
		c <- x
	}
	log.Println("End Counter ")
	close(c)
}

/*
	Squarer channel will wait for the counter channel to send a value and after receiving it will get squared
	in will control the sync of squares
	Here in acts as the receive only channel
	out acts as send only channel
*/
func squarer(out chan<- int, in <-chan int) {
	log.Println("Begin squarer")
	for v := range in {
		out <- v * v
	}
	log.Println("End squarer")
	close(out)
}

/*
	As and when squarer channel sends the data printer will print it
	In acts as receive only channel
*/

func printer(in <-chan int) {
	log.Println("Begin printer")
	for v := range in {
		log.Println(v)
	}
	log.Println("End printer")
}

func TestChan() {
	naturals := make(chan int, 0)
	squares := make(chan int, 0)

	/*
		The order of execution is not guaranteed but
		squarer function starts only when the channel input is received to it by the
		counter routine
	*/
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}
