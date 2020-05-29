package main

import (
	"fmt"
	"net"
	"strconv"
)
type result struct {
	port string
	status string
}

func scan(ip string, port string, scantype string) result {
	conn, err := net.Dial(scantype, ip+":"+port)
	if err!=nil{
		return result{port,"Closed"}
	}
	defer conn.Close()
	return result{port,"Open"}
}
func scanTop(ip string) []result {
	var finalres []result
	var ch = make(chan result)
	go func() {
		for i := 1; i <= 1024; i++ {
			results := scan(ip, strconv.Itoa(i), "tcp")
			fmt.Printf("%v \n", results)
			ch <- results
			finalres = append(finalres, results)
		}
		ch <- result{status:"q",port:"q"}
	}()
	for res := range ch {
		finalres := append(finalres, res)
		if res.status == "q" {
			return finalres
		}
	}
	return finalres
}

func main(){
	fmt.Println(scanTop("192.168.1.1"))
}
//func scanAll(ip string){
//
//}