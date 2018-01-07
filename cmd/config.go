package cmd

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

type Config map[string]interface{}
type Enum []string

// Schema is store service's configuration schema
var Schema = Config{
	"$schema":              "http://json-schema.org/draft-06/schema#",
	"type":                 "object",
	"additionalProperties": false,
	"required":             Enum{"centrex"},

	"properties": Config{

		"centrex": Config{
			"type":                 "object",
			"additionalProperties": false,
			"required":             Enum{"addr", "secure", "keypair"},

			"properties": Config{
				"addr":   Config{"type": "string"},
				"secure": Config{"type": "boolean"},
				"keypair": Config{
					"type":                 "object",
					"additionalProperties": false,
					"required":             Enum{"cert", "key"},

					"properties": Config{
						"cert": Config{"type": "string"},
						"key":  Config{"type": "string"},
					},
				},
			},
		},
	},
}

// VerifyConfig checks if config is valid.
func VerifyConfig(config Config) error {
	schemaLoader := gojsonschema.NewGoLoader(Schema)
	configLoader := gojsonschema.NewGoLoader(config)

	result, err := gojsonschema.Validate(schemaLoader, configLoader)
	if err != nil {
		return err
	}

	if result.Valid() {
		return nil
	}
	buf := &bytes.Buffer{}
	for i, desc := range result.Errors() {
		fmt.Fprintf(buf, "(%d) %s ", i, desc)
	}
	return errors.Errorf("Invalid config: %s", buf.String())
}
