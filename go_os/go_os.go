package main

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

func main() {
	fmt.Println("os.Args ", os.Args)
	fmt.Println("os.Args 0", os.Args[0])
	fmt.Println("os.Args 1", os.Args[1])
	fmt.Println("os.Args 2", os.Args[2])

	// Os.DevNull
	fmt.Println("os.DevNull ", os.DevNull)
	// os.Exit
	// osExit()

	//os.Hostname
	host, _ := os.Hostname()
	fmt.Println("os hostname ", host)

	// home directory
	homeDir, err := os.UserHomeDir()
	fmt.Println(homeDir, err)

	// current user info
	currentUser, err := user.Current()
	fmt.Printf("%#v\n", currentUser)
	fmt.Println(err)

	// get current working dir
	cwd, _ := os.Getwd()
	fmt.Println("current working directory", cwd)

	// change dir
	os.Chdir("..")
	cwd1, _ := os.Getwd()
	fmt.Println("current ch working directory", cwd1)
	// change dir
	os.Chdir(cwd)

}

func osExit() {

	time.AfterFunc(time.Second*2, func() {
		os.Exit(1)

	})
	<-time.Tick(time.Second * 3)

}
