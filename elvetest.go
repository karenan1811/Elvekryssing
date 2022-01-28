package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	crossingState := true

	firstAttempt := rand.Intn(4)
	if crossingState {
		switch firstAttempt {
		case 0:
			crossingState = false
			fmt.Println("[sheep corn  ---\\ \\human wolf/ _______|---]")
			fmt.Println("Sheep eats the corn!")
		case 1:
			crossingState = false
			fmt.Println("[sheep wolf  ---\\ \\human corn/ _______/---]")
			fmt.Println("Wolf eats the sheep!")
		case 2:
			crossingState = false
			fmt.Println("[sheep wolf corn ---\\ \\human/ _______/---]")
			fmt.Println("Sheep eats the corn!, Wolf eats the sheep!")
		default:
			fmt.Println("[wolf corn ---\\ \\human sheep / _______/---]")
			fmt.Println("Success, Human crossed the sheep, wolf and corn are on the right side.")
		}
	}

	if crossingState {
		// amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("[wolf corn ---\\  _______\\ / /human sheep]")
		fmt.Println("Human think logic and left the sheep on right side alone and crossed the left side")
		fmt.Println("[wolf corn ---\\ \\human/  _______ / sheep]")
	}

	secondAttempt := rand.Intn(2)
	thirdAttempt := rand.Intn(2)
	
	if secondAttempt == 0 {
		fmt.Println("[corn ---\\ \\wolf human/  _______ / sheep]")
		fmt.Println("Human takes the wolf and crosses it to right and leaves corn on the left")
		crossingState = true
		if thirdAttempt == 0 {
			fmt.Println("[corn ---\\  _______\\sheep human/  /wolf ]")
			crossingState = true
			fmt.Println("Human leaves the wolf on the right side and get the sheep on the boat")
			
		} else {
			fmt.Println("[corn ---\\  _______\\human/  / sheep wolf ]")
			crossingState = false
			fmt.Println("Human leaves the wolf and sheep on the right side, takes the boat and the wolf eats the sheep")
		}
	} else if secondAttempt == 1 {
		fmt.Println("[wolf ---\\ \\corn human/  _______ / sheep]")
		fmt.Println("Human takes the corn and crosses it to right and leaves the wolf on the left")
		crossingState = true
		if thirdAttempt == 0 {
			fmt.Println("[wolf ---\\  _______\\sheep human/  /corn ]")
			crossingState = true
			fmt.Println("Human leaves the corn on the right side and get the sheep on the boat")
			
		} else {
			fmt.Println("[wolf ---\\  _______\\human/  / sheep corn ]")
			crossingState = false
			fmt.Println("Human leaves the corn and sheep on the right side, takes the boat and the sheep eats the corn")
		}
	}

	fmt.Println("crossingState is currently:", crossingState)
}
