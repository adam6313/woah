package yaml

import (
	"woah/pkg/encoder"

	"gopkg.in/yaml.v2"
)

// yamlEncoder -
type yamlEncoder struct {
	encoderType string
}

// Encoder -
func (y *yamlEncoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Decode -
func (y *yamlEncoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

// Type -
func (y yamlEncoder) Type() string {
	return y.encoderType
}

// NewEncoder -
func NewEncoder() encoder.Encoder {
	return &yamlEncoder{
		encoderType: "yaml",
	}
}
