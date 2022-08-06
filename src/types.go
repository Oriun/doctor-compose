package types

type SupportedDatabase struct {
	Name    string `json:"name"`
	Image   string `json:"image"`
	Storage string `json:"storage"`
	TagUrl  string `json:"tag_url"`
	Port    string `json:"port"`
	Envs    []Env  `json:"envs"`
}

type SupportedNodeFrameworks struct {
	Name         string `json:"name"`
	Package      string `json:"package"`
	Version      string `json:"version"`
	Envs         []Env  `json:"envs"`
	BuildCommand string `json:"build_command"`
	RunCommand   string `json:"run_command"`
}

type Env struct {
	Label       string `json:"label"`
	VarName     string `json:"var_name"`
	Default     string `json:"default"`
	Mandatory   bool   `json:"mandatory"`
	Description string `json:"description"`
}

type TagResponse struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
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
