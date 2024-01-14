package main

import (
	"bufio"
	"fmt"
	"kurnosovmak/search/src/libs/Engine"
	"os"
	"strconv"
	"strings"
)

func main() {
	engine := Engine.NewEngine()

	fmt.Println("Engine starting ...")
	engine.Start()
	fmt.Println("Engine start")

	fmt.Println("Write search text ...")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		results := engine.Handle(Engine.HandleParams{
			Word: text,
		})

		for _, element := range *results.Items {
			fmt.Println("<  id - " + strconv.Itoa(element))
		}

	}
}
