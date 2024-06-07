# CandyServer

  

This project implements a candy vending machine server and a client, following a given Swagger specification for communication between the vending machine and the server. The project is implemented in Go and includes mutual TLS authentication and integration of C code.

This project was developed as a part of School 21 curriculum.

  

## Description

  

*  **Candy Purchase API**: The server handles candy purchase requests, validates input, calculates change, and returns responses in JSON format.

*  **Mutual TLS Authentication**: Both the server and the client use self-signed certificates for secure communication.

*  **C Function Integration**: The `ask_cow` C function is used to generate ASCII art responses.

  

## Prerequisites

  
-  **Golang** v 1.21

-  **Taskfile** v3.37

-  **OpenSSL** 1.1
  

## Usage

 - Navigate to src and build server and client executables

	 `task build`
 - Generate certificates

	 `task certs`
 - Run server

	 `task run-server`
 - In separate terminal, run client
 
	 `./candy-client -k AA -c 2 -m 50 -f`
	 
	 - -k: candy type
	 - -c: candy count
	 - -m: money amount
	 - -f: enable cowsay response
	 