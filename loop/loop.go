package loop

import "fmt"

func ForLoop() {
	// for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("\nfor_loop_current_index=%d", i)
	}
}

func DoWhileLoop() {
	// do while (post condition)
	loopRound := 3
	for {
		fmt.Printf("\ndo_while_current_number=%d", loopRound)
		loopRound -= 1
		if loopRound <= 0 {
			break
		}
	}
}

func WhileLoop() {
	// while do (pre condition)
	loopRound := 3
	for loopRound >= 0 {
		fmt.Printf("\nwhile_do_current_number=%d", loopRound)
		loopRound -= 1
	}
}
