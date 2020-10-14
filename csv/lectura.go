package main

import(
	"os"
	"log"
	"encoding/csv"
	"fmt"
	"io"
)


func main(){
	f, err := os.Open("pymes.csv")
	if err != nil{
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.FieldsPerRecord = 6
	
	rawData, err := r.ReadAll()
	if err != nil{
		log.Printf("error al leer la infoMMM %v", err)
	}
	
	fmt.Println(rawData)
}
