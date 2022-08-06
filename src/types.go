package types

type SupportedDatabase struct {
	Name    string
	Image   string
	Storage string
	TagUrl  string
	Port    string
	Envs    []Env
}

type BoilerPlate struct {
	Name         string
	Url          string
	Tags         []string
	BuildCommand string
	RunCommand   string
	Envs         []Env
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
