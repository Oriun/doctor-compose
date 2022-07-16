package main

import (
	"fmt"
	"io/ioutil"
	types "oriun/doctor-compose/src"
	"oriun/doctor-compose/src/database"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"
)

var qs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is the name of your project ?",
			Default: "My Project",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "What type of project are you creating ?",
			Options: []string{"Database", "Other"},
			Default: "Database",
		},
	},
}

type Service interface{}

func WriteCompose() types.Compose {

	database.ReadDatabaseJSON()

	answers := struct {
		Name string
		Type string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		panic(err)
	}

	var services = map[string]types.Service{}

	if answers.Type == "Database" {
		name, service := database.GetService()
		services[name] = service
	}
	return types.Compose{
		Version:  "3.9",
		Services: services,
	}
}

func main() {

	fmt.Println("\nWelcome to Doctor-Compose, the CLI that diagnose your app and find you the best docker-compose solution.")

	compose := WriteCompose()

	data, err := yaml.Marshal(&compose)

	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile("docker-compose.yml", data, 0777)

	if err2 != nil {
		panic(err2)
	}

}
