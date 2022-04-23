package workers

import (
	"fmt"
	"github.com/iktefish/binary-helix/nodes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func Carrier(c <-chan []byte, lc int) bool {
	go func() {
		gs := <-c
		nodeCount := 2 // TODO: Create a function that gets nodeCount. TODO: Use WaitGroups.
		// 1. Split for fasta
		counter := 1
		start := 0
		for i, b := range gs {
			if b == byte(10) {
				if counter == lc/nodeCount {
					fmt.Println("start ~~> ", start)
					fmt.Println("i ~~> ", i)
					split := gs[start:i]

					if len(gs[i:]) < lc/nodeCount {
						split = gs[start:]
					}

					fmt.Println("string(split) ~~> ", strings.TrimSpace(strings.TrimSuffix(string(split), "\n"))) // Trim everything 😂

					go func() {
						tmpFile, err := ioutil.TempFile("./tmp/", "prefix-")
						if err != nil {
							log.Fatal("Cannot create temp file", err)
						}

						if _, err = tmpFile.Write(split); err != nil {
							log.Fatal("Failed to write to temp file", err)
						}

						nodes.Client()

						time.Sleep(1500 * time.Millisecond) // NOTE: For debugging and testing!

						defer os.Remove(tmpFile.Name())
					}()
					counter = 0
					start = i
				}
				counter += 1
			}
		}

		// 2. Split for fastq
	}()
	time.Sleep(500 * time.Millisecond) // NOTE: For debugging and testing!
	return true
}
