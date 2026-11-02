package main
import (
	config "github.com/Hanniah-Heart/Blog-Aggre-Gator/internal/config"
	"fmt"
	"log"
)

func main () {
	Config, err := config.Read(); if err != nil {log.Fatal(err)}
	fmt.Printf("Run complete. Result is: %v\n", Config)
}
