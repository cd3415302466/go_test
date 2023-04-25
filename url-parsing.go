package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:password@localhost:5432/path?arg1=val1&arg2=val2"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme) //postgres

	fmt.Println(u.User)            //user:password
	fmt.Println(u.User.Username()) //user
	P, _ := u.User.Password()      //password
	fmt.Println(P)

	fmt.Println(u.Host) //localhost:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host) //localhost
	fmt.Println(port) //5432

	fmt.Println(u.Path) //path
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery) //arg1=val1&arg2=val2
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)         //map[arg1:[val1] arg2:[val2]]
	fmt.Println(m["arg1"]) //[val1]
}
