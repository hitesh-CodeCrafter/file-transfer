# Go TCP File Transfer

This project demonstrates a simple TCP-based file transfer system in Go, where files can be streamed from a client to a server over a network. The server listens for incoming connections and saves the received file to a specified directory, while the client reads a file from the local machine and sends it to the server.

## Folder Structure

Go-TCP-File-Transfer/
├── cmd/
│ └── server.go
└── client/
└── client.go

markdown


- **cmd/server.go**: Contains the code for the TCP server that receives files from clients.
- **client/client.go**: Contains the code for the TCP client that sends files to the server.

## Prerequisites

- Go 1.16 or higher installed on your machine.
- Basic understanding of Go and TCP/IP networking.

## Setup Instructions

### 1. Clone the Repository

```sh
git clone https://github.com/your-username/Go-TCP-File-Transfer.git
cd Go-TCP-File-Transfer

2. Running the Server

Navigate to the cmd directory and run the server:

sh

cd cmd
go run server.go

The server will start listening on port 8080 for incoming connections.
3. Running the Client

Navigate to the client directory and run the client:

sh

cd client
go run client.go

Before running the client, ensure that you've set the correct file path in client.go and the server's IP address.
4. File Transfer Process

    Server: The server will accept connections on port 8080 and wait for the client to send a file. It will save the file as received_file.txt in the server's working directory (you can modify the path as needed).

    Client: The client reads a file specified in the code and streams it to the server over a TCP connection.

5. Example Usage

Modify the file paths in client.go and server.go as necessary:

    Client: Specify the path of the file you want to send and the server's IP address.
    Server: Specify the path where you want to save the received file.

6. Customization

You can customize the following aspects of the project:

    Port Number: Change the port number in both server.go and client.go if you need to use a different port.
    File Paths: Modify the paths in the client.go and server.go to change the location of the source and destination files.
    Buffer Size: Adjust the buffer size used for reading and writing data in chunks if you need to optimize performance.

Error Handling

    If the connection between the client and server fails, appropriate error messages will be displayed.
    Ensure that the server is running before starting the client, otherwise, the client will be unable to connect.