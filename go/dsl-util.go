package mydsl

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func LoadYaml(filename interface{}) (interface{}, error) {
	if strFilename, ok := filename.(string); ok {
		if strings.HasSuffix(strFilename, ".yml") {
			// router.yml読み取り
			f, err := os.Open(strFilename)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			yamlInput, err := ioutil.ReadAll(f)
			if err != nil {
				return nil, err
			}
			var parsedYaml interface{}
			yamlError := yaml.UnmarshalStrict(yamlInput, &parsedYaml)
			if yamlError != nil {
				return nil, yamlError
			} else {
				return parsedYaml, nil
			}
		} else {
			return nil, errors.New(fmt.Sprintf("include: 1st argument must be .yml file. %v", strFilename))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("include: 1st argument must be string. %v", filename))
	}
}
