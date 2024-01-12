package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func registraLog(site string, status bool, statusCode int) {
	response, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	statusString := strconv.Itoa(statusCode)
	if status == true {
		response.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online " + "Status Code: " + statusString + "\n")
	} else {
		response.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - offline " + "Status Code: " + statusString + "\n")
	}
}

func imprimeLogs() {
	fmt.Println("Exibindo logs...")
	response, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(response))
}
