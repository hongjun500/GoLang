// @author hongjun500
// @date 2023/9/15 14:51
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package factory

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		name string
		want RuleConfigParser
	}{
		{
			name: "json",
			want: &RuleJsonConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
