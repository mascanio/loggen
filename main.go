package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"
)

type log struct {
	Level     string
	Timestamp time.Time
	Msg       string
	Host      string
}

func main() {
	var format = flag.String("format", "json", "valid formats: json, str")
	flag.Parse()
	last := time.Now()
	switch *format {
	case "json":
		for {
			log := log{
				Level:     "DEBUG",
				Timestamp: time.Now(),
				Msg:       time.Since(last).String(),
				Host:      "ASDF",
			}
			last = time.Now()
			data, err := json.Marshal(log)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", data)
			time.Sleep(50 * time.Millisecond)
		}
	case "str":
		for {
			fmt.Printf("[%v] [%v] %v\n", time.Now().Format(time.RFC3339), "DEBUG", time.Since(last).String())
			last = time.Now()
			time.Sleep(50 * time.Millisecond)
		}
	}
}
