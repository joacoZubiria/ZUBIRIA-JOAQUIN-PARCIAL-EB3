package main

import (
    "fmt"
    "github.com/joacoZubiria/ZUBIRIA-JOAQUIN-PARCIAL-EB3/desafio-go-bases/internal/tickets"
)

func main() {
    ticketsData := tickets.LoadTickets() // Toma los tickets del csv

    // creo los canales para las funciones gorutines
    totales := make(chan int) 
    mananas := make(chan int)
    porcentaje := make(chan float64)

    go func() {
        total, _ := tickets.GetTotalTickets("Brazil", ticketsData)
        totales <- total
    }()

    go func() {
        count, _ := tickets.GetCountByPeriod("mañana", ticketsData)
        mananas <- count
    }()

    go func() {
        percentage, _ := tickets.AverageDestination("Brazil", ticketsData)
        porcentaje <- percentage
    }()

    total := <-totales
    mornings := <-mananas
    percentage := <-porcentaje

    fmt.Printf("Tickets totales a Brazil: %d\n", total)
    fmt.Printf("Tickets en la manana: %d\n", mornings)
    fmt.Printf("Porcentaje que va a Brazil: %.2f%%\n", percentage)
}
