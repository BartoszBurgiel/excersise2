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

			<-cafe.FreeComputer

			fmt.Println("Free computer aviable")

			cafe.OccupyComputer(<-group.Tourists)

			fmt.Print("The free computer is now occupied\n\n")
		}

	}()

	// Goroutine for tourist management
	go func() {
		for !group.IsDone {

			// Increase users' time online
			for i := 0; i < 8; i++ {

				if !cafe.Computers[i].IsFree() {

					cafe.Computers[i].User.TimeOnline++

					// check if user is done
					if cafe.Computers[i].User.TimeOnline > cafe.Computers[i].User.LimitOnline {

						// remove user from computer
						cafe.KickUser(cafe.Computers[i].User)

						// free a computer
						cafe.FreeComputer <- true

					}
				}
			}

			time.Sleep(time.Millisecond)
		}

	}()

	time.Sleep(time.Second)
}
