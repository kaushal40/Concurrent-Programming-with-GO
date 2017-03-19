// Asynchronous_service_call
package main

import (
	"fmt"
	//to make web requests
	"net/http"
	//read the response from call
	"io/ioutil"
	//unmarshall the response
	"encoding/xml"
	//time package to check how much each call is taking
	"time"
)

//main function should know till what time it should run so that it will not terminate before the goroutines completed
// you can use number of CPU if you know the nuber of cores your machine is running on. to take advantage of parralism
//runtime.GOMAXCOREPROC("number of CPU")
func main() {

	//tracking variables for go routine
	numOfCompletedRoutine := 0

	companynames := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
	}

	start := time.Now()

	for _, i := range companynames {

		// added sumbol==i so each goroutines are decoupled from for loop. Because for loop will not wait for other go routines to execute
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + i)
			//close the connection
			defer resp.Body.Close()
			//this should be done carefully as if data is big it can blow up the system but in this example data is small so we can read that in single byte stream
			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Println("Company Name: = ", quote.Name, "Last Price:=", quote.LastPrice)
			numOfCompletedRoutine++
		}(i)
	}

	for numOfCompletedRoutine < len(companynames) {
		time.Sleep(10 * time.Millisecond)
	}

	timeTakenForCall := time.Since(start)
	fmt.Println("Execution time is: ", timeTakenForCall)
}

type QuoteResponse struct {
	Status    string
	Name      string
	LastPrice string
}
