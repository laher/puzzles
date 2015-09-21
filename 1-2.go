// "man has 3 bags of 30 coins each. Man goes through 30 toll gates where he has to pay toll = the number of bags of coins he has. How many does he have left at the end?"

// to test the evenly distributed algorithm (pay each penny of toll from the fullest bag)
//   go run 1.2 -even=true
// to test the lazy algorithm (empty bags as soon as possible)
//   go run 1.2 -even=false
// (default = even=true)

package main

import "fmt"
import "flag"

// I hated hard-coding 3 as the number of bags
type BagType struct {
	bags     []int
	num_bags int
}

// calculate toll given current coin balance
// toll = number of non-empty bags
func toll(b BagType) int {
	non_zero := 0

	for i := 0; i < b.num_bags; i++ {
		if b.bags[i] != 0 {
			non_zero = non_zero + 1
		}
	}

	return (non_zero)
}

// lazy system: drain first bag, then move to second bag, ...
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

// lazy be bollocks, recurse the mofo
// even payment = divide tolls across all bags
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


// calculate how much money I have left across all me money bags
func total_bucks(b BagType) int {
	bucks := 0
	for i := 0; i < b.num_bags; i++ {
		bucks += b.bags[i]
	}
	return bucks
}

// run the simulation
func main() {
	var even = flag.Bool("even", true, "Pay coins evenly")
	flag.Parse()
	
	// three money bags of 30 each
	b := BagType{make([]int, 3), 3}
	which_booth := 1
	for i := 0; i < b.num_bags; i++ {
		b.bags[i] = 30
	}

	num_toll_booths = 30
	for which_booth <= num_toll_booths {
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
