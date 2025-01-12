package langs_test

import (
	"fmt"
	"testing"

	"github.com/sheldonhull/version-bump/langs"

	changelog "github.com/anton-yurchenko/go-changelog"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	dockerRegex := []string{
		fmt.Sprintf("^LABEL .*org.opencontainers.image.version['\"= ]*[vV]?(?P<version>%v)['\"]?.*", changelog.SemVerRegex),
		fmt.Sprintf("^\\s*['\"]?org.opencontainers.image.version['\"= ]*[vV]?(?P<version>%v)['\"]?.*", changelog.SemVerRegex),
	}

	golangRegex := []string{
		fmt.Sprintf("^const [vV]ersion\\s*string = \"[vV]?(?P<version>%v)\"", changelog.SemVerRegex),
		fmt.Sprintf("^const [vV]ersion := \"[vV]?(?P<version>%v)\"", changelog.SemVerRegex),
		fmt.Sprintf("^\\s*[vV]ersion\\s*string = \"[vV]?(?P<version>%v)\"", changelog.SemVerRegex),
	}

	javaScriptJSONFields := []string{
		"version",
	}

	type test struct {
		Name           string
		ExpectedResult *langs.Language
	}

	suite := map[string]test{
		"Docker": {
			Name: "Docker",
			ExpectedResult: &langs.Language{
				Name:  "Docker",
				Files: []string{"Dockerfile"},
				Regex: &dockerRegex,
			},
		},
		"Go": {
			Name: "Go",
			ExpectedResult: &langs.Language{
				Name:  "Go",
				Files: []string{"*.go"},
				Regex: &golangRegex,
			},
		},
		"JavaScript": {
			Name: "JavaScript",
			ExpectedResult: &langs.Language{
				Name: "JavaScript",
				Files: []string{
					"package.json",
					"package-lock.json",
				},
				JSONFields: &javaScriptJSONFields,
			},
		},
		"Not Supported Language": {
			Name:           "not-supported-language",
			ExpectedResult: nil,
		},
	}

	var counter int
	for name, test := range suite {
		counter++
		t.Logf("Test Case %v/%v - %s", counter, len(suite), name)

		r := langs.New(test.Name)

		if test.Name == "not-supported-language" {
			a.Equal(test.ExpectedResult, r)
		} else {
			a.EqualValues(test.ExpectedResult, r)
		}
	}
}
