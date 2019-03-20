package mydsl

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestForCoverage(t *testing.T) {
	f, err := os.Open("test/yamls/testsuite.yml")
	if err != nil {
		t.Fatalf("open error:%v", err)
	}

	defer f.Close()
	yamlInput, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("read error:%v", err)
	}

	var testsuites []map[interface{}]interface{}
	yamlError := yaml.UnmarshalStrict(yamlInput, &testsuites)
	if yamlError != nil {
		t.Fatalf("unmarshal error:%v", err)
	}

	container := &map[string]interface{}{}
	for _, testsuite := range testsuites {
		evaluated, err := NewArgument(testsuite).Evaluate(container)
		if err != nil {
			t.Errorf("testsuite %s failed: %v", testsuite["testsuite"].([]interface{})[0], evaluated)
		}
	}
}
func TestForTestsuite(t *testing.T) {
	f, err := os.Open("test/yamls/testsuite_test.yml")
	if err != nil {
		t.Fatalf("open error:%v", err)
	}

	defer f.Close()
	yamlInput, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("read error:%v", err)
	}

	var testsuites []map[interface{}]interface{}
	yamlError := yaml.UnmarshalStrict(yamlInput, &testsuites)
	if yamlError != nil {
		t.Fatalf("unmarshal error:%v", err)
	}

	container := &map[string]interface{}{}
	for _, testsuite := range testsuites {
		evaluated, err := NewArgument(testsuite).Evaluate(container)
		if err == nil {
			t.Errorf("testsuite %s failed: %v", testsuite["testsuite"].([]interface{})[0], evaluated)
		}
	}
}
