package main

import (
	"fmt"
	"strconv"
	"strings"
)

// BNFParser 结构体
type BNFParser struct {
	tokens            []string
	currentTokenIndex int
}

// NewBNFParser 创建解析器
func NewBNFParser(input string) *BNFParser {
	tokens := tokenize(input)
	return &BNFParser{tokens: tokens}
}

// tokenize 将输入字符串拆分为 token
func tokenize(input string) []string {
	// 使用正则表达式分隔数字和运算符
	tokens := []string{}
	for _, token := range strings.Fields(input) {
		if match := isOperator(token); match {
			tokens = append(tokens, token)
		} else {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

// isOperator 检查是否为运算符
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// Parse 开始解析
func (p *BNFParser) Parse() (int, error) {
	result := p.expression()
	if p.currentTokenIndex < len(p.tokens) {
		return 0, fmt.Errorf("unexpected token: %s", p.tokens[p.currentTokenIndex])
	}
	return result, nil
}

// expression 解析表达式
func (p *BNFParser) expression() int {
	result := p.term()
	for p.currentTokenIndex < len(p.tokens) && (p.tokens[p.currentTokenIndex] == "+" || p.tokens[p.currentTokenIndex] == "-") {
		operator := p.tokens[p.currentTokenIndex]
		p.currentTokenIndex++
		rightTerm := p.term()
		if operator == "+" {
			result += rightTerm
		} else {
			result -= rightTerm
		}
	}
	return result
}

// term 解析项
func (p *BNFParser) term() int {
	result := p.factor()
	for p.currentTokenIndex < len(p.tokens) && (p.tokens[p.currentTokenIndex] == "*" || p.tokens[p.currentTokenIndex] == "/") {
		operator := p.tokens[p.currentTokenIndex]
		p.currentTokenIndex++
		rightFactor := p.factor()
		if operator == "*" {
			result *= rightFactor
		} else {
			result /= rightFactor
		}
	}
	return result
}

// factor 解析因子
func (p *BNFParser) factor() int {
	token := p.tokens[p.currentTokenIndex]
	if isNumber(token) {
		p.currentTokenIndex++
		value, _ := strconv.Atoi(token)
		return value
	} else if token == "(" {
		p.currentTokenIndex++
		result := p.expression()
		if p.currentTokenIndex < len(p.tokens) && p.tokens[p.currentTokenIndex] == ")" {
			p.currentTokenIndex++ // Skip ')'
			return result
		} else {
			panic("missing closing parenthesis")
		}
	}
	panic(fmt.Sprintf("unexpected token: %s", token))
}

// isNumber 检查是否为数字
func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

// 主函数
func main() {
	expression := "3 + 5 * (2 - 8)"
	parser := NewBNFParser(expression)
	result, err := parser.Parse()
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}
}
