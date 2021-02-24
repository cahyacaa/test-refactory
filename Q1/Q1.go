package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strconv"
	"strings"
)
func rightjust(s string, n int, fill string) string {
	return strings.Repeat(fill, n) + s
}




func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner2 := bufio.NewScanner(os.Stdin)
	var arr []string
	var arr2 []string
	for {
		fmt.Print("Enter Items: ")
		scanner.Scan()
		items := scanner.Text()		
		fmt.Print("Enter Price: ")
		scanner2.Scan()
		price := scanner2.Text()
		if len(items) != 0 {
			arr  =  append(arr, items)
			arr2 =  append(arr2, price)
		} else {
			break
		}

	}
	if len(arr) == len(arr2) {
		name:="Cahya Nugroho"
		total_price:=0
		var total_price1 string
		dt := time.Now()
		fmt.Println("    Warung Makan Sederhana   ")
		fmt.Println("Tanggal : ",dt.Format("2006-01-02 15:04:05"))
		fmt.Println( "Nama Kasir",rightjust(name,(29-(len("Name Kasir")+len(name))), ".")  )
		fmt.Println(strings.Repeat("=",30))
        for i := range arr {
			if x, err := strconv.Atoi(arr2[i]); err == nil {
				total_price+=x
			}
			total_price1 = strconv.Itoa(int(total_price))
			fmt.Println(arr[i],rightjust("Rp."+arr2[i],(26-(len(arr[i])+len(arr2[i]))), ".") )
			
		}
		fmt.Println(" ")
		fmt.Println("Total",rightjust("Rp."+total_price1,(26-(len("Total")+len(total_price1))), ".") )
    }else{
		fmt.Println("There is data not valid")
	}
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}