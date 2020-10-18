package main

import (
	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
	"log"
	"net"
	"time"
	//"chat"

	"google.golang.org/grpc"
	"fmt"
	"github.com/streadway/amqp"
	"encoding/json"
	"strconv"


)


type Entrega struct {
	Id_paquete string
	Tipo string
	Valor int
	Intentos int	
	Estado string
}

type PaqueteEnMarcha struct {
	Idpaquete     string
	Estado        string
	Idcamion      string
	Idseguimiento string
	Intentos      string
	Origen        string
	Destino       string
	Timestamp     time.Time
}


type Registro struct {
	Timestamp   time.Time
	Idpaquete   string
	Tipo        string
	Nombre      string
	Valor       string
	Origen      string
	Destino     string
	Seguimiento string
}


func InArr(id string, arr []string) bool {
    for _, i := range arr {
        if i == id {
            return true
        }
    }
    return false
}

func tipoYvalor (idpaquete string, 
	listaRegistro []chat.Registro,
				) (string, int){

	var tipo string
	var valor int
	for _, registro := range listaRegistro {
		if registro.Idpaquete == idpaquete {
			tipo = registro.Tipo
			v, _ := strconv.Atoi(registro.Valor )
			valor = v
            return tipo, valor
        }
	}

	return tipo, valor
}

func procesarEntregas(paquetesProcesados []string, 
				   	 paqueteEnMarcha []chat.PaqueteEnMarcha, 
					 listaRegistro []chat.Registro) ([]string , []Entrega) {
	
	var entregasProcesadas []Entrega

	for _, Paquete := range paqueteEnMarcha {

		IntIntentos, _ := strconv.Atoi(Paquete.Intentos)
		if ((Paquete.Estado == "Recibido" || Paquete.Estado == "No Recibido" ) && !InArr( Paquete.Idpaquete, paquetesProcesados)) {
			
			paquetesProcesados = append(paquetesProcesados, Paquete.Idpaquete)
			
			var tipo string
			var valor int
			tipo, valor = tipoYvalor(Paquete.Idpaquete, listaRegistro) 

			ent := &Entrega{Id_paquete: Paquete.Idpaquete, 
							Tipo: tipo, 
							Valor: valor, 
							Intentos: IntIntentos,
							Estado: Paquete.Estado}		

			entregasProcesadas = append(entregasProcesadas, *ent)    		
		}
	}	
	return paquetesProcesados, entregasProcesadas
}


func enviarRabbit(entregasProcesadas []Entrega) {
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

	
	for _, entrega := range entregasProcesadas  {

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
	}
}

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

	var paquetesProcesados []string
	var entregasProcesadas []Entrega

	go func() {
		for {
			time.Sleep(2 * time.Second)
			
			paquetesProcesados, entregasProcesadas = procesarEntregas(paquetesProcesados,
																	  s.PaquetesEnMarcha,
																	  s.ListaRegistro)

			fmt.Println(s.PaquetesEnMarcha) //Borrar................
			fmt.Println(s.ListaRegistro)
			fmt.Println("____________") //Borrar................

			enviarRabbit(entregasProcesadas)
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9000: %v", err)
	}

	if err2 := grpcServer.Serve(lis2); err2 != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9001: %v", err)
	}

}
