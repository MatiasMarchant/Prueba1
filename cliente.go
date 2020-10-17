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
	"strconv"
	"strings"
	"time"

	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
)

func ingresarordenespymes(nombreexcel string, tiempoespera string, c chat.ChatServiceClient) bool {
	csvfile, err := os.Open(nombreexcel)
	tiempoesperaint, _ := strconv.Atoi(strings.TrimSuffix(tiempoespera, "\n")) // DANGER
	if err != nil {
		log.Fatalln("No pude abrir el csv:", err)
	}
	defer csvfile.Close()
	var prioritarioBool bool
	r := csv.NewReader(csvfile)
	r.Read() // Saltarse la gracias primera linea
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return true
		}
		valorint32, _ := strconv.ParseInt(row[2], 10, 32)

		if row[5] == "1" {
			prioritarioBool = true
		} else {
			prioritarioBool = false
		}
		orden := chat.Ordenclientepymes{
			Id:          row[0],
			Producto:    row[1],
			Valor:       int32(valorint32),
			Tienda:      row[3],
			Destino:     row[4],
			Prioritario: prioritarioBool,
		}

		response, err := c.RecibirOrdenPymes(context.Background(), &orden)
		if err != nil {
			log.Fatalf("Error usando RecibirOrdenPymes: %s", err)
		}

		log.Printf("Codigo de seguimiento: %s", response.Nordenseguimiento)
		time.Sleep(time.Second * time.Duration(int64(tiempoesperaint)))
	}

}

func ingresarordenesretail(nombreexcel string, tiempoespera string, c chat.ChatServiceClient) bool {
	csvfile, err := os.Open(nombreexcel)
	tiempoesperaint, _ := strconv.Atoi(strings.TrimSuffix(tiempoespera, "\n")) // CUIDADO \r POR WINDOWS
	if err != nil {
		log.Fatalln("No pude abrir el csv:", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	r.Read() // Saltarse la gracias primera linea
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return true
		}
		valorint32, _ := strconv.ParseInt(row[2], 10, 32)

		orden := chat.Ordenclienteretail{
			Id:       row[0],
			Producto: row[1],
			Valor:    int32(valorint32),
			Tienda:   row[3],
			Destino:  row[4],
		}

		c.RecibirOrdenRetail(context.Background(), &orden)

		time.Sleep(time.Second * time.Duration(int64(tiempoesperaint)))
	}
}

func preguntasiniciales() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Preguntas iniciales:")
	fmt.Println("Pymes o Retail?")
	fmt.Printf("> ")
	tipotienda, _ := reader.ReadBytes('\n')
	//fmt.Println("Respuesta:", string(tipotienda))
	fmt.Println("Tiempo en segundos de espera entre el envío de órdenes")
	fmt.Printf("> ")
	tiempoespera, _ := reader.ReadBytes('\n')
	//fmt.Println("Respuesta:", string(tiempoespera))
	return string(tipotienda), string(tiempoespera)
}

func main() {
	var tipotienda, tiempoespera string
	tipotienda, tiempoespera = preguntasiniciales()

	//fmt.Println("tipotienda:", tipotienda)
	//fmt.Println("tiempoespera:", tiempoespera)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist37:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	if string(tipotienda) == "pymes\n" { // CUIDADO HAY Q CAMBIAR A SOLO \n CUANDO LO PASEMOS A LAS MAQUINAS VIRTUALES (CREOOOOOOOOOOOOOOOOOOOOOOO) XD EL \r ES PQ ESTOY EN WINDOWS, FUCK WINDOWS
		fmt.Println("Entre a tipotienda == pymes")
		go ingresarordenespymes("pymes.csv", tiempoespera, c)
	} else if string(tipotienda) == "retail\n" {
		fmt.Println("Entre a tipotienda == retail")
		ingresarordenesretail("retail.csv", tiempoespera, c)
	}

	leeropcion := bufio.NewReader(os.Stdin)
	if string(tipotienda) == "pymes\n" {
		for true {
			fmt.Println("Ingrese codigo de seguimiento, si quiere salir ingrese exit")
			fmt.Printf("> ")
			opcion, _ := leeropcion.ReadString('\n')
			opcion = strings.Replace(opcion, "\n", "", -1)
			if opcion == "exit\n" {
				break
			} else { // Codigo de seguimiento
				codigoseguimiento, _ := leeropcion.ReadString('\n')
				codigoseguimiento = strings.TrimSuffix(codigoseguimiento, "\n")
				orden := chat.Ordenseguimiento{
					Nordenseguimiento: codigoseguimiento,
				}
				respuesta, err := c.CodigoSeguimiento(context.Background(), &orden)
				if err != nil {
					log.Fatalf("Error usando c.CodigoSeguimiento")
				}
				fmt.Println("El estado de su pedido es:", respuesta.Estado)
			}
		}
	}
}
