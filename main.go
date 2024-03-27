package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()
        router.GET("/hello-world", getHelloWorld)
        log.Fatal(router.Run(":8080"))
}

func getHelloWorld(c *gin.Context) {
        fmt.Println("received /hello-world request\n")
        c.String(http.StatusOK,"Hello, World!")
}
