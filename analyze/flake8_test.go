package analyze_test

import (
	"io"
	"os/exec"
	"testing"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

const code = `
def test(ξ):
    thisfunctiondoesnotexist("Expect an error.") 

def too_complex():
    def b(z, x, c, v, b, n, p, o, i, u, y, t):
        if z == x:
            pass
        if x == c:
            pass
        if c == v:
            pass
        if b == n:
            pass
        if n == p:
            pass
        if p == o:
            pass
        if o == i:
            pass
        if i == u:
            pass
        if u == y:
            pass
        if y == t:
            pass
        if z == t:
            pass

    b(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
`

func newUUID(s string) uuid.UUID {
	return uuid.Must(uuid.FromString(s))
}

func resetExecCommand() {
	analyze.ExecCommand = exec.Command
}

func TestWithFlake8CmdStdoutPipeError(t *testing.T) {
	analyze.ExecCommand = func(command string, args ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		cmd.Stdout = &io.PipeWriter{}
		return cmd
	}
	defer resetExecCommand()

	_, err := analyze.WithFlake8("bot", code)
	assert.EqualError(t, err, "exec: Stdout already set")
}

func TestWithFlake8CmdStart(t *testing.T) {
	analyze.ExecCommand = func(_ string, args ...string) *exec.Cmd {
		return exec.Command("", args...) // non-existing command
	}
	defer resetExecCommand()

	_, err := analyze.WithFlake8("bot", code)
	assert.EqualError(t, err, "fork/exec : no such file or directory")
}

func TestWithFlake8CmdStdoutInvalidFormat(t *testing.T) {
	analyze.ExecCommand = func(command string, args ...string) *exec.Cmd {
		args[2] = "--format=none"
		return exec.Command(command, args...)
	}
	defer resetExecCommand()

	_, err := analyze.WithFlake8("bot", code)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")
}

func TestWithFlake8CmdErrorExitStatus(t *testing.T) {
	analyze.ExecCommand = func(command string, args ...string) *exec.Cmd {
		args = append(args[:5], args[6:]...) // remove `--exit-zero`
		return exec.Command(command, args...)
	}
	defer resetExecCommand()

	_, err := analyze.WithFlake8("bot", code)
	assert.EqualError(t, err, "exit status 1")
}

func TestWithFlake8HitCache(t *testing.T) {
	first_result, _ := analyze.WithFlake8("bot", code)
	second_result, _ := analyze.WithFlake8("bot", code)

	assert.Equal(t, first_result, second_result)
}

func TestWithFlake8EmptyFile(t *testing.T) {
	result, err := analyze.WithFlake8("", "")

	assert.Equal(t, nil, err)
	assert.Equal(t, "", result.Username)
	assert.Equal(t, uint(0), result.Score)
	assert.Equal(t, "python3", result.Language)
}

func TestWithFlake8Simple(t *testing.T) {
	result, err := analyze.WithFlake8("bot", code)
	if err != nil {
		t.Error(err)
	}

	errorTests := []model.RoastError{
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID("23c24854-d36d-59a1-9709-6a6649e07903"),
				Row:         3,
				Column:      5,
				Engine:      "flake8",
				Name:        "F821",
				Description: "undefined name 'thisfunctiondoesnotexist'",
			},
		},
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID("7f28cf1f-7902-535a-bf0e-13de2466e71b"),
				Row:         5,
				Column:      1,
				Engine:      "flake8",
				Name:        "E302",
				Description: "expected 2 blank lines, found 1",
			},
		},
	}

	warningTests := []model.RoastWarning{
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID("72d0c6d0-e473-50d0-932a-b375c98ed45c"),
				Row:         3,
				Column:      49,
				Engine:      "flake8",
				Name:        "W291",
				Description: "trailing whitespace",
			},
		},
		{
			RoastMessage: model.RoastMessage{
				Hash:        newUUID("cdb7ea70-a75d-5508-8056-b266929ef808"),
				Row:         5,
				Column:      1,
				Engine:      "flake8",
				Name:        "C901",
				Description: "'too_complex' is too complex (13)",
			},
		},
	}

	for i, expected := range errorTests {
		assert.Exactly(t, expected, result.Errors[i])
	}

	for i, expected := range warningTests {
		assert.Exactly(t, expected, result.Warnings[i])
	}

	assert.Equal(t, "bot", result.Username)
	assert.Equal(t, uint(10), result.Score)
}

func BenchmarkWithFlake8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := analyze.WithFlake8("bot", code)
		if err != nil {
			b.Error(err)
		}
	}
}
