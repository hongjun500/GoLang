// @author hongjun500
// @date 2023/9/15 14:35
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 工厂方法模式

package factory

// IRuleConfigParserFactory 工厂接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type JsonRuleConfigParserFactory struct {
}

func (j JsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &JsonRuleConfigParser{}
}

type YamlRuleConfigParserFactory struct {
}

func (y YamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &YamlRuleConfigParser{}
}

// NewIRuleConfigParserFactory 工厂方法
func NewIRuleConfigParserFactory(s string) IRuleConfigParserFactory {
	switch s {
	case "json":
		return &JsonRuleConfigParserFactory{}
	case "yaml":
		return &YamlRuleConfigParserFactory{}
	}
	return nil
}
