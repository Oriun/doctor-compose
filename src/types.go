package types

type SupportedDatabase struct {
	Name    string `json:"name"`
	Image   string `json:"image"`
	Storage string `json:"storage"`
	TagUrl  string `json:"tag_url"`
	Port    string `json:"port"`
	Envs    []Env  `json:"envs"`
}

type Env struct {
	Label       string `json:"label"`
	VarName     string `json:"var_name"`
	Default     string `json:"default"`
	Mandatory   bool   `json:"mandatory"`
	Description string `json:"description"`
}

type Service struct {
	Container_name string            `yaml:"container_name,omitempty"`
	Image          string            `yaml:"image,omitempty"`
	Volumes        []string          `yaml:"volumes,omitempty"`
	Restart        string            `yaml:"restart,omitempty"`
	Ports          []string          `yaml:"ports,omitempty"`
	Env            map[string]string `yaml:"env,omitempty"`
	Env_file       []string          `yaml:"env_file,omitempty"`
}

type Compose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}
