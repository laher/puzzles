// "man has 3 bags of 30 coins each. Man goes through 30 toll gates where he has to pay toll = the number of bags of coins he has. How many does he have left at the end?"

package main

import "fmt"
import "flag"

type BagType struct {
	bags     []int
	num_bags int
}

// calculate toll given current coin balance
func toll(b BagType) int {
	non_zero := 0

	for i := 0; i < b.num_bags; i++ {
		if b.bags[i] != 0 {
			non_zero = non_zero + 1
		}
	}

	return (non_zero)
}

func pay_toll_lazy(b BagType, toll int) {
	this_bag := 0
	to_pay := toll

	for to_pay > 0 && this_bag < 3 {
		if b.bags[this_bag] == 0 {
			this_bag = this_bag + 1
			continue
		}

		if b.bags[this_bag] >= to_pay {
			b.bags[this_bag] = b.bags[this_bag] - to_pay
			to_pay = 0
		} else {
			to_pay -= b.bags[this_bag]
			b.bags[this_bag] = 0
			this_bag++
		}
	}
}

// lazy bollocks, recurse the mofo
func pay_toll_even(b BagType, toll int) {
//	fmt.Println(b, " ", toll)
	fullest_bag := 0
	amount_in_fullest_bag := 0

	for i := 0; i < b.num_bags; i++ {
		if amount_in_fullest_bag < b.bags[i] {
			amount_in_fullest_bag = b.bags[i]
			fullest_bag = i
		}
	}

	// sanity check: there's money to take, right? (this is the base case for recursion)
	if b.bags[fullest_bag] == 0 {
		fmt.Println("WTF?! need to pay", toll, "but the fullest bag is empty!")
		return
	}

//	fmt.Println("paying a coin from bag", fullest_bag)
	b.bags[fullest_bag]--
	if toll-1 > 0 {
		pay_toll_even(b, toll-1)
	}
}

func total_bucks(b BagType) int {
	bucks := 0
	for i := 0; i < b.num_bags; i++ {
		bucks += b.bags[i]
	}
	return bucks
}

func main() {
	var even = flag.Bool("even", true, "Pay coins evenly")
	flag.Parse()
	
	b := BagType{make([]int, 3), 3}
	which_booth := 1
	for i := 0; i < b.num_bags; i++ {
		b.bags[i] = 30
	}

	for which_booth <= 30 {
		t := toll(b)
		fmt.Print("Toll gate ", which_booth, ", Toll = ", t, ", Start = $", total_bucks(b))
		if *even {
			pay_toll_even(b, t)
		} else {
			pay_toll_lazy(b, t)
		}
		fmt.Print(", End = $", total_bucks(b))
		fmt.Println()
		which_booth = which_booth + 1
	}
}
