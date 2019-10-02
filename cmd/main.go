package main

import (
	"fmt"
	"internetCafe/cafe"
	"internetCafe/tourist"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Excersise 3 - Internet cafe")
	fmt.Println("")

	// Fetch args
	nComputer, _ := strconv.ParseInt(os.Args[1], 10, 64)
	nTourists, _ := strconv.ParseInt(os.Args[2], 10, 64)

	// Define structs
	group := tourist.NewGroup()
	cafe := cafe.NewCafe()

	// Goroutine for computer management
	go func() {

		for {
			// Proceed if and only if there's a free computer
			<-cafe.FreeComputer

			// Occupy the free computer
			cafe.OccupyComputer(<-group.Tourists)
		}
	}()

	// Goroutine for tourist management
	go func() {
		for {

			// Iterate over all computers
			for i := 0; i < 8; i++ {

				// If computer is occupied
				if !cafe.Computers[i].IsFree() {

					// Temporary variable as shortcut for the user of the current computer
					tempUser := cafe.GetUser(i)

					// Increase users' time online by one (minute)
					tempUser.TimeOnline++

					// Check if user is done
					if tempUser.TimeOnline > tempUser.LimitOnline {

						// Remove user from computer
						cafe.KickUser(tempUser)

						// Add one to 'users that were already online'
						group.UserCount++

						// Check if all users used the computer
						if group.UserCount == 25 {

							// Let the program continue
							group.IsDone <- true
						}

						// Free the computer
						cafe.FreeComputer <- true
					}
				}
			}

			// "Ticker"
			time.Sleep(time.Millisecond)
		}

	}()

	<-group.IsDone

	fmt.Println("Everyone used the computer. Not it's time to party!")
}
