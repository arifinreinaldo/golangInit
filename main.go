package main

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
	"strings"
)

var global = "Cibai"

func main() {
	var quantity uint
	const total uint = 50
	var remaining = total
	//var array = []string{}
	//var nameArray []string
	var data = strconv.FormatUint(uint64(total), 10)
	fmt.Println(data)
	var name = "Reinaldo Arifin"
	var names = strings.Fields(name)
	var mapData = make(map[string]string)       //map
	var mapsData = make([]map[string]string, 0) //list of map
	mapData["Gurara"] = "Guririr"
	mapsData = append(mapsData, mapData)
	if mapData["Guririr"] == "" {
		fmt.Println("Ga ketemu")
	}
	fmt.Println(mapData["Gurara2"])
	fmt.Println(names[0])
	util.Cabecabean()
	fmt.Println(util.GlobalTiket)
	for {
		fmt.Println("Welcome Ticket Manager")
		fmt.Println("Remaining :", total)
		fmt.Println("Want to Buy ", remaining, "?")
		fmt.Scan(&quantity)
		remaining = remaining - quantity
		fmt.Println("Remainding ", remaining)
		//array = append(array, "remaining")
		//fmt.Println(len(array))
		//fmt.Println(len(nameArray))
		firstName := []string{}
		fmt.Println("First Name ", len(firstName))
		//validateName(name)
		fmt.Println(validateName(name))
		satu, dua, tiga := multipleReturn(name)
		fmt.Println(satu, dua, tiga)
		if remaining == 50 {
			break
		} else if remaining == 10 {
			break
		} else {
			continue
		}
	}

}

func multipleReturn(name string) (string, string, string) { //multiple return value
	return name, "", "Hohoh"
}

func validateName(name string) string {
	return "Okee " + name
}

type Orang struct { //kaya data class di kotlin
	firstName   string
	lastName    string
	qty         uint
	cibabababai map[string]string
}
