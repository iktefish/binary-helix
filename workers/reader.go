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
	if utils.Verify_FileExt(path) != true {
		log.Fatal("FAIL [PANIC]: Please enter path to a '.fa' or '.fastq' file!")
	}

	genome, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("FAIL: Please enter a proper path to a file!")
		utils.HandleError(err)
	}

	var processed []byte
	var line_count int

	if utils.Verify_Fasta(path) {
		processed, line_count = pre_process_fa(genome)
	}

	if utils.Verify_Fastq(path) {
		processed, line_count = pre_process_fq(genome)
	}

	return filepath.Ext(path), processed, line_count
}

func pre_process_fa(b []byte) ([]byte, int) {
	var wg sync.WaitGroup
	wg.Add(1)

	line_count := 0
	anchor := 0

	var output []byte

	go func() {
		defer wg.Done()

		for i, bc := range b {
			if bc == byte(62) {
				anchor += 1
			}

			if bc == byte(10) {
				line_count += 1
				if anchor != 0 {
					output = b[i:]
					anchor = 0
				}
			}
		}
	}()

	wg.Wait()
	return output, line_count
}

func pre_process_fq(b []byte) ([]byte, int) {
	var wg sync.WaitGroup
	wg.Add(1)

	line_count := 0
	anchor := 0

	var output []byte

	go func() {
		defer wg.Done()

		for i, bc := range b {
			if bc == byte(10) {
				line_count += 1

				if line_count%2 != 0 {
					anchor = i
				}

				if line_count%2 == 0 {
					output = append(output, b[anchor:i]...)
				}
			}
		}
	}()

	wg.Wait()
	return output, line_count
}
