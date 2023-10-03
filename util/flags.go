package util

import (
	"encoding/json"
	"flag"
	"io"
	"os"
)

// configurable options set either by flags or json congig file
type Flags struct {
	WHP         *int `json:"wHP"`
	Arrows      *int `json:"arrows"`
	BpChance    *int `json:"bpChance"`
	BatsChance  *int `json:"batsChance"`
	ArrowChance *int `json:"arrowChance"`
	Rows        *int `json:"rows"`
	Cols        *int `json:"cols"`
}

func parseJsonConfig(jsonConfig string) Flags {
	jsonFile, err := os.Open(jsonConfig)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var flags Flags
	err = json.Unmarshal(byteValue, &flags)
	if err != nil {
		panic(err)
	}

	return flags
}

func ParseFlags() Flags {
	jsonConfig := flag.String("jsonConfig", "", "Path to the json config file")
	flag.Parse()

	var flags Flags
	if *jsonConfig != "" {
		flags = parseJsonConfig(*jsonConfig)
	} else {
		flags = Flags{
			flag.Int("wHP", 3, "Health of the Wombat"),
			flag.Int("arrows", 5, "Starting number of arrows"),
			flag.Int("bpChance", 10, "Chance of each room being a bottomless pit"),
			flag.Int("arrowChance", 10, "Chance of each room having an arrow"),
			flag.Int("batsChance", 5, "Chance of each room having bats"),
			flag.Int("rows", 5, "Number of rows"),
			flag.Int("cols", 5, "Number of cols"),
		}
		flag.Parse()
	}

	return flags
}
