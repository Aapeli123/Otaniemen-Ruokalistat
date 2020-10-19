package main

import (
	"bufio"
	"fmt"
	"os"

	"otaniemenruokalistat.tk/telegram"
	"otaniemenruokalistat.tk/webapi"
)

func main() {
	fmt.Println("Initializing...")
	go telegram.Init() // Kommentoi tämä rivi pois sammuttaaksesi telegrambotin
	go webapi.Init()   // Sama mutta webserver
	fmt.Println("PRESS ENTER TO EXIT")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
