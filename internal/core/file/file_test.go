package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		fileName  string
		url       string
		fileType  string
		expectErr bool
	}{
		{
			name:      "valid png file",
			fileName:  "image.png",
			url:       "https://cdn.example.com/files/image.png",
			fileType:  "image/png",
			expectErr: false,
		},
		{
			name:      "invalid file type",
			fileName:  "file.zip",
			url:       "https://cdn.example.com/files/file.zip",
			fileType:  "application/zip",
			expectErr: true,
		},
		{
			name:      "empty file type",
			fileName:  "file.png",
			url:       "https://cdn.example.com/files/file.png",
			fileType:  "",
			expectErr: true,
		},
		{
			name:      "missing file name",
			fileName:  "",
			url:       "https://cdn.example.com/files/image.png",
			fileType:  "image/png",
			expectErr: true,
		},
		{
			name:      "missing url",
			fileName:  "image.png",
			url:       "",
			fileType:  "image/png",
			expectErr: true,
		},
		{
			name:      "file name too long",
			fileName:  string(make([]byte, 101)),
			url:       "https://cdn.example.com/files/long.jpg",
			fileType:  "image/png",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := New(tt.fileName, tt.url, tt.fileType)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, file)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, file)
				assert.Equal(t, tt.fileName, file.Name)
				assert.Equal(t, tt.url, file.URL)
				assert.Equal(t, tt.fileType, file.Type.String())
			}
		})
	}
}
