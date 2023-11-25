package main

import (
	"context"
	"fmt"
	"time"
)

func contextCancelMethod() {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second*5)

	defer cancel()

	networkReceiver := make(chan string)
	databaseReceiver := make(chan string)

	go getDataFromNetwork(ctx, networkReceiver)
	go getDataFromDatabase(ctx, databaseReceiver)

	defer close(networkReceiver) // close the channel
	defer close(databaseReceiver)

	for i := 0; i < 3; i++ {
		select { // wait on multiple channel operations, select value from channel that first fulfills
		case networkData := <-networkReceiver:
			fmt.Println("Acquired:", networkData)
		case databaseData := <-databaseReceiver:
			fmt.Println("Acquired:", databaseData)
		case <-ctx.Done():
			fmt.Println("Cli: Process timed-out")
		}
	}

	time.Sleep(time.Second * 4) // wait for 4 second until program terminated
}

func getDataFromNetwork(ctx context.Context, dataReceiver chan string) {
	startTime := time.Now()

	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		fmt.Println("Surfing Network..")
		if time.Since(startTime).Seconds() >= 4 {
			ticker.Stop()
			break
		}
	}

	if ctx.Err() == nil {
		dataReceiver <- "Network Data"
	}
}

func getDataFromDatabase(ctx context.Context, dataReceiver chan string) {
	startTime := time.Now()

	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		fmt.Println("Exploring DB..")
		if time.Since(startTime).Seconds() >= 6 {
			ticker.Stop()
			break
		}
	}

	if ctx.Err() == nil {
		dataReceiver <- "Database Data"
	}
}
