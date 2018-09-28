package main

import (
	"fmt"
	"time"
)

func main() {
	var inputString string
	var inputInteger int64
	start := time.Now()

	_, err := fmt.Scan(&inputString, &inputInteger)
	if err != nil {
		fmt.Println("Error with reading Input from Terminal.")
		fmt.Println("Input sample: <0x82125AFe01819Dff1535D0D6276d57045291B6c0 30>")
	}

	result, err := GetTransactions(inputString, inputInteger)
	if err != nil {
		fmt.Println("Error with getting the transactions")
	} else {
		fmt.Println("GetTransaction() was called succesfully ")
	}
	fmt.Println("The length of the returned list is: ", len(result))
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

// Sample Input: <0x82125AFe01819Dff1535D0D6276d57045291B6c0,60>
// User must set the key for the EtherScan API in GetTransactions()
