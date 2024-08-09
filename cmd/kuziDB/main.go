package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvs()

	address, _ := os.LookupEnv("DEFAULT_ADDRESS")
	if address == "" {
		log.Fatal("Missing required environment variable: DEFAULT_ADDRESS")
		return
	}
	log.Printf("Will listen on %s", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("Can't start listen on %s by %s", address, err)
	}

	for {
		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		if err != nil {
			log.Fatalf("Can't start listen on %s by %s", address, err)
			return
		}

		go handleConnection(conn) // обрабатываем запросы клиента в отдельной го-рутине
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // закрываем сокет при выходе из функции

	buf := make([]byte, 128) // буфер для чтения клиентских данных
	for {
		readLen, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			log.Fatal(append([]byte("Can't read from socket"), err.Error()...))
			return
		}

		_, err = conn.Write(buf[:readLen])
		if err != nil {
			log.Fatal("Can't write to socket")
			return
		} // пишем в сокет
	}
}

func loadEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}
