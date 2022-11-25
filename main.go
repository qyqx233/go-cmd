package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var host, password, database, user, path, raw string
	var port int
	switch os.Args[1] {
	case "telnet":
		telnetCommand := flag.NewFlagSet("telnet", flag.ExitOnError)
		telnetCommand.StringVar(&host, "host", "localhost", "host")
		telnetCommand.IntVar(&port, "port", 80, "port")
		telnetCommand.Parse(os.Args[2:])
		telnet(host, port)
	case "sha256":
		sha256Command := flag.NewFlagSet("sha256", flag.ExitOnError)
		sha256Command.StringVar(&path, "path", "", "path")
		sha256Command.StringVar(&raw, "raw", "", "raw")
		sha256Command.Parse(os.Args[2:])
		sha256_(path, raw)
	case "mysql":
		mysqlCommand := flag.NewFlagSet("mysql", flag.ExitOnError)
		mysqlCommand.StringVar(&host, "host", "localhost", "host")
		mysqlCommand.StringVar(&database, "db", "database", "database")
		mysqlCommand.StringVar(&user, "user", "root", "user")
		mysqlCommand.StringVar(&password, "pass", "password", "password")
		mysqlCommand.IntVar(&port, "port", 3306, "port")
		mysqlCommand.Parse(os.Args[2:])
		mysql(host, port, user, password, database)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
