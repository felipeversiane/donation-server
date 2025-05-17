package documenttype

import (
	"errors"
	"strings"
)

type DocumentType string

const (
	CPF  DocumentType = "cpf"
	CNPJ DocumentType = "cnpj"
)

var validTypes = map[DocumentType]bool{
	CPF:  true,
	CNPJ: true,
}

func New(value string) (DocumentType, error) {
	t := DocumentType(strings.ToLower(strings.TrimSpace(value)))
	if !validTypes[t] {
		return "", errors.New("invalid document type: must be 'cpf' or 'cnpj'")
	}
	return t, nil
}

func (d DocumentType) String() string {
	return string(d)
}
