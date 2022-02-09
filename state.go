// package state

// func ViewState() string {
//     // Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
//     return "[kylling rev korn hs ---\\ \\__/ _________________/---]"
// }

package state

import (
	"fmt"
	// "math/rand"
	// "time"
)

func ViewState(firstPrøve int, secondPrøve int, thirdPrøve int) string {
	// rand.Seed(time.Now().UnixNano())
	crossingState := true
	initialState := "[sheep corn human wolf  ---\\ \\/ _________________|---]"
	//firstPrøve := rand.Intn(4)
	// if firstPrøve == 0 {
	// 	crossingState = false
	// 	fmt.Println("[sheep corn  ---\\ \\human wolf/ _________________|---]")
	// 	fmt.Println("Sheep eats the corn!")
	// 	initialState = "[sheep corn  ---\\ \\human wolf/ _________________|---]"
	// } else if firstPrøve == 1 {
	// 	crossingState = false
	// 	fmt.Println("[sheep wolf  ---\\ \\human corn/ _________________/---]")
	// 	fmt.Println("Wolf eats the sheep!")
	// 	initialState = "[sheep wolf  ---\\ \\human corn/ _________________/---]"
	// } else if firstPrøve == 2 {
	// 	crossingState = false
	// 	fmt.Println("[sheep wolf corn ---\\ \\human/ _________________/---]")
	// 	fmt.Println("Sheep eats the corn!, Wolf eats the sheep!")
	// 	initialState = "[sheep wolf corn ---\\ \\human/ _________________/---]"
	// } else {
	// 	fmt.Println("[wolf corn ---\\ \\human sheep / _________________/---]")
	// 	fmt.Println("Success, Human crossed the sheep, wolf and corn are on the right side.")
	// 	initialState = "[wolf corn ---\\ \\human sheep / _________________/---]"
	// }
	if crossingState {
		switch firstPrøve {
		case 0:
			crossingState = false
			fmt.Println("[sheep corn  ---\\ \\human wolf/ _________________|---]")
			fmt.Println("Sheep eats the corn!")
			initialState = "[sheep corn  ---\\ \\human wolf/ _________________|---]"

		case 1:
			crossingState = false
			fmt.Println("[sheep wolf  ---\\ \\human corn/ _________________/---]")
			fmt.Println("Wolf eats the sheep!")
			initialState = "[sheep wolf  ---\\ \\human corn/ _________________/---]"

		case 2:
			crossingState = false
			fmt.Println("[sheep wolf corn ---\\ \\human/ _________________/---]")
			fmt.Println("Sheep eats the corn!, Wolf eats the sheep!")
			initialState = "[sheep wolf corn ---\\ \\human/ _________________/---]"

		default:
			fmt.Println("[wolf corn ---\\ \\human sheep / _________________/---]")
			fmt.Println("Success, Human takes the sheep, leaves wolf and corn on the right side.")
			initialState = "[wolf corn ---\\ \\human sheep / _________________/---]"

		}
	}

	if crossingState {

		fmt.Println("[wolf corn ---\\  _________________\\ / /human sheep]")
		fmt.Println("Human think logic and left the sheep on right side alone and crossed the left side")
		fmt.Println("[wolf corn ---\\ \\human/  _________________ / sheep]")
		initialState = "[wolf corn ---\\ \\human/  _________________ / sheep]"
	}

	// secondPrøve := rand.Intn(2)
	// thirdPrøve := rand.Intn(2)
	if secondPrøve == 0 && crossingState {
		fmt.Println("[corn ---\\ \\wolf human/  _________________ / sheep]")
		fmt.Println("Human takes the wolf and crosses it to right and leaves corn on the left")
		crossingState = true
		if thirdPrøve == 0 {
			fmt.Println("[corn ---\\  _________________\\sheep human/  /wolf ]")
			crossingState = true
			fmt.Println("Human leaves the wolf on the right side and get the sheep on the boat")
			initialState = "[corn ---\\  _________________\\sheep human/  /wolf ]"
		} else {
			fmt.Println("[corn ---\\  _________________\\human/  / sheep wolf ]")
			crossingState = false
			fmt.Println("Human leaves the wolf and sheep on the right side, takes the boat and the wolf eats the sheep")
			initialState = "[corn ---\\  _________________\\human/  / sheep wolf ]"
		}
	} else if secondPrøve == 1 && crossingState {
		fmt.Println("[wolf ---\\ \\corn human/  _________________ / sheep]")
		fmt.Println("Human takes the corn and crosses it to right and leaves the wolf on the left")
		crossingState = true
		initialState = "[wolf ---\\ \\corn human/  _________________ / sheep]"
		if thirdPrøve == 0 {
			fmt.Println("[wolf ---\\  _________________\\sheep human/  /corn ]")
			crossingState = true
			fmt.Println("Human leaves the corn on the right side and get the sheep on the boat")
			initialState = "[wolf ---\\  _________________\\sheep human/  /corn ]"
		} else {
			fmt.Println("[wolf ---\\  _________________\\human/  / sheep corn ]")
			crossingState = false
			fmt.Println("Human leaves the corn and sheep on the right side, takes the boat and the sheep eats the corn")
			initialState = "[wolf ---\\  _________________\\human/  / sheep corn ]"
		}
	}
	if secondPrøve == 0 && thirdPrøve == 0 && crossingState {
		fmt.Println("[corn sheep human---\\ \\/  _________________  /wolf ]")
		fmt.Println("Human think logic again and leave the sheep on left side alone and take the corn")
		fmt.Println("[sheep  ---\\ \\human corn/  _________________ /wolf]")
		initialState = "[sheep  ---\\ \\human corn/  _________________ /wolf]"
	} else if secondPrøve == 1 && thirdPrøve == 0 && crossingState {
		fmt.Println("[wolf sheep human---\\ \\/  _________________  /corn ]")
		fmt.Println("Human think logic again and leave the sheep on left side alone and take the corn")
		fmt.Println("[sheep  ---\\ \\human wolf/  _________________ /corn]")
		initialState = "[sheep  ---\\ \\human wolf/  _________________ /corn]"
	}
	fmt.Println("crossingState is currently:", crossingState)
	if crossingState {
		fmt.Println("[sheep  ---\\  _________________ \\ / / human wolf corn]")
		fmt.Println("______________________________________________________ ")
		fmt.Println("Human think logic 3rd time and does not want to repeat whole proccess and return alone to left side and leave wolf and corn on the right side")
		fmt.Println("[  ---\\ \\human sheep /   _________________ / wolf corn]")
		fmt.Println("Human is able to cross wolf, sheep and corn without problem")
		initialState = "[  ---\\    _________________\\  / / human sheep wolf corn]"
		fmt.Println("crossingState is currently:", crossingState)
		//fmt.Print(initialState)
	}
	return initialState
}
