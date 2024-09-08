package main

import(
	"fmt"
	"flag"
)

func main(){
	var host *string = flag.String("host", "localhost", "database hostname") //* name, default, description
	var port *int = flag.Int("port", 0, "database port") //* name, default, description
	var username *string = flag.String("username", "root", "database username") //* name, default, description
	var password *string = flag.String("password", "root", "database password") //* name, default, description

	flag.Parse()

	fmt.Println("hostname", *host)
	fmt.Println("port", *port)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
}