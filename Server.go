package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8091")
	// accept connection on port
	conn, _ := ln.Accept()
	// run loop forever (or until ctrl-c)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Server command> ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		text = text[:len(text)-1]

		switch text {
		case "1":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Destiny: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			break
		case "2":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Destiny: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			break
		case "3":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Destiny: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			break
		case "4":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Command: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			parameter := bufio.NewReader(os.Stdin)
			fmt.Print("Parameter: ")
			tparameter, _ := parameter.ReadString('\n')
			fmt.Fprintf(conn, tparameter+"\n")
			break
		case "5":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message Received:", string(message))
			break
		default:
			break
		}
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Status command:", string(message))
	}
}
