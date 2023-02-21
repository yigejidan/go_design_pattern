package _2_factory

type IRule interface {
	getRule()
}

type JsonRule struct {
}

func (j JsonRule) getRule() {}

type YamlRule struct {
}

func (y YamlRule) getRule() {}

func NewIRule(t string) IRule {
	switch t {
	case "json":
		return JsonRule{}
	case "yaml":
		return YamlRule{}
	default:
		return nil
	}
}
