package main

import (
	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
	"log"
	"net"
	"time"
	//"chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Falle al escuchar puerto 9000: %v", err)
	}

	lis2, err2 := net.Listen("tcp", ":9001")
	if err2 != nil {
		log.Fatalf("Falle al escuchar puerto 9001: %v", err)
	}

	var listaRegistro []chat.Registro
	s := chat.Server{
		//chat.listaRegistro: listaRegistro,
		//seguimiento:        "0",
	}
	s.ListaRegistro = listaRegistro
	s.Seguimiento = "0"

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Printf("Paquetes:\n")
			log.Println(s.ColaRetail)
			log.Printf("\n")
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9000: %v", err)
	}

	if err2 := grpcServer.Serve(lis2); err2 != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9001: %v", err)
	}

}
