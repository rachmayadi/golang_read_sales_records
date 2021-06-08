package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// setup reader
	csvIn, err := os.Open("./sales_records.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	// setup writer
	csvOut, err := os.Create("./sales_records_new.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	// handle header
	rec, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	rec = append(rec, "score")
	if err = w.Write(rec); err != nil {
		log.Fatal(err)
	}

	for {
		rec, err = r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// get float value
		value := rec[1]
		// floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal("Record, error: %v, %v", value, err)

		}
		//change order_date
		new_order_date := rec[5]
		rec[5] = new_order_date
		//change ship date
		new_ship_date := rec[7]
		rec[7] = new_ship_date
		// calculate scores; THIS EXTERNAL METHOD CANNOT BE CHANGED
		score := 1.0

		scoreString := strconv.FormatFloat(score, 'f', 8, 64)
		rec = append(rec, scoreString)

		if err = w.Write(rec); err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}
}
