package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	const requiredArgsCount = 2
	var timeout time.Duration
	var wg sync.WaitGroup

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [--timeout duration] host port\n", os.Args[0])
	}

	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout for server connection")
	flag.Parse()

	if flag.NArg() != requiredArgsCount {
		flag.Usage()
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect: %s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "...Connected to %s\n", net.JoinHostPort(host, port))

	sendCh := make(chan struct{}, 1)
	defer close(sendCh)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := client.Send(); err != nil {
			fmt.Fprintf(os.Stderr, "send error: %s\n", err)
		}
		sendCh <- struct{}{}
	}()

	receiveCh := make(chan struct{}, 1)
	defer close(receiveCh)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := client.Receive(); err != nil {
			fmt.Fprintf(os.Stderr, "receive error: %s\n", err)
		}
		receiveCh <- struct{}{}
	}()

	select {
	case <-sendCh:
		fmt.Fprintf(os.Stderr, "...EOF\n")
	case <-receiveCh:
		fmt.Fprintf(os.Stderr, "...Connection was closed by peer\n")
	}
	client.Close()

	wg.Wait()
}
