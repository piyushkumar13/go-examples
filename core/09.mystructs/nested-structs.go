package main

import (
	"encoding/json"
	"fmt"
)

/*
	Following nested structs are helpful in json serialization and deserialization.

But, not when we want to initialize object manually.
*/
type Config struct {
	Server struct {
		Port int
		Host string
	}

	Postgres struct {
		ConnectionString string
		User             string
		Password         string
		Database         string
	}
}

/* Nested structs are helpful in this case which is doing json serialization and deserialization. */
type JsonConfig struct {
	Server struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"server"`

	Postgres struct {
		ConnectionString string `json:"connectionString"`
		User             string `json:"user"`
		Password         string `json:"password"`
		Database         string `json:"database"`
	} `json:"postgres"`
}

func main() {

	/* Since we are manually initializing, nested structs in this way will be more verbose to initialize object as you can see below. */
	configuration := Config{

		Server: struct {
			Port int
			Host string
		}{
			Port: 8080,
			Host: "My server Host",
		},

		Postgres: struct {
			ConnectionString string
			User             string
			Password         string
			Database         string
		}{
			ConnectionString: "my connection string",
			User:             "Piyush",
			Password:         "Pass123",
			Database:         "MyDb",
		},
	}

	fmt.Println("The configuration is ::: ", configuration)
	fmt.Println("The server configuration is ::: ", configuration.Server)
	fmt.Println("The postgres configuration is ::: ", configuration.Postgres)

	/* Here, we will be unmarshalling this below json to JsonConfig struct which is nested struct. */
	sampleJson := []byte(`{
    "server": {
        "port": 8080,
        "host": "My server Host"
    },
    "postgres": {
        "connectionString": "my connection string",
        "user": "Piyush",
        "password": "Pass123",
        "database": "MyDb"
    }
}`)

	var jsonConfig JsonConfig

	err := json.Unmarshal(sampleJson, &jsonConfig)

	if err != nil {
		panic("Unmarshalling error")
	}

	fmt.Println("Unmarshalled object ::: ", jsonConfig)
}
