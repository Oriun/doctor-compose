package base

import "github.com/savioxavier/termlink"

type PopulateFields map[string]string
type Namable interface {
	getName() string
}
type SupportedDatabase struct {
	Name    string
	Image   string
	Storage string
	TagUrl  string
	Port    string
	Envs    []Env
}

func (data SupportedDatabase) getName() string {
	return data.Name
}

type BoilerPlate struct {
	Name         string
	Url          string
	Tags         []string
	BuildCommand string
	RunCommand   string
	Envs         []Env
	CloneCommand string
}

func (data BoilerPlate) getName() string {
	return data.Name
}
func (data BoilerPlate) GetLink() string {
	return termlink.Link("ctrl+click to see on github", data.Url)
}

type NodeFrameWorkVerions struct {
	BuildCommand string
	RunCommand   string
	Envs         []Env
	BoilerPlate  []BoilerPlate
}

type SupportedNodeFrameworks struct {
	Name    string
	Package string
	Version map[string]NodeFrameWorkVerions
}

func (data SupportedNodeFrameworks) getName() string {
	return data.Name
}

type Env struct {
	Label       string
	VarName     string
	Default     string
	Mandatory   bool
	Static      bool
	Description string
}

type TagResponse struct {
	Results []struct {
		Name string
	}
}

type Service struct {
	Description string            `yaml:"#description,omitempty"`
	Image       string            `yaml:"image,omitempty"`
	Volumes     []string          `yaml:"volumes,omitempty"`
	Restart     string            `yaml:"restart,omitempty"`
	Ports       []string          `yaml:"ports,omitempty"`
	Env         map[string]string `yaml:"environment,omitempty"`
}

type Compose struct {
	Version  string                 `yaml:"version"`
	Services map[string]Service     `yaml:"services"`
	Volumes  map[string]interface{} `yaml:"volumes,omitempty"`
	Configs  map[string]interface{} `yaml:"configs,omitempty"`
	Secrets  map[string]interface{} `yaml:"secrets,omitempty"`
	Networks map[string]interface{} `yaml:"networks,omitempty"`
}
