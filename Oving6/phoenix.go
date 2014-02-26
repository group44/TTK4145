package main

import (
	"fmt"
	"os/exec"
	"os"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func SendHeartbeat(conn *net.UDPConn, addr *net.UDPAddr) {
	conn.WriteToUDP([]byte, addr)
}

func CheckHeartbeat(conn *net.UDPConn) {
	
}


var data = []byte("test")
var backup, master bool

func main() {
	var name string
	
	fmt.Println(len(os.Args))
	
	if len(os.Args) == 2 {
		name = "127.0.0.1:15000"
		fmt.Println("test1")
		backup = true
		master = false
	} else {
		name = "129.241.187.255:15000"
		fmt.Println("test2")
		backup = false
		master = true
	}

	addr, err := net.ResolveUDPAddr("udp", name)
	conn, err := net.ListenUDP("udp", addr)
	CheckError(err)
	fmt.Println("test3")
	
	conn.WriteToUDP(data, addr)
	buf := make([]byte, 512)
	conn.ReadFromUDP(buf)
	fmt.Println("test4")
	
	
	
	for {
		switch backup {
		case true:
			fmt.Println("test5")
		case false:
			cmd := exec.Command("mate-terminal", "-x", "go", "run",  "phoenix.go", "slave")
			err = cmd.Start()
			CheckError(err)
			backup = true
			master = true
			fmt.Println("test6")
		}
		
		
	}
	
	fmt.Println("end")
}
