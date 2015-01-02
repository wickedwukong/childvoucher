package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("ChildcareVoucher-template.htm")
	if err != nil {
		fmt.Println("Error reading template file")
		fmt.Printf("Error: %s", err)
		log.Fatal(err)
		panic(err)
	} else {
		for i := 4; i <= 12; i++ {
			fmt.Println("Value of i is now:", i)
			yearMonth := strconv.Itoa(201400 + i)
			voucherContents := strings.Replace(string(contents), "$voucherMonth$", yearMonth, 1)
			voucherFile, err := os.Create("ChildcareVoucher-" + yearMonth + ".html")

			defer voucherFile.Close()

			if err != nil {
				fmt.Println("Error creating a voucher file")
				fmt.Println("Error: %s", err)
				log.Fatal(err)
				panic(err)
			} else {
				voucherFile.WriteString(voucherContents)
				voucherFile.Sync()
			}
		}
	}
}
