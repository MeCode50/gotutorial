package main 

import (
	"log"
	"github.com/sikozonpc/ecom/cmd/api"
)
// setting up and running the server 
func main () {
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
if err := server.Run(); err != nil {
	log.Fatal(err)

}

}