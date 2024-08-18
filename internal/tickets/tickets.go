package tickets

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ticket struct { // Utilice los nombres de los atributos en funcion del csv 
	Id          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       int
}

func LoadTickets() []Ticket { // Devuelve un array de tickets de todos los tickets del csv
	file, err := os.Open("tickets.csv") // abre el archivo
	if err != nil {
		log.Fatal(err) // vi que se utiliza el log.Fatal en algunos videos de playground, opte por esto en vez de crear mensajes personalizados
	}
	defer file.Close() // por si falla la apertura

	reader := csv.NewReader(file) // abro un reader para empezar a meter los tickets del csv para el array
	var tickets []Ticket

	for {
		record, err := reader.Read()
		if err == io.EOF { // utilizo end of file para cortar la lectura cuando se llega al final
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id, _ := strconv.Atoi(record[0]) // parseo los datos que tienen que ser int
		price, _ := strconv.Atoi(record[5])

		tickets = append(tickets, Ticket{ // voy agregando al array
			Id:          id,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			Time:        record[4],
			Price:       price,
		})
	}
	return tickets
}

func GetTotalTickets(destination string, tickets []Ticket) (int, error) {
	var count int
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			count++
		} // voy matcheando con la destinacion pasada por parametro, si es igual al recorrer los tickets del array voy contando
	}
	return count, nil 
}
func GetCountByPeriod(period string, tickets []Ticket) (int, error) {
	periods := map[string][]int{ // la mejor opcion para mi es utilizar un map, me parecio lo mas organizado, tambien puede ser un switch
		"madrugada": {0, 6},
		"mañana":    {7, 12},
		"tarde":     {13, 19},
		"noche":     {20, 23},
	}

	var count int
	for _, ticket := range tickets {
		timeParts := strings.Split(ticket.Time, ":")
		hour, _ := strconv.Atoi(timeParts[0])

		if hour >= periods[period][0] && hour <= periods[period][1] {
			count++
		} // agarro las horas de los tiempos y las comparo si estan dentro del lapso del periodo
	}
	return count, nil
}
func AverageDestination(destination string, tickets []Ticket) (float64, error) {
	totalTickets, err := GetTotalTickets(destination, tickets) // cuento la cantidad de destinos que hubieron de un destino
	if err != nil {
		return 0, err
	}
	total := len(tickets)
	return (float64(totalTickets) / float64(total)) * 100, nil // calculamos el promedio parseando en float64
}
