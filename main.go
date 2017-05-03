package main

// go run main.go -f hacked_point or just execute go run main.go
/*
1. Read input from file
2. Decode data
3. Base convertion
4. Generate CSV as output: result.csv on same directory
*/

import (
	"encoding/csv"
	"flag"
	"github.com/Rakanixu/epic/player/data"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f := flag.String("f", "hacked_points", "Path to codified puntuations for Epic Invaders XXIII")

	flag.Parse()

	if *f == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	s, err := os.Open(*f)
	if err != nil {
		log.Fatal(err)
	}

	playersData, err := proccessInput(s)
	if err != nil {
		log.Fatal(err)
	}

	// 3 steeps, calculate the base,
	//decode the input to numeric and finally convert from X base to 10 base
	for _, v := range playersData {
		v.CalcBase()
		v.Decode()
		v.ToBase10()
	}

	// Sort the results
	sort.Sort(data.SortByPoints(playersData))

	// Generates a CSV file from playersData
	generateCSV(playersData)
}

// proccessInput reads a file, removes break lines and convert its content to a slice of players data
func proccessInput(f *os.File) ([]*data.Data, error) {
	// Read file
	b := make([]byte, 1000)
	count, err := f.Read(b)
	if err != nil {
		return nil, err
	}

	// Remove break lines from input
	records := strings.Split(string(b[:count]), "\n")

	var playersData []*data.Data

	// Generate players Data
	for _, v := range records {
		record := strings.Split(v, ",")

		p := &data.Data{
			PlayerName:  record[0],
			CodedBase:   record[1],
			CodedPoints: record[2],
		}

		playersData = append(playersData, p)
	}

	return playersData, nil
}

func generateCSV(data []*data.Data) {
	// Create file
	f, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// New CSV Writter
	w := csv.NewWriter(f)
	defer w.Flush()

	// Write into the file
	for _, v := range data {
		err := w.Write([]string{
			v.PlayerName,
			strconv.Itoa(int(v.Points)),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
