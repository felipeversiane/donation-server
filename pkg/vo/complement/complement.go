package complement

import "strings"

type Complement struct {
	value string
}

func New(value string) (Complement, error) {
	return Complement{value: strings.TrimSpace(value)}, nil
}

func (c Complement) String() string {
	return c.value
}
