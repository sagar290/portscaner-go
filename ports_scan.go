package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

func worker(host string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()

		results <- p
	}
}

func main() {

	n_worker := flag.Int("worker", 5, "number of worker, default is 5")
	host := flag.String("host", "scanme.nmap.org", "host name, default is \"scanme.nmap.org\"")
	min := flag.Int("min", 0, "minimum port number to scan, default is \"0\"")
	max := flag.Int("max", 100, "maximum port number to scan, default is \"100\"")
	flag.Parse()
	fmt.Printf("host: %s port-min: %d; port-max: %d \n", *host, *min, *max)

	ports := make(chan int, *n_worker)

	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {

		go worker(*host, ports, results)
	}

	go func(min int, max int) {
		for i := min; i <= max; i++ {
			fmt.Printf("push port %d to worker for scanning \n", i)
			ports <- i
		}

	}(*min, *max)

	for i := *min; i < *max; i++ {
		port := <-results
		if port != 0 {
			fmt.Printf("found port %d as open\n", port)
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

}
