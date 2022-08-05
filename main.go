package main

import (
	"bytes"
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

func WriteCompose() (types.Compose, string) {

	answers := struct {
		Name string
		Type string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		panic(err)
	}

	var services = map[string]types.Service{}
	var envs string
	if answers.Type == "Database" {
		name, service, env := database.GetService()
		services[name] = service
		envs = env + envs
	}
	return types.Compose{
		Version:  "3.9",
		Services: services,
	}, envs
}

func main() {

	fmt.Println("\nWelcome to Doctor-Compose, the CLI that diagnose your app and find you the best docker-compose solution.")

	compose, env := WriteCompose()

	data, err := yaml.Marshal(&compose)

	if err != nil {
		panic(err)
	}

	/* allow comments  from the field '#description'*/
	data = bytes.Replace(data, []byte("'#description':"), []byte("#"), -1)

	err = ioutil.WriteFile("docker-compose.yml", data, 0777)

	if err != nil {
		panic(err)
	}

	if len(env) > 0 {
		err = ioutil.WriteFile(".env", []byte(env), 0777)

		if err != nil {
			panic(err)
		}
	}

}
