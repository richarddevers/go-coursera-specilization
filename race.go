// to check if there's a race condition, run it using:
// go run -race race.go

// My ouput:

// WARNING: DATA RACE
// Read at 0x00c00011a008 by goroutine 7:
//   main.increase_counter()
//       /workspaces/go-coursera/race.go:5 +0x34
// Previous write at 0x00c00011a008 by goroutine 6:
//   main.increase_counter()
//       /workspaces/go-coursera/race.go:5 +0x48
// Goroutine 7 (running) created at:
//   main.main()
//       /workspaces/go-coursera/race.go:12 +0x88
// Goroutine 6 (running) created at:
//   main.main()
//       /workspaces/go-coursera/race.go:11 +0x68
// ==================
// Found 1 data race(s)
// exit status 66

//explanation:
// the same variable "counter" is accessed by two goroutine
// which try in an endless loop to increse it's value.
// During one of the update
// at the machine level the value is read from the memory
// but then another goroutine increse it's value before the current operation
// which invalidate the operation since the value has change between the current read and update operation

package main

func increase_counter(counter *int) {
	for {
		*counter += 1
	}
}

func main() {
	var counter int
	go increase_counter(&counter)
	go increase_counter(&counter)
}
