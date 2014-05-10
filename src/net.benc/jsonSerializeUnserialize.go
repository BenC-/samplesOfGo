/**
 * Example of marshalling and unmarshalling JSON
 * Channel, complex, and function types cannot be encoded
 */
package main

import (
	"encoding/json"
	"fmt"
)

type HttpRequest struct {
	Method string
	Body   string
}

func main() {

	// Serialize in JSON
	originatedRequest := HttpRequest{"POST", "Here is my content"}
	b, err := json.Marshal(originatedRequest)
	if err != nil {
		fmt.Println("Error while marshalling")
	}

	// Unserialize from JSON
	b = []byte(`{"Method":"POST","Body":"Here is my content"}`)
	var terminatedRequest HttpRequest
	err = json.Unmarshal(b, &terminatedRequest)
	if err != nil {
		fmt.Println("Error while unmarshalling")
	}

	if terminatedRequest == originatedRequest {
		fmt.Println("SUCCESS")
	}

}
