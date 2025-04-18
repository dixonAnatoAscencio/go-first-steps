package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Simula un trabajador que recibe tareas desde el canal y las procesa
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
        fmt.Printf("👷 Worker %d: procesando tarea %d\n", id, task)
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simula tiempo de trabajo
        fmt.Printf("✅ Worker %d: terminó tarea %d\n", id, task)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    const numWorkers = 3
    const numTasks = 10

    tasks := make(chan int, numTasks) // Canal con búfer
    var wg sync.WaitGroup

    // Lanzamos varios workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, &wg)
    }

    // Enviamos tareas al canal
    for t := 1; t <= numTasks; t++ {
        tasks <- t
        fmt.Printf("📤 Enviada tarea %d al canal\n", t)
    }

    close(tasks) // Cerramos el canal para indicar que no habrá más tareas
    fmt.Println("🚪 Canal cerrado, esperando a que terminen los workers...")

    wg.Wait() // Esperamos a que todos los workers terminen
    fmt.Println("🎉 ¡Todas las tareas fueron procesadas!")
}
