package main

import(
	"flag"
	"fmt"
	"github.com/pbssubhash/goportscan/pkg"
)

func main(){
	verbose := flag.Bool("v", false,"Enable Verbose Output" )
	flag.Parse()
	printVersionAndDefaults()
	if *verbose {
		fmt.Println("[+] Verbose Mode enabled. [+]")
	}
}

func printVersionAndDefaults(){
	banner := `

====================================================================================
  ________                __________              __   _________                     
 /  _____/  ____          \______   \____________/  |_/   _____/ ____ _____    ____  
/   \  ___ /  _ \   ______ |     ___/  _ \_  __ \   __\_____  \_/ ___\\__  \  /    \ 
\    \_\  (  <_> ) /_____/ |    |  (  <_> )  | \/|  | /        \  \___ / __ \|   |  \
 \______  /\____/          |____|   \____/|__|   |__|/_______  /\___  >____  /___|  /
        \/                                                   \/     \/     \/     \/ 
=====================================================================================
Go-PortScan Version 1.0
`
	fmt.Println(banner)
	flag.PrintDefaults()
}
