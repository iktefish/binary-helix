package utils

import (
	"fmt"
	"io/ioutil"
	"log"
)

// This func must be Exported, Capitalized, and comment added.
func Reader(path string) {

	genome, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	processed, lineCount := preProccess(genome)
	fmt.Println("lineCount ~~> ", <-lineCount)
	fmt.Println("processed ~~> ", <-processed)
	Carrier(processed, <-lineCount)
}

func preProccess(b []byte) (<-chan []byte, <-chan int) {

	outByte := make(chan []byte)
	outInt := make(chan int)
	go func() {
		counter := 0
		for i, n := range b {
			if n == byte(62) { // Check for ">"
				for _, test := range b[i:] {
					if test == byte(10) { // Chech for "\n"
						break
					}
					counter += 1
				}
			}

			if counter == 0 {
				myArr := b[i+1:] // The +1 is there to skip the trailing \n
				lineCount := 0
				for _, breakChar := range b[i : len(b)-1] { // The len(b)-1 is there to skip the final \n
					if breakChar == byte(10) {
						lineCount += 1
					}
				}
				go func() {
					outByte <- myArr
				}()

				go func() {
					outInt <- lineCount
				}()

			} else {
				counter -= 1
			}
		}
	}()

	// time.Sleep(500 * time.Millisecond) // NOTE: For debugging and testing!

	return outByte, outInt
}

// func Carrier(c <-chan []byte, lc int) bool {
// 	go func() {
// 		gs := <-c
// 		nodeCount := 2 // TODO: Create a function that gets nodeCount
// 		// 1. Split for fasta
// 		counter := 1
// 		start := 0
// 		for i, b := range gs {
// 			if b == byte(10) {
// 				if counter == lc/nodeCount {
// 					fmt.Println("start ~~> ", start)
// 					fmt.Println("i ~~> ", i)
// 					split := gs[start:i]
//
// 					if len(gs[i:]) < lc/nodeCount {
// 						split = gs[start:]
// 					}
//
// 					fmt.Println("string(split) ~~> ", strings.TrimSpace(strings.TrimSuffix(string(split), "\n"))) // Trim everything 😂
//
// 					go func() {
// 						tmpFile, err := ioutil.TempFile("./tmp/", "prefix-")
// 						if err != nil {
// 							log.Fatal("Cannot create temp file", err)
// 						}
//
// 						if _, err = tmpFile.Write(split); err != nil {
// 							log.Fatal("Failed to write to temp file", err)
// 						}
//
// 						defer os.Remove(tmpFile.Name())
// 					}()
// 					counter = 0
// 					start = i
// 				}
// 				counter += 1
// 			}
// 		}
//
// 		// 2. Split for fastq
// 	}()
// 	time.Sleep(500 * time.Millisecond) // NOTE: For debugging and testing!
// 	return true
// }
