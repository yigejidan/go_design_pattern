package _5_chain

/*
责任链模式：使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止。
责任链模式（Chain of Responsibility）是一种处理请求的模式，它让多个处理器都有机会处理该请求，直到其中某个处理成功为止。责任链模式把多个处理器串成链，然后让请求在链上传递：
*/

// SensitiveWordFilter 敏感词过滤器
type SensitiveWordFilter interface {
	filter(content string) bool
}

// SensitiveWordFilterChain 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (s *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	s.filters = append(s.filters, filter)
}

func (s *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range s.filters {
		if !filter.filter(content) {
			return filter.filter(content)
		}
	}
	return true
}

type AdSensitiveWordFilter struct {
}

func (a *AdSensitiveWordFilter) filter(content string) bool {
	return false
}

// PoliticalWordFilter 政治敏感
type PoliticalWordFilter struct {
}

func (p *PoliticalWordFilter) filter(content string) bool {
	return false
}
