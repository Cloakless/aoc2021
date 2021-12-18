package main

import "fmt"

func main() {
	// fishes := []int{3,4,3,1,2}
	fishes := []int{2,1,2,1,5,1,5,1,2,2,1,1,5,1,4,4,4,3,1,2,2,3,4,1,1,5,1,1,4,2,5,5,5,1,1,4,5,4,1,1,4,2,1,4,1,2,2,5,1,1,5,1,1,3,4,4,1,2,3,1,5,5,4,1,4,1,2,1,5,1,1,1,3,4,1,1,5,1,5,1,1,5,1,1,4,3,2,4,1,4,1,5,3,3,1,5,1,3,1,1,4,1,4,5,2,3,1,1,1,1,3,1,2,1,5,1,1,5,1,1,1,1,4,1,4,3,1,5,1,1,5,4,4,2,1,4,5,1,1,3,3,1,1,4,2,5,5,2,4,1,4,5,4,5,3,1,4,1,5,2,4,5,3,1,3,2,4,5,4,4,1,5,1,5,1,2,2,1,4,1,1,4,2,2,2,4,1,1,5,3,1,1,5,4,4,1,5,1,3,1,3,2,2,1,1,4,1,4,1,2,2,1,1,3,5,1,2,1,3,1,4,5,1,3,4,1,1,1,1,4,3,3,4,5,1,1,1,1,1,2,4,5,3,4,2,1,1,1,3,3,1,4,1,1,4,2,1,5,1,1,2,3,4,2,5,1,1,1,5,1,1,4,1,2,4,1,1,2,4,3,4,2,3,1,1,2,1,5,4,2,3,5,1,2,3,1,2,2,1,4}
	fish_each_day := [7]int{0,0,0,0,0,0,0}
	for i := range fishes {
		fish_each_day[fishes[i]] += 1
	}
	num_days := 256
	new_fish, day1, day2 := 0, 0, 0
	for i := 0; i < num_days; i++ {
		curr_day := i % 7
		new_fish, day1, day2 = day1, day2, fish_each_day[curr_day]
		fish_each_day[curr_day] += new_fish
	}
	acc := 0
	for j := range fish_each_day {
		acc += fish_each_day[j]
	}
	acc += day1 + day2
	fmt.Println(acc)
	// // Go through the generation each day
	// for i:= 0; i < num_days; i++ {
	// 	next_gen := []int{}
	// 	// fmt.Println(fishes)
	// 	for _, fish := range fishes {
	// 		if fish != 0 {
	// 			next_gen = append(next_gen, fish-1)
	// 		} else {
	// 			next_gen = append(next_gen, 6, 8)
	// 		}
	// 	}
	// 	fishes = next_gen
	// 	// fmt.Println(next_gen)
	// 	fmt.Println(i)
	// }
	// fmt.Println(len(fishes))
	
}
