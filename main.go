package main

import (
	"os"
	"fmt"
	"sync"
	"strconv"
	"golang.org/x/sync/semaphore"
)

var print_speed = 750
var available_sems = 10
var available_waiters = 3
var waiter_names = []string{"Rudolf", "Raplh", "Rihanna"}

func main () {

	//Command line argument validation
	if len(os.Args) > 1 {

		if arg := os.Args[1]; arg != "" {
			print_speed, _ = strconv.Atoi(arg)
		}

		if arg := os.Args[2]; arg != "" {
			available_sems, _ = strconv.Atoi(arg)
		}
		
	}

	//Create wait group
	wg := new(sync.WaitGroup)
	wg.Add(2 + available_waiters)

	//Create the Seats array and Waiters
	customer_seats := make([]*Seat, available_sems)
	waiters := make([]*Waiter, available_waiters)

	//Create semaphgre arrays for Waiters and Customers
	seat_sems := make([]*semaphore.Weighted, available_sems)
	serv_sems := make([]*semaphore.Weighted, available_sems)

	//Initialize semaphore arrays
	for i := 0; i < available_sems; i++ {

		seat_sems[i] = semaphore.NewWeighted(int64(1))
		serv_sems[i] = semaphore.NewWeighted(int64(1))
		
	}

	//Start goroutines for even and oddtable numbers
	go sitCustomers(customer_seats, even, seat_sems,wg)
	go sitCustomers(customer_seats, odd, seat_sems,wg)

	//Start waiter goroutines to serve the different tables
	for i,waiter := range waiters {
		
		waiter = new(Waiter)
		waiter.Name = waiter_names[i]
		
		fmt.Println(&waiter)
		go waiter.ServeCustomers(serv_sems, customer_seats, wg)
	}

	//Print the customers using tables
	go showCustomers(customer_seats,print_speed)

	wg.Wait()

}
