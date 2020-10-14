package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strings"

	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
)

func abrircsvpymes() *Reader {
	csvfile, err := os.Open("pymes.csv")
	if err != nil {
		log.Fatalln("No pude abrir el csv", err)
	}

	// Parsear el csvfile
	r := csv.NewReader(csvfile)

	// Iterar csv

	/*for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("1: %s, 2: %s, 3: %s, 4: %s, 5: %s, 6: %s\n", record[0], record[1], record[2], record[3], record[4], record[5])
	}*/

	return r

}

func preguntasiniciales() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Preguntas iniciales:")
	fmt.Println("Pymes o Retail?")
	fmt.Printf("> ")
	tipotienda, _ := reader.ReadString('\n')
	tipotienda = strings.Replace(tipotienda, "\n", "", -1)
	fmt.Println("Respuesta:", tipotienda)
	fmt.Println("Tiempo en segundos de espera entre el envío de órdenes")
	fmt.Printf("> ")
	tiempoespera, _ := reader.ReadString('\n')
	tiempoespera = strings.Replace(tiempoespera, "\n", "", -1)
	fmt.Println("Respuesta:", tiempoespera)
	return tipotienda, tiempoespera
}

func main() {
	var tipotienda, tiempoespera string
	tipotienda, tiempoespera = preguntasiniciales()

	fmt.Println("tipotienda:", tipotienda)
	fmt.Println("tiempoespera:", tiempoespera)

	//abrircsvpymes()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.178:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	leeropcion := bufio.NewReader(os.Stdin)

	if tipotienda == "pymes" {
		fmt.Printf("Entre a tipotienda == pymes")
		r := abrircsvpymes()

		for true {
			fmt.Println("Desea consultar un código de seguimiento? (si/no)")
			fmt.Printf("> ")
			opcion, _ := leeropcion.ReadString('\n')
			opcion = strings.Replace(opcion, "\n", "", -1)
			if opcion == "si" { // Consulta seguimiento
				fmt.Println("Ingrese código de seguimiento")
				fmt.Printf("> ")
				codigoseguimiento, _ := leeropcion.ReadString('\n')
				codigoseguimiento = strings.Replace(codigoseguimiento, "\n", "", -1)
				/*
					respuesta, err := c.CodigoSeguimiento(context.Background(), codigoseguimiento)
					if err != nil {
						log.Fatalf("Error usando c.CodigoSeguimiento")
					}
					fmt.Println("El estado de su pedido es:", respuesta)
				*/

			}

			/*
				// Esperar el tiempo ya determinado
				// Lanzar orden (preguntar antes si esque aun hay ordenes)
			*/

		}

	} else if tipotienda == "retail" {
		fmt.Printf("Entre a tipotienda == retail")

	}

	message := chat.Ordenclientepymes{
		Id:          "Id",
		Producto:    "Producto",
		Valor:       15,
		Tienda:      "Tienda",
		Destino:     "Destino",
		Prioritario: true,
	}

	response, err := c.RedecirOrdenPymes(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Producto)
}
