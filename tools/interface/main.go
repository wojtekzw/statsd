package main

import (
	"fmt"
	"github.com/wojtekzw/statsd"
	"reflect"
	// "os"
)

func main() {
	var (
		statsdClient statsd.Statser
		noopClient   *statsd.NoopClient
	)

	statsdClient, err := statsd.New(statsd.Address("localhost:8125"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		// statsdClient = &statsd.NoopClient{}
		statsdClient = noopClient
	}

	fmt.Printf("(%v %T)\n", statsdClient, statsdClient)

	fmt.Printf("Init OK - sending metric\n")
	t := statsdClient.NewTiming()
	for i := 0; i < 10000; i++ {
		statsdClient.Increment("test.metric" + string(i))
		t.Send("test.metric.timing")
	}

	fmt.Printf("StatsD: %s, Timing: %s\n", reflect.TypeOf(statsdClient), reflect.TypeOf(t))
}
