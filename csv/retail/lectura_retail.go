package main

import(
	"os"
	"log"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type retail struct {
	id string
	producto string
	valor int
	tienda string
	destino string
}

func main(){

	f, err := os.Open("retail.csv")
	if err != nil{
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.FieldsPerRecord = 5

	if _, err := r.Read(); err != nil{
		panic(err)
	}


	var retails []retail
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		p := retail{
			id: record[0],
			producto: record[1],
			tienda: record[3],
			destino: record[4],
		}

		if record[2] != ""{
			i, err := strconv.Atoi(record[2])
			if err != nil{
				log.Printf("error al procesar el valor: %v", err)
				continue
			}
			p.valor = i
		}

		retails = append(retails, p)
	}
	fmt.Println(retails)
}
