// @author hongjun500
// @date 2023/10/18 10:56
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 职责链模式

package main

// SensitiveWordFilter 敏感词过滤器
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// AddFilter 添加过滤器
func (s *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	s.filters = append(s.filters, filter)
}

// Filter 过滤
func (s *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range s.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter 广告敏感词过滤器
type AdSensitiveWordFilter struct {
}

func (a *AdSensitiveWordFilter) Filter(content string) bool {
	return false
}

// PoliticalWordFilter 政治敏感词过滤器
type PoliticalWordFilter struct {
}

func (p *PoliticalWordFilter) Filter(content string) bool {
	return true
}
