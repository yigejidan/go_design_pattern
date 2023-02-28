package _1_interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
解释器模式：给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子。
解释器模式（Interpreter）是一种针对特定问题设计的一种解决方案。例如，匹配字符串的时候，由于匹配条件非常灵活，使得通过代码来实现非常不灵活。举个例子，针对以下的匹配条件：
	以+开头的数字表示的区号和电话号码，如+861012345678；
	以英文开头，后接英文和数字，并以.分隔的域名，如www.liaoxuefeng.com；
	以/开头的文件路径，如/path/to/file.txt；
...
因此，需要一种通用的表示方法——正则表达式来进行匹配。正则表达式就是一个字符串，但要把正则表达式解析为语法树，然后再匹配指定的字符串，就需要一个解释器。
*/

// 假设我们现在有一个监控系统
// 现在需要实现一个告警模块，可以根据输入的告警规则来决定是否触发告警
// 告警规则支持 &&、>、< 3种运算符
// 其中 >、< 优先级比  && 更高

type AlertRule struct {
	expression IExpression
}

func NewAlertRule(rule string) (*AlertRule, error) {
	exp, err := NewAndExpression(rule)
	return &AlertRule{expression: exp}, err
}

func (r *AlertRule) Interpret(stats map[string]float64) bool {
	return r.expression.Interpret(stats)
}

type IExpression interface {
	Interpret(stats map[string]float64) bool
}

type GreaterExpression struct {
	key   string
	value float64
}

func (g *GreaterExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v > g.value
}

func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invaild: %s", exp)
	}
	return &GreaterExpression{
		key:   data[0],
		value: val,
	}, nil
}

type LessExpression struct {
	key   string
	value float64
}

func (g *LessExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v < g.value
}

func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &LessExpression{
		key:   data[0],
		value: val,
	}, nil
}

type AndExpression struct {
	expressions []IExpression
}

func (e AndExpression) Interpret(stats map[string]float64) bool {
	for _, expression := range e.expressions {
		if !expression.Interpret(stats) {
			return false
		}
	}
	return true
}

func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]IExpression, len(exps))

	for i, e := range exps {
		var expression IExpression
		var err error

		switch {
		case strings.Contains(e, ">"):
			expression, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expression, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid: %s", exp)
		}

		if err != nil {
			return nil, err
		}

		expressions[i] = expression
	}

	return &AndExpression{expressions: expressions}, nil
}
