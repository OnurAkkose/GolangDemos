package main

import (
	"context"
	"fmt"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
)

func main() {
	connStr := "Endpoint=sb://demoeventmaestro.servicebus.windows.net/;SharedAccessKeyName=ilke;SharedAccessKey=JxZskjveQJ4Omv/0E4z3q50bVaQ/XST1JnnISbO/wu8=;EntityPath=firstentity"
	hub, err := eventhub.NewHubFromConnectionString(connStr)

	if err != nil {
		fmt.Println("dgsdg", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	runtimeInfo, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(runtimeInfo)
	}
	// send a single message into a random partition
	// err = hub.Send(ctx, eventhub.NewEventFromString(""))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Println("sent")
	// }
	//eventhub.NewEvent([]byte("df"))

	// handler := func(c context.Context, event *eventhub.Event) error {
	// 	fmt.Println(string(event.Data))
	// 	return nil
	// }

	// // listen to each partition of the Event Hub
	// runtimeInfo, err := hub.GetRuntimeInformation(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for _, partitionID := range runtimeInfo.PartitionIDs {
	// 	// Start receiving messages
	// 	//
	// 	// Receive blocks while attempting to connect to hub, then runs until listenerHandle.Close() is called
	// 	// <- listenerHandle.Done() signals listener has stopped
	// 	// listenerHandle.Err() provides the last error the receiver encountered
	// 	listenerHandle, err := hub.Receive(ctx, partitionID, handler)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(listenerHandle)
	// }

	// // Wait for a signal to quit:
	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt, os.Kill)
	// <-signalChan

	err = hub.Close(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
