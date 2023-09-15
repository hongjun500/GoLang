// @author hongjun500
// @date 2023/9/15 14:51
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 抽象工厂模式

package factory

type RuleConfigParser interface {
	Parse(data []byte)
}

type RuleJsonConfigParser struct {
}

func (r *RuleJsonConfigParser) Parse(data []byte) {
	panic("implement me")
}

type SystemConfigParser interface {
	ParserSystem(data []byte)
}

type SystemJsonConfigParser struct {
}

func (s *SystemJsonConfigParser) ParserSystem(data []byte) {
	panic("implement me")
}

type ConfigParserFactory interface {
	CreateRuleParser() RuleConfigParser
	CreateSystemParser() SystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() RuleConfigParser {
	return &RuleJsonConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() SystemConfigParser {
	return &SystemJsonConfigParser{}
}
