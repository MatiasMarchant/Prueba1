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
	//Origen string
	//Destino string
	Intentos int	
	//Fecha_entrega time.Time
	Estado string
}



type Cola struct {
	Idpaquete   string
	Seguimiento string
	Tipo        string
	Valor       string
	Intentos    string
	Estado      string
	Origen      string
	Destino     string
}

func InArr(id string, arr []string) bool {
    for _, i := range arr {
        if i == id {
            return true
        }
    }
    return false
}

func iterarCola(paquetesProcesados []string, cola []chat.Cola, entregasProcesadas []Entrega) ([]string, []Entrega){
	
	for _, Paquete := range cola {

		if ((Paquete.Estado == "Recibido" || Paquete.Estado == "No Recibido" ) && !InArr( Paquete.Idpaquete, paquetesProcesados)) {
			
			IntIntentos, _ := strconv.Atoi(Paquete.Intentos)
			IntValor, _ := strconv.Atoi(Paquete.Valor)		

			paquetesProcesados = append(paquetesProcesados, Paquete.Idpaquete)
	
			ent := &Entrega{Id_paquete: Paquete.Idpaquete, 
							Tipo: Paquete.Tipo, 
							Valor: IntValor, 
							Intentos: IntIntentos,
							Estado: Paquete.Estado}		

			entregasProcesadas = append(entregasProcesadas, *ent)    		
		}
	}

	return paquetesProcesados, entregasProcesadas
	
}


func procesarEntregas(paquetesProcesados []string, 
					 colaRetail []chat.Cola,
					 colaPrioritario []chat.Cola,
					 colaNormal []chat.Cola) ([]string , []Entrega) {
	
	var entregasProcesadas []Entrega

	paquetesProcesados, entregasProcesadas = iterarCola(paquetesProcesados, colaRetail, entregasProcesadas) 
	paquetesProcesados, entregasProcesadas = iterarCola(paquetesProcesados, colaPrioritario, entregasProcesadas) 
	paquetesProcesados, entregasProcesadas = iterarCola(paquetesProcesados, colaNormal, entregasProcesadas) 
	

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
																	  s.ColaRetail,
																	  s.ColaPrioritario,
																	  s.ColaNormal,
																	)

			fmt.Println(s.PaquetesEnMarcha) //Borrar................
			fmt.Println(paquetesProcesados) //Borrar................
			fmt.Println(entregasProcesadas) //Borrar................
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
