package main

import (
    "fmt"
    "sync"
    "time"
)

func saludar(nombre string, wg *sync.WaitGroup) {
    defer wg.Done() // Indica que esta goroutine terminó
    for i := 0; i < 5; i++ {
        fmt.Println("Hola", nombre)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1) // Esperamos una goroutine

    go saludar("Goroutine", &wg)

    saludar("Main", &wg)
    wg.Wait() // Esperamos que todas las goroutines terminen
}
