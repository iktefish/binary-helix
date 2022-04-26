package workers

import (
	"github.com/iktefish/binary-helix/utils"
	"log"
	"strings"
	"sync"
)

func Splitter(p string, b []byte, lc int) []string {
	var wg sync.WaitGroup
	wg.Add(1)

	var splits []string

	nodeCount := utils.Get_ActiveNodes()
	if nodeCount == 0 {
		log.Fatal("FAIL [PANIC]: No compute nodes are active!")
	}

	if p == ".fa" {
		splits = splitFa(&wg, b, lc, nodeCount)
	}

	if p == ".fastq" {
		splits = splitFq(&wg, b, lc, nodeCount)
	}

	wg.Wait()
	return splits
}

func splitFa(wg *sync.WaitGroup, b []byte, lc int, n int) []string {
	defer wg.Done()

	var splits []string

	counter := 1
	start := 0
	for i, bc := range b {
		if bc == byte(10) {
			if counter == lc/n {
				split := b[start:i]

				if len(b[i:]) < lc/n {
					split = b[start:]
				}

				splits = append(splits, strings.TrimSpace(strings.TrimSuffix(string(split), "\n")))

				counter = 0
				start = i
			}
			counter += 1
		}
	}

	return splits
}

func splitFq(wg *sync.WaitGroup, b []byte, lc int, n int) []string {
	defer wg.Done()

	var splits []string

	var ideal int
	for i := lc / 2; i >= 1; i-- {
		if i%n == 0 {
			ideal = i
			break
		}
	}

	counter := 1
	start := 0
	for i, bc := range b {
		if bc == byte(10) {
			if counter == ideal {
				split := b[start:i]

				if len(b[i:]) < ideal {
					split = b[start:]
				}

				splits = append(splits, strings.TrimSpace(strings.TrimSuffix(string(split), "\n")))

				counter = 0
				start = i
			}
			counter += 1
		}
	}

	return splits
}
