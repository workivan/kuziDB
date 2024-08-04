package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
)

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
