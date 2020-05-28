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
func scan(ip string, port string, scantype string)result{
	var status string
	conn, err := net.Dial(scantype, ip+":"+port)
	if err!=nil{
		status = "Closed"
		return result{port,status}
	}
	defer conn.Close()
	status = "Open"
	return result{port, status}
}

func scanTop(ip string) [] result{
	var finalres []result
	for i:=79;i<=181;i++{
		resultx := scan(ip,strconv.Itoa(i),"tcp")
		finalres = append(finalres,resultx)
		fmt.Printf("%v \n",resultx)
	}
	return finalres
}
func main(){
	fmt.Println(scanTop("192.168.1.1"))
}
//func scanAll(ip string){
//
//}