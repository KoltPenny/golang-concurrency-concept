package main

import(
	"fmt"
	"time"
	"sync"
	"golang.org/x/sync/semaphore"
)

const (
	even int64 = 0
	odd int64 = 1
)

type Waiter struct {
	Name string
	Services int64
	Served []*Customer
}

func (w *Waiter) serve(c *Customer) {
	/*
- c: Customer pointer
*/
	//When serving add Customer to Served list, and add 1 to the counter.
	
	c.Served = true
	w.Services += 1
	w.Served = append(w.Served, c)
	fmt.Println("Waiter",w.Name,"served",w.Services,"customers.")
}

func (w *Waiter) AcquireService(sem *semaphore.Weighted, seat *Seat) {
	/*
- sem: Semaphore pointer.
- seat: Seat pointer.
*/
	//Consume a semaphore to serve table.
	if sem.TryAcquire(1) {

		//Test if there are both a seat and a Customer available.
		if seat != nil && seat.Cust != nil {

			if c := seat.Cust; !c.Served {
				w.serve(c)
			}
			
		}

		sem.Release(1)
	}

}

func (w *Waiter) ServeCustomers (service_sems []*semaphore.Weighted, seats []*Seat, wg *sync.WaitGroup) {
	/*
- service_sems: Semaphore array.
- seats: Array of seats in the restaurant.
- wg: Wait Group.
*/
	//Iterates over all tables to try and serve them.
	defer wg.Done()

	for index := 0; index < len(seats); index++ {
		
		if index == len(seats)-1 {
			
			index = 0
			
		}

		w.AcquireService(service_sems[index], seats[index])
		
		time.Sleep(800 * time.Millisecond)
		
	}

	
}
