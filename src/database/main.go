package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	types "oriun/doctor-compose/src"
	"os"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func ReadDatabaseJSON() []types.SupportedDatabase {
	// Open our jsonFile
	jsonFile, err := os.Open("src/database/data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var dbs []types.SupportedDatabase

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'dbs' which we defined above
	json.Unmarshal(byteValue, &dbs)

	return dbs

}

func getNames(vs []types.SupportedDatabase) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = v.Name
	}
	return vsm
}

func GetService() (string, types.Service) {
	var data = ReadDatabaseJSON()
	service := types.Service{}

	answers := struct {
		Type        string
		PersistMode string
		Volume      string
	}{}

	err := survey.Ask([]*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What type of database would you create ?",
				Options: getNames(data),
				Default: "MongoDB",
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	err2 := survey.Ask([]*survey.Question{
		{
			Name: "persistMode",
			Prompt: &survey.Select{
				Message: "Do you want to persist the database on disk or on a volume?",
				Options: []string{"Disk", "Volume", "None"},
				Default: "Disk",
			},
		},
	}, &answers)
	if err2 != nil {
		panic((err2))
	}

	var currentDB types.SupportedDatabase

	for i := range data {
		if data[i].Name == answers.Type {
			currentDB = data[i]
			break
		}
	}

	if answers.PersistMode == "Disk" {
		err3 := survey.Ask([]*survey.Question{
			{
				Name: "volume",
				Prompt: &survey.Input{
					Message: "Where to persist ?",
					Default: "./" + strings.ToLower(answers.Type) + ".db",
				},
			},
		}, &answers)
		if err3 != nil {
			panic((err3))
		} else {
			valid, err4 := regexp.MatchString(`^[a-z:\.]{0,3}\/.+`, answers.Volume)
			if err4 != nil {
				panic((err4))
			}
			if !valid {
				answers.Volume = "./" + answers.Volume
			}
		}
	} else if answers.PersistMode == "Volume" {
		err5 := survey.Ask([]*survey.Question{
			{
				Name: "volume",
				Prompt: &survey.Input{
					Message: "Where to persist ?",
					Default: strings.ToLower(answers.Type) + ".db",
				},
			},
		}, &answers)
		if err5 != nil {
			panic((err5))
		} else {
			if strings.Contains(answers.Volume, "/") {
				split := strings.Split(answers.Volume, "/")
				last := split[len(split)-1]
				ok, err6 := regexp.MatchString("[a-zA-Z0-9]+", last)
				if err6 != nil {
					panic(err6)
				} else if ok {
					fmt.Printf("Invalid volume name, %q will be used instead", last)
					answers.Volume = last
				} else {
					panic("Invalid volume name")
				}
			}
		}
	}

	service.Container_name = "doctor-" + strings.ToLower(answers.Type)
	service.Image = currentDB.Image
	service.Volumes = []string{
		answers.Volume + ":" + currentDB.Storage,
	}

	return "tolo", service
}
