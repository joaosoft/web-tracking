package validator

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"
)

func (v *Validator) initPwd() {
	var err error
	blackList, err := initPwdBlackList()
	if err != nil {
		v.logger.Info(err)
	}

	v.pwd = &pwd{
		settings: &PwdSettings{
			MinNumeric:     constMinNumeric,
			MinUpper:       constMinUpper,
			MinLower:       constMinLower,
			MinLetter:      constMinLetter,
			MinSpace:       constMinSpace,
			MinSymbol:      constMinSymbol,
			MinPunctuation: constMinPunctuation,
			MinLength:      constMinLength,
			BlackList:      blackList,
		},
	}
}
func initPwdBlackList() (_ map[string]empty, err error) {
	blackList := make(map[string]empty)
	var file *os.File

	file, err = os.Open(constPwdBlackListFile)
	if err != nil {
		return blackList, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		blackList[strings.TrimSuffix(line, "\n")] = empty{}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return blackList, err
	}

	return blackList, nil
}

func (p *PwdSettings) check(value string) {
	for _, ch := range value {
		switch {
		case unicode.IsNumber(ch):
			p.MinNumeric++
		case unicode.IsUpper(ch):
			p.MinUpper++
			p.MinLetter++
		case unicode.IsLower(ch):
			p.MinLower++
			p.MinLetter++
		case ch == ' ':
			p.MinSpace++
		case unicode.IsSymbol(ch):
			p.MinSymbol++
		case unicode.IsPunct(ch):
			p.MinPunctuation++
		}
		p.MinLength++
	}
}

func (p *PwdSettings) Compare(value string) (errs []error) {
	newCheck := &PwdSettings{}

	newCheck.check(value)

	if newCheck.MinNumeric < p.MinNumeric ||
		newCheck.MinUpper < p.MinUpper ||
		newCheck.MinLower < p.MinLower ||
		newCheck.MinLetter < p.MinLetter ||
		newCheck.MinSpace < p.MinSpace ||
		newCheck.MinSymbol < p.MinSymbol ||
		newCheck.MinPunctuation < p.MinPunctuation {
		return []error{ErrorInvalidValue}
	}

	if _, ok := p.BlackList[strings.ToLower(strings.TrimSpace(value))]; ok {
		return []error{ErrorInvalidValue}
	}

	return errs
}
