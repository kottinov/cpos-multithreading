## Concurrent Data Collection with Goroutines and Channels
This is a simple Go program that demonstrates how to collect data from multiple sensors and TV streams concurrently using goroutines and channels.

## Prerequisites
This program requires the Go language to be installed on your system. Please refer to the official Go website for installation instructions.

## How to Run
To run the program, simply execute the following command in your terminal:

`go run main.go`

## Explanation
- The program creates a pool of goroutines to limit the number of concurrent threads using a buffered channel threadPool. Each sensor and TV stream is started in its own goroutine. The readSensorData and decodeTVStream functions simulate the collection of sensor and TV stream data, respectively. They both send the collected data to the corresponding channel (sensorData or tvData).

- The program uses two channels to collect the sensor and TV stream data: sensorData and tvData. These channels have a maximum buffer size of 100 to avoid overflowing the memory. Once all data has been collected, the program iterates through the channels and prints the collected data.
