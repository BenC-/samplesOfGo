/**
 * Example of computing sha1 of a string
 */
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	myStringToHash := "JDç!../qr"

	/* Generate a sha1 hasher */
	hasher := sha1.New()

	/* Append data to the running hash */
	hasher.Write([]byte(myStringToHash))

	/* Display the sha1 */
	fmt.Println(hasher.Sum(nil)) // Sum returns the SHA1 checksum of the data.

	/* Or directly compute the hash without storing data */
	hasher = sha1.New()
	fmt.Println(hasher.Sum([]byte(myStringToHash)))

}
