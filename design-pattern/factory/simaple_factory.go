// @author hongjun500
// @date 2023/9/14 15:39
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 简单工厂

package factory

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type IRuleConfigParser interface {
	Parser(data []byte)
}

type JsonRuleConfigParser struct {
	JsonData
}
type JsonData struct {
	Json  string `json:"json"`
	Count int    `json:"count"`
}

func (j *JsonRuleConfigParser) Parser(data []byte) {
	if data != nil {
		json.Unmarshal(data, j)
	}
}

type YamlRuleConfigParser struct {
	Yaml  string   `yaml:"yaml"`
	Tag   []string `yaml:"tag"`
	Count int      `yaml:"count"`
}

func (y *YamlRuleConfigParser) Parser(data []byte) {
	if data != nil && len(data) > 0 {
		yaml.Unmarshal(data, y)
	}
}

// NewIRuleConfigParser 简单工厂 返回接口
func NewIRuleConfigParser(s string) IRuleConfigParser {
	switch s {
	case "json":
		return &JsonRuleConfigParser{}
	case "yaml":
		return &YamlRuleConfigParser{}
	}
	return nil
}
