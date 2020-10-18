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
	Origen string
	Destino string
	Intentos int	
	Fecha_entrega time.Time
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

func tipoYvalor (idpaquete string, 
				 colaRetail []chat.Cola,
				 colaPrioritario []chat.Cola,
				 colaNormal []chat.Cola) (string, int){

	var tipo string
	var valor int
	for _, cola := range colaRetail {
		if cola.Idpaquete == idpaquete {
			tipo = cola.Tipo
			v, _ := strconv.Atoi(cola.Valor )
			valor = v
            return tipo, valor
        }
	}
	for _, cola := range colaPrioritario {
		if cola.Idpaquete == idpaquete {
			tipo = cola.Tipo
			v, _ := strconv.Atoi(cola.Valor )
			valor = v
            return tipo, valor
        }
	}
	for _, cola := range colaNormal {
		if cola.Idpaquete == idpaquete {
			tipo = cola.Tipo
			v, _ := strconv.Atoi(cola.Valor )
			valor = v
            return tipo, valor
        }
	}

	return tipo, valor
}

func procesarEntregas(paquetesProcesados []string, 
				   	 paqueteEnMarcha []chat.PaqueteEnMarcha, 
					 colaRetail []chat.Cola,
					 colaPrioritario []chat.Cola,
					 colaNormal []chat.Cola) ([]string , []Entrega) {
	
	var entregasProcesadas []Entrega

	for _, Paquete := range paqueteEnMarcha {

		IntIntentos, _ := strconv.Atoi(Paquete.Intentos)
		if ((Paquete.Estado == "Recibido" || Paquete.Estado == "No Recibido" ) && !InArr( Paquete.Idpaquete, paquetesProcesados)) {
			
			paquetesProcesados = append(paquetesProcesados, Paquete.Idpaquete)
			
			var tipo string
			var valor int
			tipo, valor = tipoYvalor(Paquete.Idpaquete, colaRetail, colaPrioritario, colaNormal ) 

			ent := &Entrega{Id_paquete: Paquete.Idpaquete, 
							Tipo: tipo, 
							Valor: valor, 
							Origen: Paquete.Origen, 
							Destino: Paquete.Destino,
							Intentos: IntIntentos,
							Fecha_entrega: Paquete.Timestamp,
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
