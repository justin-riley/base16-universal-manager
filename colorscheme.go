package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

type B16Colors struct {
	Name   string `yaml:"scheme"`
	Author string `yaml:"author"`

	Color00 string `yaml:"base00"`
	Color01 string `yaml:"base01"`
	Color02 string `yaml:"base02"`
	Color03 string `yaml:"base03"`
	Color04 string `yaml:"base04"`
	Color05 string `yaml:"base05"`
	Color06 string `yaml:"base06"`
	Color07 string `yaml:"base07"`
	Color08 string `yaml:"base08"`
	Color09 string `yaml:"base09"`
	Color10 string `yaml:"base0A"`
	Color11 string `yaml:"base0B"`
	Color12 string `yaml:"base0C"`
	Color13 string `yaml:"base0D"`
	Color14 string `yaml:"base0E"`
	Color15 string `yaml:"base0F"`
}

func (c *B16Colors) getColors(url string) *B16Colors {

	yamlFile, err := DownloadYAML(url)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func GetB16Colorscheme(name string) (*B16Colors, error) {
	return nil
}