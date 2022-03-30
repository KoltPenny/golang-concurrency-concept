package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
	"golang.org/x/sync/semaphore"
)

const maxWaitTime = 10

type Customer struct {
	Secs int64
	id int64
	Served bool
}

type Seat struct {
	Cust *Customer
}

/************** [UNUSED] Refer to README********************/
/*
func initSeats(seats []*Seat) {

	for index,_ := range seats {
		seats[index] = new(Seat)
	}
	
}


func (c *Customer) GetSeat(sem *semaphore.Weighted, seat *Seat) (happy bool) {

	receipt := make(chan bool)
	
	for {

		if sem.TryAcquire(1) { break }
		time.Sleep(500 * time.Millisecond)

	}
	
	_ = seat
	
	seat.Cust = c
	
	go c.wait(receipt)
	happy = <- receipt

	return
}*/
/************** [UNUSED] Refer to README********************/


func sitCustomers(
	seats []*Seat,             
	table_order int64,         
	sems []*semaphore.Weighted,
	wg *sync.WaitGroup) {

/*
- seats: Array of seats in the restaurant.
- table_order: Order type 0: even; type 1: odds.
- sems: Semaphore array (same length as seats).
- wg: Wait Group.
*/
	
	defer wg.Done()

	//Sets an automatic Customer ID.
	var count int64
	count = table_order

	//Resets iterator if it overflows.
	for i:= table_order; true; i+=2 {
		
		if i >= int64(len(seats)) {

			i = table_order
			
		}
		
		//If seat is not taken.
		if seats[i] == nil {

			success := sems[i].TryAcquire(1)

			if !success {
				continue
			}
			
			seats[i] = new(Seat)
			seats[i].initCustomer(count)
			
			count += 2
			
			receipt := make(chan bool)

			//Customer starts eating.
			go seats[i].Cust.wait(receipt)
			go func(r chan bool, idx int64){

				select {
					
				case _ = <- r:
					if seats[idx] != nil{
						seats[idx] = nil
						sems[idx].Release(1)
					}
				}

			}(receipt,i)
			
		}
		
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Seat) initCustomer(index int64) {

	min := maxWaitTime - 4
	seconds := int64(rand.Intn(maxWaitTime - min) + min)
	
	s.Cust = &Customer{ seconds, index, false}
	
}


func (c *Customer) wait(receipt chan bool) {

	for c.Secs > 0 && !c.Served {

		time.Sleep(1000 * time.Millisecond)
		c.Secs -= 1

	}

	receipt <- c.Served

}

func showCustomers(seats []*Seat, speed int) {

	for {

		time.Sleep(time.Duration(speed) * time.Millisecond)

		for idx, seat := range seats {

			if seat != nil {
				fmt.Printf("Table [%d] used by -> %v\n", idx, seat.Cust)
			} else {
				fmt.Printf("Table [%d] is empty.\n", idx)
			}

		}
		print("\n")
	}

}
