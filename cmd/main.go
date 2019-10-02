package main

import (
	"fmt"
	"internetCafe/cafe"
	"internetCafe/tourist"
	"time"
)

func main() {
	fmt.Println("Excersise 3 - Internet cafe")

	// Define structs
	group := tourist.NewGroup()
	cafe := cafe.NewCafe()

	// fmt.Println(group)
	// fmt.Println(cafe)

	// Goroutine for computer management
	go func() {

		// Go as long as the group isn't done
		for !group.IsDone {

			// Proceed if and only if there's a free computer
			<-cafe.FreeComputer

			fmt.Println("Free computer aviable")

			cafe.OccupyComputer(<-group.Tourists)

			fmt.Print("The free computer is now occupied\n\n")
		}

	}()

	// Goroutine for tourist management
	go func() {
		for !group.IsDone {

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

						// Free the computer
						cafe.FreeComputer <- true

					}
				}
			}

			// "Ticker"
			time.Sleep(time.Millisecond)
		}

	}()

	time.Sleep(time.Second)
}
