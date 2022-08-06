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

func useExistingCompose() (types.Compose, bool) {
	currentCompose := types.Compose{}
	currentComposeFile, noCurrent := ioutil.ReadFile("docker-compose.yml")
	if noCurrent != nil {
		fmt.Println("\nNo docker-compose.yml file found, creating one...")
		return currentCompose, false
	}
	currentComposeString := bytes.Replace(currentComposeFile, []byte("#"), []byte("'#description':"), -1)
	err := yaml.Unmarshal(currentComposeString, &currentCompose)
	if err != nil {
		fmt.Println("\nError while parsing docker-compose.yml file, creating one...")
		return currentCompose, false
	} else if currentCompose.Version != "3.9" {
		fmt.Println("\nYour docker-compose.yml file is not compatible with this version of Doctor-Compose, creating one...")
		return currentCompose, false
	}
	return currentCompose, true
}

func mergeCompose(from *types.Compose, into *types.Compose, overwrite bool) {
	for k, v := range from.Services {
		_, ok := into.Services[k]
		if !ok || overwrite {
			into.Services[k] = v
		}
		if ok && !overwrite {
			fmt.Println("\nService " + k + " already exists in your docker-compose.yml file.")
			fmt.Println("\n To overwrite any existing field, use the -f option. Skipping...")
		}
	}
	for k, v := range from.Volumes {
		if len(into.Volumes) == 0 {
			into.Volumes = map[string]interface{}{}
		}
		_, ok := into.Volumes[k]
		if !ok || overwrite {
			into.Volumes[k] = v
		}
		if ok && !overwrite {
			fmt.Println("\n[ERROR] Volume " + k + " already exists in your docker-compose.yml file.")
			fmt.Println("To overwrite any existing field, use the -f option. Skipping...")
		}
	}
	for k, v := range from.Networks {
		if len(into.Networks) == 0 {
			into.Networks = map[string]interface{}{}
		}
		_, ok := into.Networks[k]
		if !ok || overwrite {
			into.Networks[k] = v
		}
		if ok && !overwrite {
			fmt.Println("\n[ERROR] Network " + k + " already exists in your docker-compose.yml file.")
			fmt.Println("To overwrite any existing field, use the -f option. Skipping...")
		}
	}
	for k, v := range from.Configs {
		if len(into.Configs) == 0 {
			into.Configs = map[string]interface{}{}
		}
		_, ok := into.Configs[k]
		if !ok || overwrite {
			into.Configs[k] = v
		}
		if ok && !overwrite {
			fmt.Println("\n[ERROR] Config " + k + " already exists in your docker-compose.yml file.")
			fmt.Println("To overwrite any existing field, use the -f option. Skipping...")
		}
	}
	for k, v := range from.Secrets {
		if len(into.Secrets) == 0 {
			into.Secrets = map[string]interface{}{}
		}
		_, ok := into.Secrets[k]
		if !ok || overwrite {
			into.Secrets[k] = v
		}
		if ok && !overwrite {
			fmt.Println("\n[ERROR] Secret " + k + " already exists in your docker-compose.yml file.")
			fmt.Println("To overwrite any existing field, use the -f option. Skipping...")
		}
	}
}

func main() {

	fmt.Println("\nWelcome to Doctor-Compose, the CLI that diagnose your app and find you the best docker-compose solution.")

	force := true
	currentCompose, _ := useExistingCompose()

	compose, env := WriteCompose()

	mergeCompose(&compose, &currentCompose, force)

	data, err := yaml.Marshal(&currentCompose)

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
