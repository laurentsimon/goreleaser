package nametemplate

import (
	"testing"

	"github.com/goreleaser/goreleaser/config"
	"github.com/goreleaser/goreleaser/context"
	"github.com/goreleaser/goreleaser/internal/artifact"
	"github.com/tj/assert"
)

func TestNameTemplate(t *testing.T) {
	var ctx = context.New(config.Project{
		ProjectName: "proj",
		Archive: config.Archive{
			NameTemplate: "{{.Binary}}_{{.ProjectName}}_{{ .Os }}_{{ .Arch }}{{ if .Arm }}v{{ .Arm }}{{ end }}",
			Replacements: map[string]string{
				"windows": "windao",
			},
		},
	})
	s, err := Apply(ctx, artifact.Artifact{
		Goos:   "windows",
		Goarch: "amd64",
		Name:   "winbin",
	}, "bin")
	assert.NoError(t, err)
	assert.Equal(t, "bin_bin_windao_amd64", s)
	s, err = Apply(ctx, artifact.Artifact{
		Goos:   "darwin",
		Goarch: "amd64",
		Name:   "winbin",
	}, "bin")
	assert.NoError(t, err)
	assert.Equal(t, "bin_bin_darwin_amd64", s)
}

func TestInvalidNameTemplate(t *testing.T) {
	var ctx = context.New(config.Project{
		ProjectName: "proj",
		Archive: config.Archive{
			NameTemplate: "{{.Binary}",
		},
	})
	s, err := Apply(ctx, artifact.Artifact{
		Goos:   "windows",
		Goarch: "amd64",
		Name:   "winbin",
	}, "bin")
	assert.EqualError(t, err, `template: archive_name:1: unexpected "}" in operand`)
	assert.Empty(t, s)
}
