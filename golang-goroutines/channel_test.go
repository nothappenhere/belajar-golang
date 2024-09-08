package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"

)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // Selalu di close

	/*
		channel <- "Akbar"  //* Mengirim data kedalam channel
		data := <- channel //* Menerima data dari channel
		fmt.Println(<- channel) //* Atau bisa seperti ini
	*/

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Rizky Akbar"
		fmt.Println("DONE")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Rizky Akbar"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // Selalu di close

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Rizky Akbar"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // Selalu di close

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		channel <- "Akbar"
		channel <- "Rizky"
		channel <- "Muhammad"
	}()

	go func ()  {
		fmt.Println(<- channel)
		// fmt.Println(<- channel)
		// fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("DONE")
}


func TestRangeChannel(t *testing.T){
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	
	for data := range channel{
		fmt.Println("Menerima data", data)
	}

	fmt.Println("DONE")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	defer close(channel1)
	channel2 := make(chan string)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	defer close(channel1)
	channel2 := make(chan string)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}