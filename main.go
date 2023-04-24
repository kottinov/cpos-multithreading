package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxConcurrentThreads = 10
	maxBufferSize        = 100
)

var wg sync.WaitGroup
var sensorData = make(chan string, maxBufferSize)
var tvData = make(chan string, maxBufferSize)

func main() {
	// initializing sensor and TV streams
	sensorStreams := []string{"Sensor1", "Sensor2", "Sensor3"}
	tvStreams := []string{"TVStream1", "TVStream2", "TVStream3"}

	// creating a goroutine pool to limit the number of concurrent threads
	threadPool := make(chan struct{}, maxConcurrentThreads)

	// we start a separate thread for each sensor stream
	for _, stream := range sensorStreams {
		threadPool <- struct{}{}
		wg.Add(1)
		go func(stream string) {
			defer wg.Done()
			readSensorData(stream)
			<-threadPool
		}(stream)
	}

	// we start a separate thread for each TV stream
	for _, stream := range tvStreams {
		threadPool <- struct{}{}
		wg.Add(1)
		go func(stream string) {
			defer wg.Done()
			decodeTVStream(stream)
			<-threadPool
		}(stream)
	}

	// here we patiently wait for all threads to complete before exiting the program
	wg.Wait()

	fmt.Println("Sensor data:")
	close(sensorData)
	for data := range sensorData {
		fmt.Println(data)
	}

	fmt.Println("TV stream data:")
	close(tvData)
	for data := range tvData {
		fmt.Println(data)
	}
}

func readSensorData(sensor string) {
	// here we simulate reading data from a sensor 
	for i := 1; i <= 10; i++ {
		data := fmt.Sprintf("Sensor %s data point %d", sensor, i)
		sensorData <- data
		time.Sleep(time.Millisecond * 500) // here im simulating time-consuming task
	}
}

func decodeTVStream(stream string) {
	// same thing from a TV stream
	for i := 1; i <= 10; i++ {
		data := fmt.Sprintf("Decoded data from TV stream %s, frame %d", stream, i)
		tvData <- data
		time.Sleep(time.Millisecond * 500) // same here
	}
}
