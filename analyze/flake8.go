package analyze

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/satori/go.uuid"
)

const (
	engineName   = "flake8"
	languageName = "python3"
)

var command = commandPath("./analyze/python3/analyze.py")
var domainUUID = uuid.Must(
	uuid.FromString("badefdd6-8997-425d-9d9e-ae31a01daf0c"))

type Flake8Message struct {
	Code         string `json:"code"`
	Filename     string `json:"filename"`
	LineNumber   uint   `json:"line_number"`
	ColumnNumber uint   `json:"column_number"`
	Text         string `json:"text"`
	PhysicalLine string `json:"physical_line"`
}

type Flake8Result map[string][]Flake8Message

func (f Flake8Result) toRoast() (roast model.RoastResult) {
	result := f["stdin"]

	for _, message := range result {
		id := uuid.NewV5(domainUUID, message.Code)

		switch []rune(message.Code)[0] {
		case 'W', 'C':
			roast.AddWarning(id.Bytes(),
				message.LineNumber,
				message.ColumnNumber,
				engineName,
				message.Code,
				message.Text)
		case 'F', 'I', 'R', 'E':
			roast.AddError(id.Bytes(),
				message.LineNumber,
				message.ColumnNumber,
				engineName,
				message.Code,
				message.Text)
		}
	}

	return
}

func executablePath() string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(exe)
}

func commandPath(p string) string {
	return path.Join(executablePath(), p)
}

func WithFlake8(code io.Reader) (result model.RoastResult, err error) {
	cmd := exec.Command(command)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		_, err = io.Copy(stdin, code)
	}()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return model.RoastResult{}, err
	}

	if err := cmd.Start(); err != nil {
		return model.RoastResult{}, err
	}

	if err := json.NewDecoder(stdout).Decode(&result); err != nil {
		return model.RoastResult{}, err
	}

	if err := cmd.Wait(); err != nil {
		return model.RoastResult{}, err
	}

	return
}
