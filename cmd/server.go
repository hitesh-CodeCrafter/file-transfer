package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn, savePath string) {
	defer conn.Close()

	// Create or open a file to save the received data
	outFile, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Read the data in chunks
	buffer := make([]byte, 4096) // 4KB buffer
	for {
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from connection:", err)
			return
		}
		if n == 0 {
			break
		}

		if _, err := outFile.Write(buffer[:n]); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File received successfully")
}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, "/home/altair/go/src/file-transfer/output.mp4")
	}
}
