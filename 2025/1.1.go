package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://adventofcode.com/2025/day/1

func AOC_1_1() {
	f, err := os.Open("./input/1.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pos := 50
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		v, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		v %= 100
		if line[0:1] == "L" {
			pos -= v
			if pos < 0 {
				pos += 100
			}
		} else {
			pos += v
			if pos > 99 {
				pos -= 100
			}
		}
		if pos == 0 {
			count += 1
		}
		// fmt.Printf("%d\n", pos)
	}
	fmt.Printf("%d\n", count)
}
