package filetype

import (
	"errors"
	"strings"
)

type FileType string

const (
	PNG  FileType = "image/png"
	JPEG FileType = "image/jpeg"
	JPG  FileType = "image/jpg"
)

var validTypes = map[FileType]bool{
	PNG:  true,
	JPEG: true,
	JPG:  true,
}

func New(value string) (FileType, error) {
	ft := FileType(strings.ToLower(strings.TrimSpace(value)))
	if !validTypes[ft] {
		return "", errors.New("invalid file type")
	}
	return ft, nil
}

func (f FileType) String() string {
	return string(f)
}
