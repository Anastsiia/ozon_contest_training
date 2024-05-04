package main_test

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	testCases := getTestCases("3_cases")
	fmt.Println("3333333")
	for _, tc := range testCases {
		start := time.Now()

		inputFile, err := os.Open(tc.inputFile)
		if err != nil {
			t.Fatalf("error opening input file: %v", err)
		}
		defer inputFile.Close()

		expectedOutputFile, err := os.Open(tc.expectedOutputFile)
		if err != nil {
			t.Fatalf("error opening expected output file: %v", err)
		}
		defer expectedOutputFile.Close()

		reader := bufio.NewReader(inputFile)
		writer := bufio.NewWriter(expectedOutputFile)

		// Modify reader and writer to be of type *os.File
		os.Stdin = inputFile
		os.Stdout = expectedOutputFile

		main()

		err = writer.Flush()
		if err != nil {
			t.Fatalf("error flushing output: %v", err)
		}

		expectedOutputScanner := bufio.NewScanner(expectedOutputFile)
		outputScanner := bufio.NewScanner(os.Stdout)

		for expectedOutputScanner.Scan() && outputScanner.Scan() {
			expectedOutput := expectedOutputScanner.Text()
			output := outputScanner.Text()

			if expectedOutput != output {
				t.Errorf("expected output: %s, got: %s", expectedOutput, output)
			}
		}

		if expectedOutputScanner.Scan() || outputScanner.Scan() {
			t.Errorf("expected and actual outputs have different number of lines")
		}

		fmt.Printf("Test case %s passed in %v\n", tc.inputFile, time.Since(start))
	}
}

func getTestCases(testDir string) []struct {
	inputFile          string
	expectedOutputFile string
} {
	var testCases []struct {
		inputFile          string
		expectedOutputFile string
	}

	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Mode().IsRegular() {
			if filepath.Ext(path) == ".a" {
				base := filepath.Base(path)
				testCase := base[:len(base)-len(filepath.Ext(base))]

				inputFile := filepath.Join(testDir, testCase)
				expectedOutputFile := path

				testCases = append(testCases, struct {
					inputFile          string
					expectedOutputFile string
				}{inputFile, expectedOutputFile})
			}
		}

		return nil
	})

	if err != nil {
		panic(fmt.Errorf("error reading test directory: %v", err))
	}

	return testCases
}
