package analyze_test

import (
	"strings"
	"testing"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func newUUID(b []uint8) uuid.UUID {
	return uuid.Must(uuid.FromBytes(b))
}

func TestWithFlake8(t *testing.T) {
	code := `
def test(ξ):
    ξ is None 
    thisfunctiondoesnotexist("Expect an error.")

def too_complex():
    def b():
        def c():
            pass
        c()
    b()
`

	result, err := analyze.WithFlake8(strings.NewReader(code))
	if err != nil {
		t.Error(err)
	}

	errorTests := []model.RoastError{
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID([]uint8{0x23, 0xc2, 0x48, 0x54, 0xd3, 0x6d, 0x59, 0xa1, 0x97, 0x9, 0x6a, 0x66, 0x49, 0xe0, 0x79, 0x3}),
				Row:         0x4,
				Column:      0x5,
				Engine:      "flake8",
				Name:        "F821",
				Description: "undefined name 'thisfunctiondoesnotexist'",
			},
		},
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID([]uint8{0x7f, 0x28, 0xcf, 0x1f, 0x79, 0x2, 0x53, 0x5a, 0xbf, 0xe, 0x13, 0xde, 0x24, 0x66, 0xe7, 0x1b}),
				Row:         0x6,
				Column:      0x1,
				Engine:      "flake8",
				Name:        "E302",
				Description: "expected 2 blank lines, found 1",
			},
		},
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID([]uint8{0x72, 0xd0, 0xc6, 0xd0, 0xe4, 0x73, 0x50, 0xd0, 0x93, 0x2a, 0xb3, 0x75, 0xc9, 0x8e, 0xd4, 0x5c}),
				Row:         0x3,
				Column:      0xe,
				Engine:      "flake8",
				Name:        "W291",
				Description: "trailing whitespace",
			},
		},
	}

	warningTests := []model.RoastWarning{
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID([]uint8{0x72, 0xd0, 0xc6, 0xd0, 0xe4, 0x73, 0x50, 0xd0, 0x93, 0x2a, 0xb3, 0x75, 0xc9, 0x8e, 0xd4, 0x5c}),
				Row:         0x3,
				Column:      0xe,
				Engine:      "flake8",
				Name:        "W291",
				Description: "trailing whitespace",
			},
		},
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID([]uint8{0xcd, 0xb7, 0xea, 0x70, 0xa7, 0x5d, 0x55, 0x8, 0x80, 0x56, 0xb2, 0x66, 0x92, 0x9e, 0xf8, 0x8}),
				Row:         0x6,
				Column:      0x1,
				Engine:      "flake8",
				Name:        "C901",
				Description: "'too_complex' is too complex (3)",
			},
		},
	}

	for i, message := range result.Errors {
		assert.Exactly(t, errorTests[i], message)
	}

	for i, message := range result.Warnings {
		assert.Exactly(t, warningTests[i], message)
	}
}
