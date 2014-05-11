/**
 * Example of marshalling and unmarshalling XML
 * For more information, go to http://golang.org/pkg/encoding/xml
 */
package main

import (
	"encoding/xml"
	"fmt"
)

type HttpRequest struct {
	XMLName xml.Name `xml:"httpRequest"` // Override xml root element name
	Method  string   `xml:"method"`
	Body    string   `xml:"body"`
}

func main() {

	/* Serialize in XML */
	originatedRequest := HttpRequest{Method: "POST", Body: "Here is my content"}
	b, err := xml.Marshal(originatedRequest)
	if err != nil {
		fmt.Println("Error while marshalling")
	}

	fmt.Println("Binary representation from struct : " + string(b))

	/* Unserialize XML */
	b = []byte(`<httpRequest><method>POST</method><body>Here is my content</body></httpRequest>`)

	fmt.Println("Compared with its equivalent byte [] : " + string(b))

	var terminatedRequest HttpRequest
	err = xml.Unmarshal(b, &terminatedRequest)
	if err != nil {
		fmt.Println("Error while unmarshalling")
	}

	/* Check result */
	fmt.Print("Initial struct object : ")
	fmt.Println(originatedRequest)
	fmt.Print("Struct object after xmlisation : ")
	fmt.Println(terminatedRequest)

}
