package main

import(
	"os"
	"log"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type pymes struct {
	id string
	producto string
	valor int
	tienda string
	destino string
	propietario int
}

func main(){
	f, err := os.Open("pymes.csv")
	if err != nil{
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.FieldsPerRecord = 6

	var pyme []pymes
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		p := pymes{
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

		if record[5] != ""{
			j, err := strconv.Atoi(record[5])
			if err != nil{
				log.Printf("error al procesar el propietario: %v", err)
				continue
			}
			p.propietario = j
		}

		pyme = append(pyme, p)
	}
	fmt.Println(pyme)
}
