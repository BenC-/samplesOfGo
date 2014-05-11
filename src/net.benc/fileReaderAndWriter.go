/**
 * Example of reading and writing to file
 * This example uses bufio and io/ioutil packages
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FILENAME string = "/Users/benjaminchenebault/myfile.txt"

func main() {

	/* First method : Use scanner and read line by line */
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("File not found")
	}
	defer file.Close()

	var fileContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent += scanner.Text()
	}

	fmt.Println(fileContent)

	/* Second method : Load full file content at once */
	binaryFileContent, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatal("File not found")
	}

	fmt.Println(string(binaryFileContent))
}
