package main

import (
	"flag"
	"log"
)

func main() {
	// init resources
	configFile := flag.String("c", "nickel.conf", "configuration setup file")
	interval := flag.Int("i", 3600, "Interval for process billing file")
	flag.Parse()
	var err error
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	s3Resource := NewS3Resource(config)
	csvResource := NewCSVResource(config)
	logstashResource := NewLogstashResource(config)
}
