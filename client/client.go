package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open("/media/altair/Fade/animes/black clover/Black Clover (TV) (Sub) Episode 077 - Watch Black Clover (TV) (Sub) Episod.mp4") // Replace with the file you want to send
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Stream the file to the server in chunks
	buffer := make([]byte, 4096) // 4KB buffer
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			return
		}
		if n == 0 {
			break
		}

		// Write the chunk to the connection
		if _, err := conn.Write(buffer[:n]); err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}

	fmt.Println("File sent successfully")
}
