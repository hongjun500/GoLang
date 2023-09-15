// @author hongjun500
// @date 2023/9/14 16:04
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIRuleConfigParser(t *testing.T) {
	var jdata any = `
		{
		"json":"json_data",
		"count":18
		}
	`
	jsonParser := NewIRuleConfigParser("json")
	jsonParser.Parser([]byte(jdata.(string)))

	assert.IsType(t, jsonParser, &JsonRuleConfigParser{})
	assert.NotNil(t, jsonParser.(*JsonRuleConfigParser))
	assert.Equal(t, jsonParser.(*JsonRuleConfigParser).Json, "json_data")

	var ydata any = `
    yaml: yaml_data
    count: 2
    tag:
     - tag1
     - tag2
    `
	yamlParser := NewIRuleConfigParser("yaml")
	yamlParser.Parser([]byte(ydata.(string)))
	assert.IsType(t, yamlParser, &YamlRuleConfigParser{})
	assert.NotNil(t, yamlParser.(*YamlRuleConfigParser))
	assert.Equal(t, yamlParser.(*YamlRuleConfigParser).Yaml, "yaml_data")
}
