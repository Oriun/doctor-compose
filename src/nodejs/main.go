package nodejs

import (
	"bufio"
	"fmt"
	base "oriun/doctor-compose/src"
	nodejs_data "oriun/doctor-compose/src/nodejs/data"
	"os"
	"os/exec"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	semver "github.com/hashicorp/go-version"
)

type Answers struct {
	Framework       string
	Cwd             string
	Name            string
	Selected        base.SupportedNodeFrameworks
	Version         string
	BoilerplateName string
	Boilerplate     base.BoilerPlate
}

func automaticConfiguration(answers *Answers) {

}

func mannualConfiguration(answers *Answers) {
	names := base.GetNames(nodejs_data.Data)
	var selectedFramework base.SupportedNodeFrameworks
	var version string

	err := survey.Ask([]*survey.Question{
		{
			Name: "framework",
			Prompt: &survey.Select{
				Message: "What type of framework do you want to use ?",
				Options: names,
				Default: names[0],
			},
		},
	}, answers)
	if err != nil {
		panic((err))
	}

	for i := range nodejs_data.Data {
		if nodejs_data.Data[i].Name == answers.Framework {
			selectedFramework = nodejs_data.Data[i]
			break
		}
	}
	answers.Selected = selectedFramework

	for key := range selectedFramework.Version {
		if version == "" {
			version = key
		} else {
			current, err := semver.NewVersion(version)
			if err != nil {
				panic(err)
			}
			other, err := semver.NewVersion(key)
			if err != nil {
				panic(err)
			}
			if current.LessThan(other) {
				version = key
			}
		}

	}
	answers.Version = version

	boilerplates, ok := selectedFramework.Version[version]

	if !ok {
		panic("No boilerplate found for this version")
	}

	err = survey.Ask([]*survey.Question{
		{
			Name: "boilerplateName",
			Prompt: &survey.Select{
				Message: "What boilerplate do you want to use ?",
				Options: base.GetNames(boilerplates.BoilerPlate),
				Default: base.GetNames(boilerplates.BoilerPlate)[0],
				Description: func(value string, index int) string {
					return boilerplates.BoilerPlate[index].GetLink()
				},
			},
		},
	}, answers)

	for i := range boilerplates.BoilerPlate {
		if boilerplates.BoilerPlate[i].Name == answers.BoilerplateName {
			answers.Boilerplate = boilerplates.BoilerPlate[i]
			break
		}
	}

	clean1 := regexp.MustCompile(`[^a-zA-Z0-9]`)
	clean2 := regexp.MustCompile(`(^[^a-zA-Z]|[^a-zA-Z]$)`)
	answers.Name = clean2.ReplaceAllString(clean1.ReplaceAllString(answers.Cwd, ""), "-")

	fmt.Printf("Creating %s app...", answers.Name)
	shellcmd := base.Populate(answers.Boilerplate.CloneCommand, base.PopulateFields{"APP_NAME": answers.Name})
	fmt.Printf("Running \n%s\n", shellcmd)
	cmd := exec.Command("/bin/sh", "-c", shellcmd)
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Print(m + " ")
	}
	err = cmd.Wait()

	if err != nil {
		panic(err)
	}

	fmt.Println("\nSuccess!")

}

func GetService() (string, base.Service, string) {
	service := base.Service{}
	answers := Answers{}
	var name = ""
	var env_string = ""

	err := survey.Ask([]*survey.Question{
		{
			Name: "cwd",
			Prompt: &survey.Input{
				Message: "What is the app name (or directory path if already exists) ?",
				Default: "doctor-node",
			},
		},
	}, &answers)

	// Read current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Look if answers.Cwd is a directory
	if _, err := os.Stat(cwd + "/" + answers.Cwd); os.IsNotExist(err) {
		mannualConfiguration(&answers)
	} else {
		fmt.Println("Found app directory, skipping creation")
		// Read package.json in answers.Cwd directory
		automaticConfiguration(&answers)
	}

	/*
	 * Do things here
	 */

	return name, service, env_string
}
