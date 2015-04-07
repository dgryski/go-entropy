package main

import (
	"bufio"
	"fmt"
	"os"

	//"github.com/davecheney/profile"
	"github.com/dgryski/go-entropy"
)

func main() {

	//defer profile.Start(profile.CPUProfile).Stop()

	ex := entropy.NewExact()
	sk := entropy.NewEstimate(100)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ex.Push(scanner.Bytes(), 1)
		sk.Push(scanner.Bytes(), 1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error during scan: ", err)
	}

	fmt.Println("entropy :", ex.Entropy())
	fmt.Println("estimate:", sk.Entropy())
}
