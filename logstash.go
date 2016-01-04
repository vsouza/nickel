package main

type Logstash struct {
	Hostname string
	Port     int
}

func NewLogstashResource(config *configFile) *Logstash {
	logs := &Logstash{
		Hostname: config.Logstash.host,
		Port:     config.Logstash.port,
	}
	return logs
}

func ProcessDataBfSend(data []byte) []byte {
	return data
}

func Publish() error {
	var err error

	return err
}
