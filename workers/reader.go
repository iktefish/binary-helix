package workers

import (
	"fmt"
	"io/ioutil"
	"sync"

	"log"
	"path/filepath"

	"github.com/iktefish/binary-helix/utils"
)

func Reader(path string) (string, []byte, int) {
	if utils.Verify_FileExt(path) {
		log.Fatal("FAIL [PANIC]: Please enter path to a '.fa' or '.fastq' file!")
	}

	genome, err := ioutil.ReadFile(path)
	if err != nil {
        fmt.Println("FAIL: Please enter a proper path to a file!")
		utils.HandleError(err)
	}

	var processed []byte
	var lineCount int

	if filepath.Ext(path) == ".fa" {
		processed, lineCount = preProcess_fa(genome)
	}

	if filepath.Ext(path) == ".fastq" {
		processed, lineCount = preProcess_fq(genome)
	}

	return filepath.Ext(path), processed, lineCount
}

func preProcess_fa(b []byte) ([]byte, int) {
	var wg sync.WaitGroup
	wg.Add(1)

	lineCount := 0
	anchor := 0

	var output []byte

	go func() {
		defer wg.Done()

		for i, bc := range b {
			if bc == byte(62) {
				anchor += 1
			}

			if bc == byte(10) {
				lineCount += 1
				if anchor != 0 {
					output = b[i:]
					anchor = 0
				}
			}
		}
	}()

	wg.Wait()
	return output, lineCount
}

func preProcess_fq(b []byte) ([]byte, int) {
	var wg sync.WaitGroup
	wg.Add(1)

	lineCount := 0
	anchor := 0

	var output []byte

	go func() {
		defer wg.Done()

		for i, bc := range b {
			if bc == byte(10) {
				lineCount += 1

				if lineCount%2 != 0 {
					anchor = i
				}

				if lineCount%2 == 0 {
					output = append(output, b[anchor:i]...)
				}
			}
		}
	}()

	wg.Wait()
	return output, lineCount
}
