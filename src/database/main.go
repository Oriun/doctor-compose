package database

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	types "oriun/doctor-compose/src"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

func getNames(vs []types.SupportedDatabase) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = v.Name
	}
	return vsm
}

func fetchTags(url string, wg *sync.WaitGroup, result chan []string) {
	defer wg.Done()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		result <- []string{"lts"}
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		result <- []string{"lts"}
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		result <- []string{"lts"}
		return
	}

	var tags types.TagResponse
	err = json.Unmarshal([]byte(resBody), &tags)
	if err != nil {
		fmt.Printf("client: could not parse response body: %s\n", err)
		result <- []string{"lts"}
		return
	}

	tagList := []string{}
	var recommendedTag string = "lts"
	canBeRecommended, err := regexp.Compile(`^\d+(\.\d+(\.\d+)?)$`)
	if err != nil {
		panic(err)
	}
	for i := range tags.Results {
		name := tags.Results[i].Name
		if name == "lts" {
			continue
		}
		if canBeRecommended.MatchString(name) && recommendedTag == "lts" {
			recommendedTag = name
		}
		tagList = append(tagList, name)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(tagList)))
	result <- append([]string{recommendedTag}, tagList...)
}

func randomString(length int) string {
	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = chars[v%byte(len(chars))]
	}
	return string(bytes)
}

func populate(str string) string {
	var s = str
	s = strings.Replace(s, "${RANDOM_STRING}", randomString(16), -1)
	return s
}

func GetService() (string, types.Service, string) {
	/* Final service data */
	service := types.Service{}

	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))
	fmt.Println(randomString(16))

	/* Answers object */
	answers := struct {
		Type         string
		PersistMode  string
		Volume       string
		Expose       string
		Name         string
		Restart      string
		Tag          string
		Env_location string
		Env_set      string
		Envs         map[string]string
	}{}

	/* 1. Choose a type of database */
	err := survey.Ask([]*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What type of database would you create ?",
				Options: getNames(Data),
				Default: "MongoDB",
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	/* 1.1 Get chosen database config */
	var currentDB types.SupportedDatabase
	for i := range Data {
		if Data[i].Name == answers.Type {
			currentDB = Data[i]
			break
		}
	}

	/* 1.2 Fetch image tags asynchronously */
	var tagFetchRoutine sync.WaitGroup

	tagFetchRoutine.Add(1)
	tagChannel := make(chan []string, 100)
	go fetchTags(currentDB.TagUrl, &tagFetchRoutine, tagChannel)

	/* 2. Choose the type of persistance, then volume or directory name */
	err = survey.Ask([]*survey.Question{
		{
			Name: "persistMode",
			Prompt: &survey.Select{
				Message: "Do you want to persist the database on disk or on a volume?",
				Options: []string{"Disk", "Volume", "None"},
				Default: "Disk",
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	if answers.PersistMode == "Disk" {
		err := survey.Ask([]*survey.Question{
			{
				Name: "volume",
				Prompt: &survey.Input{
					Message: "Where to persist ?",
					Default: "./" + strings.ToLower(answers.Type) + ".db",
				},
			},
		}, &answers)
		if err != nil {
			panic((err))
		} else {
			valid, err2 := regexp.MatchString(`^[a-z:\.]{0,3}\/.+`, answers.Volume)
			if err2 != nil {
				panic((err2))
			}
			if !valid {
				answers.Volume = "./" + answers.Volume
			}
		}
	} else if answers.PersistMode == "Volume" {
		err = survey.Ask([]*survey.Question{
			{
				Name: "volume",
				Prompt: &survey.Input{
					Message: "Where to persist ?",
					Default: strings.ToLower(answers.Type) + ".db",
				},
			},
		}, &answers)
		if err != nil {
			panic((err))
		} else {
			if strings.Contains(answers.Volume, "/") {
				split := strings.Split(answers.Volume, "/")
				last := split[len(split)-1]
				ok, err2 := regexp.MatchString("[a-zA-Z0-9]+", last)
				if err2 != nil {
					panic(err2)
				} else if ok {
					fmt.Printf("Invalid volume name, %q will be used instead", last)
					answers.Volume = last
				} else {
					panic("Invalid volume name")
				}
			}
		}
	}

	/* 3. Configure ports */
	var defaultPort = currentDB.Port
	err = survey.Ask([]*survey.Question{
		{
			Name: "expose",
			Prompt: &survey.Input{
				Message: "What external port do you want to access it on ? (0 to not expose it)",
				Default: defaultPort,
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	/* 4. Set restart policy */
	err = survey.Ask([]*survey.Question{
		{
			Name: "restart",
			Prompt: &survey.Select{
				Message: "Choose a restart policy",
				Options: []string{"always", "on-failure", "unless-stopped", "never"},
				Default: "unless-stopped",
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}
	/* 5. Choose image tag */
	tagFetchRoutine.Wait() // Wait for the fetch routine to finish
	tags := <-tagChannel
	err = survey.Ask([]*survey.Question{
		{
			Name: "tag",
			Prompt: &survey.Select{
				Message: "Choose an image tag",
				Options: tags,
				Default: tags[0],
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	/* 6. Choose service name */
	err = survey.Ask([]*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the name of the service ?",
				Default: "doctor-" + strings.ToLower(answers.Type),
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	/* 7. Set envs from `.env` or `docker-compose.yml` */
	err = survey.Ask([]*survey.Question{
		{
			Name: "env_location",
			Prompt: &survey.Select{
				Message: "Set environment variables from .env or compose file` ?",
				Options: []string{".env", "docker-compose"},
				Default: ".env",
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	/* 8. Choose to set all environment variable or only mandatories */
	err = survey.Ask([]*survey.Question{
		{
			Name: "env_set",
			Prompt: &survey.Select{
				Message: "Set all environment variables in this tool ?",
				Options: []string{"All", "Only mandatories"},
				Default: "Only mandatories",
			},
		},
	}, &answers)
	if err != nil {
		panic((err))
	}

	/* 9. Set environment variables */
	answers.Envs = make(map[string]string)
	for _, env := range currentDB.Envs {
		if !env.Mandatory && answers.Env_set == "Only mandatories" {
			continue
		}
		tmp := struct {
			Value string
		}{}
		err = survey.Ask([]*survey.Question{
			{
				Name: "value",
				Prompt: &survey.Input{
					Message: env.Label + " (default to randomly generated)",
					Help:    env.Description,
					Default: populate(env.Default),
				},
			},
		}, &tmp)
		if err != nil {
			panic((err))
		}
		answers.Envs[env.VarName] = tmp.Value
	}
	if len(answers.Envs) == 0 {
		fmt.Println("No environment variables to set")
	}

	service.Description = answers.Type + " database service"
	service.Image = currentDB.Image + ":" + answers.Tag
	service.Restart = answers.Restart
	if answers.Expose != "" && answers.Expose != "0" {
		service.Ports = []string{answers.Expose + ":" + currentDB.Port}
	}
	service.Volumes = []string{answers.Volume + ":" + currentDB.Storage}

	var env_string string
	if answers.Env_location == ".env" {
		safeName := strings.Replace(answers.Name, "-", "_", -1)
		service.Env = map[string]string{}
		for k, v := range answers.Envs {
			env_string += safeName + "_" + k + "=" + v + "\n"
			service.Env[k] = "${" + safeName + "_" + k + "}"
		}
	} else {
		service.Env = answers.Envs
	}

	fmt.Println("Service created")

	return answers.Name, service, env_string
}
