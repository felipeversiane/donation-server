package document

import (
	"errors"
	"regexp"
)

type Document struct {
	value string
}

func New(value string) (Document, error) {
	if value == "" {
		return Document{}, errors.New("document is required")
	}
	if isValidCPF(value) || isValidCNPJ(value) {
		return Document{value: value}, nil
	}
	return Document{}, errors.New("invalid document")
}

func (d Document) String() string {
	return d.value
}

func isValidCPF(cpf string) bool {
	re := regexp.MustCompile(`^\d{11}$`)
	return re.MatchString(cpf)
}

func isValidCNPJ(cnpj string) bool {
	re := regexp.MustCompile(`^\d{14}$`)
	return re.MatchString(cnpj)
}
