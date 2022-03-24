package main

/* This program is designed to work on the infamous 3n + 1 problem until
 * a shutdown signal is recieved (in the form of the creation of a /var/opt/3n1data file)
 * Upon recieving the shutdown signal, it saves it's work so that it can pick back up later
 */

import (
	"os"
	"errors"
	"log"
	"encoding/json"
	"fmt"
)

// Maths //

func threeNPlusOne(n int) int {
	if n % 2 == 0 {
		return n / 2
	} 

	n = (3 * n + 1)

	for n % 2 == 0 {
		n = n / 2
	}
	return n
}

func testCollatz(n int) []int {
	sequence := []int {n}
	for !in(sequence[1:], n) {
		n = threeNPlusOne(n)
		sequence = append([]int {n}, sequence...)
	}
	return sequence
}

func in(s []int, n int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

// System //

func restoreState() int {
	

	_, val := decodeData()
	err := os.Remove("/var/opt/3n1shutdown")
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func shouldShutdown() bool {
	_, err := os.Stat("/var/opt/3n1shutdown")
	return !errors.Is(err, os.ErrNotExist)
}

func shutdownSave(s []int, n int) {
	jsonFile, err3 := os.OpenFile("/var/opt/3n1data", os.O_RDWR, 0644)
	if err3 != nil {
		log.Fatal(err3)
	}
	defer jsonFile.Close()

	enc := json.NewEncoder(jsonFile)

	err4 := enc.Encode(s)
	if err4 != nil {
		log.Fatal(err4)
	}

	err := enc.Encode(n) 
	if err != nil {
		log.Fatal(err)
	}
}

func decodeData() ([]int, int) {
	jsonFile, err := os.OpenFile("/var/opt/3n1data", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	dec := json.NewDecoder(jsonFile)
	
	var s []int
	err0 := dec.Decode(&s)
	if err0 != nil {
		log.Fatal(err0)
	}

	var n int
	err1 := dec.Decode(&n)
	if err1 != nil {
		log.Fatal(err1)
	}

	return s, n
}

// Main //

func main() {

	n := restoreState()
	fmt.Println(n)
	s := []int{1}
	for s[0] == 1 {
		s = testCollatz(n)
		n ++

		if shouldShutdown() {
			break
		}
		fmt.Println(n, s)
	}
	shutdownSave(s, n)
}