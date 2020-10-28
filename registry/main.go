package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type RegistryConfig struct {
	Registry struct {
		Discovery struct {
			Modules  string `yaml:"modules"`
			Login    string `yaml:"login"`
			Provider string `yaml:"provider"`
		} `yaml:"discovery"`
	} `yaml:"registry"`
}

func main() {
	c := flag.String("c", "telm.yaml", "the configuration file to use")
	flag.Parse()

	ac, err := filepath.Abs(*c)
	if err != nil {
		log.Fatalf("could not find the configuration file: %s, %s\n", *c, err)
	}

	yc, err := ioutil.ReadFile(ac)
	if err != nil {
		log.Fatalf("could not read configuration file: %s, %s\n", ac, err)
	}

	var conf RegistryConfig
	err = yaml.Unmarshal(yc, &conf)
	if err != nil {
		log.Fatalf("error while parsing configuration file: %s, %s\n", ac, err)
	}

	fmt.Printf("config: %v\n", conf)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"registry": "ok",
		})
	})

	r.GET("/.well-known/terraform.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"modules.v1": conf.Registry.Discovery.Modules,
		})
	})

	log.Fatal(r.Run())
}
