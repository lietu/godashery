package godashery

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)


type Settings struct {
	ListenAddress string		`yaml:"listen_address"`
	ListenPort    int           `yaml:"listen_port"`
	WwwPath       string        `yaml:"www_path"`
	UpdatesPerSec int	        `yaml:"updates_per_sec"`
}


func (s *Settings) ToYaml() []byte {
	result, err := yaml.Marshal(s)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return result
}

func GetSettings() Settings {
	settings := Settings{
		"0.0.0.0",
		8080,
		"../../www/",
		3,
	}

	data, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	yaml.Unmarshal(data, &settings)

	return settings
}
