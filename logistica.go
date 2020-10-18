package main

import (
	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
	"log"
	"net"
	"time"
	//"chat"

	"google.golang.org/grpc"
	"fmt"
	//"github.com/streadway/amqp"
	//"encoding/json"
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

func IntInArr(a int, arr []int) bool {
    for _, i := range arr {
        if i == a {
            return true
        }
    }
    return false
}

func procesarEntregas(paquetesProcesados []string, paqueteEnMarcha []chat.PaqueteEnMarcha) ([]string , []Entrega) {
	
	var entregasProcesadas []Entrega

	for _, Paquete := range paqueteEnMarcha {

		IntIntentos, _ := strconv.Atoi(Paquete.Intentos)
		if ((Paquete.Estado == "Recibido" || Paquete.Estado == "No Recibido" ) && !IntInArr( Paquete.Id_paquete, paquetesProcesados)) {
			
			paquetesProcesados = append(paquetesProcesados, IntIdpaquete)

			ent := &Entrega{Id_paquete: Paquete.Id_paquete, 
							Tipo: "Malo, cambiar", 
							Valor: 0, 
							Origen: Paquete.Origen, 
							Destino: Paquete.Destino,
							Intentos: IntIntentos,
							Fecha_entrega: Paquete.Timestamp}		

			entregasProcesadas = append(entregasProcesadas, *ent)    		
		}
	}	
	return paquetesProcesados, entregasProcesadas
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
	var entregasProcesadas []Entrega

	go func() {
		for {
			time.Sleep(2 * time.Second)
			
			paquetesProcesados, entregasProcesadas = procesarEntregas(paquetesProcesados, s.PaquetesEnMarcha)

			fmt.Println(s.PaquetesEnMarcha)
			fmt.Println(paquetesProcesados)
			fmt.Println(entregasProcesadas)
			fmt.Println("____________")

			//procesarEntregas(paquetesProcesados, s.PaquetesEnMarcha)
						


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
