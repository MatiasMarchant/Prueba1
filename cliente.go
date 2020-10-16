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
	tiempoesperaint, _ := strconv.Atoi(strings.TrimSuffix(tiempoespera, "\r\n"))
	if err != nil {
		log.Fatalln("No pude abrir el csv:", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return true
		}
		valorint32, _ := strconv.ParseInt(row[2], 10, 32)

		orden := chat.Ordenclientepymes{
			Id:       row[0],
			Producto: row[1],
			Valor:    int32(valorint32),
		}
	}
}

func ingresarordenespymes(nombreexcel string, tiempoespera string, c chat.ChatServiceClient bool) {
	csvfile, err := os.Open(nombreexcel)
	tiempoesperaint, _ := strconv.Atoi(strings.TrimSuffix(tiempoespera, "\r\n")) // DANGER
	if err != nil {
		log.Fatalln("No pude abrir el csv:", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	for{
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return true
		}
		valorint32, _ := strconv.ParseInt(row[2], 10, 32)

		if row[5] == 1 {
			prioritarioBool = true
		} else {
			prioritarioBool = false
		}
		orden := chat.Ordenclientepymes{
			Id: row[0],
			Producto: row[1],
			Valor: row[2],
			Tienda: row[3],
			Destino: row[4],
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
	tiempoesperaint, _ := strconv.Atoi(strings.TrimSuffix(tiempoespera, "\r\n")) // CUIDADO \r POR WINDOWS
	if err != nil {
		log.Fatalln("No pude abrir el csv:", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
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

		response, err := c.RedecirOrdenRetail(context.Background(), &orden)
		if err != nil {
			log.Fatalf("Error usando RedecirOrdenRetail: %s", err)
		}

		log.Printf("Respuesta de RedecirOrdenRetail: %s", response.Producto)
		time.Sleep(time.Second * time.Duration(int64(tiempoesperaint)))
	}
}

func preguntasiniciales() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Preguntas iniciales:")
	fmt.Println("Pymes o Retail?")
	fmt.Printf("> ")
	tipotienda, _ := reader.ReadBytes('\n')
	//tipotienda = strings.Replace(tipotienda, "\n", "", -1)
	fmt.Println("Respuesta:", string(tipotienda))
	fmt.Println("Tiempo en segundos de espera entre el envío de órdenes")
	fmt.Printf("> ")
	tiempoespera, _ := reader.ReadBytes('\n')
	//tiempoespera = strings.Replace(tiempoespera, "\n", "", -1)
	fmt.Println("Respuesta:", string(tiempoespera))
	return string(tipotienda), string(tiempoespera)
}

func main() {
	var tipotienda, tiempoespera string
	tipotienda, tiempoespera = preguntasiniciales()

	fmt.Println("tipotienda:", tipotienda)
	fmt.Println("tiempoespera:", tiempoespera)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	if string(tipotienda) == "pymes\r\n" { // CUIDADO HAY Q CAMBIAR A SOLO \n CUANDO LO PASEMOS A LAS MAQUINAS VIRTUALES (CREOOOOOOOOOOOOOOOOOOOOOOO) XD EL \r ES PQ ESTOY EN WINDOWS, FUCK WINDOWS
		fmt.Println("Entre a tipotienda == pymes")
		go ingresarordenespymes("pymes.csv", tiempoespera, c)
	} else if string(tipotienda) == "retail\r\n" {
		fmt.Println("Entre a tipotienda == retail")
		go ingresarordenesretail("retail.csv", tiempoespera, c)
	}

	leeropcion := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Ingrese codigo de seguimiento, si quiere salir ingrese exit")
		fmt.Printf("> ")
		opcion, _ := leeropcion.ReadString('\n')
		opcion = strings.Replace(opcion, "\n", "", -1)
		if opcion == "exit\r\n" {
			break
		} /* else { // Codigo de seguimiento
			codigoseguimiento, _ := leeropcion.ReadString('\n')
			codigoseguimiento = strings.TrimSuffix(codigoseguimiento, "\r\n")
			respuesta, err := c.CodigoSeguimiento(context.Background(), codigoseguimiento)
			if err != nil {
				log.Fatalf("Error usando c.CodigoSeguimiento")
			}
			fmt.Println("El estado de su pedido es:", respuesta)
		}*/
	}

	/*
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
	*/
}
