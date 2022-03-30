# Restaurant concept program for Parallelism and Concurrency

## Requirements

Before you start your program, you need have the following versions installed:

- go version 1.16.15

Run `go mod tidy` to install all the necessary dependencies.

## Getting Started

To start the program, run `go run *.go`.

You can also run the program with the following custom parameters:

- `go run *.go <printing_frequency (milliseconds)> <semaphore_quantity>`

## To Do

- Create `Customer` queues and the logic for handling them (see commented functions in `Customer.go`).
- Initialize persistent seats
- Generate random names for `Waiter` and `Customer` to improve readability.
- Add more parameters for customization (cores, printing info, etc).

## Additional Info

- This project was tested on a `Fedora Linux 35 - x86_64` system, last update on the 25th of March.
- Printing needs improvement.
