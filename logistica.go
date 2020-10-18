package main

import (
	//"github.com/MatiasMarchant/Prueba1/tree/master/chat"
	"log"
	"net"
	"time"
	//"chat"

	"google.golang.org/grpc"
	"fmt"
	//"github.com/streadway/amqp"
	"encoding/json"
	
	//"reflect" //-------------
	"github.com/MatiasMarchant/Prueba1/tree/centos4/chat"
)


import chat "github.com/MatiasMarchant/Prueba1/tree/centos4/chat"


type Entrega struct {
	Id_paquete int
	Tipo string
	Valor int
	Origen string
	Destino string
	Intentos int	
	Fecha_entrega int
}



func procesarEntregas(paquetesProcesados []int, ColaRetail []chat.Cola){ //([]int , []Entrega){
	
	fmt.Println(ColaRetail)	
	for _, Paquete := range ColaRetail {

		
		
		//fmt.Println(Paquete)	

		slcB, _ := json.Marshal(Paquete)
    	fmt.Println(string(slcB))
		
		//chat.Cola

		//entregaProcesada := &Entrega{Id_paquete:int, Tipo:str, Valor:int, Origen:str, Destino:str, Intentos:int, Fecha_entrega:int}

		/*for _, idPaquete := range paquetesProcesados {
			if idPaquete == entregaProcesada.Id_paquete {
				// Found!
			}
		}*/

	}





	//iterar ColaRetail
	
	//return paquetesProcesados, entregasProcesadas
}


/*func enviarRabbit(entregasProcesadas []Entrega) {
	conn, err := amqp.Dial("amqp://mqadmin:mqadminpassword@10.6.40.180:5672/")
	if err != nil {
		fmt.Println("Falla inicializando conección")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)
	if err != nil {
		fmt.Println(err)
	}

	//ent := &Entrega{Id_paquete:ranId, Tipo:tipos[randT], Valor:ranV, Origen:"tienda-A", Destino:"casa-A", Intentos:randI, Fecha_entrega:0}
	
	for _, entrega := entregasProcesadas  {

		js, err := json.Marshal(entrega)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ch.Publish(
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType:  "application/json",
				Body:         js,
			},
		)

		if err != nil {
			fmt.Println(err)
		}
}*/

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

	var paquetesProcesados []int

	go func() {
		for {
			time.Sleep(2 * time.Second)
			//paquetesProcesados, entregasProcesadas = procesarEntregas(paquetesProcesados, s.ColaRetail)
			procesarEntregas(paquetesProcesados, s.ColaRetail)
						
			//enviarRabbit(entregasProcesadas)
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9000: %v", err)
	}

	if err2 := grpcServer.Serve(lis2); err2 != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9001: %v", err)
	}

}
