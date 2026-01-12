package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/saumyan/payment_gateway/internals/server"
)



func main(){

	mux := http.NewServeMux()

	ser := server.New(":8080",mux)
//go routine to start a server
	go func(){
		log.Println("server started on :8080")
		if err:= ser.Start();err != nil && err != http.ErrServerClosed{
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop,os.Interrupt,syscall.SIGTERM)

	<-stop
	log.Println("shutdown signal")

	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ser.Close(ctx); err != nil{
		log.Fatalf("gracefull shoutdown: %v" ,err)
	}

	log.Println("Server stopped")



	
}