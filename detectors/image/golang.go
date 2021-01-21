package image

import (
	"os/exec"
	"regexp"
)

const (
	GOLANG_1_13 = "1.13"
	GOLANG_1_14 = "1.14"
	GOLANG_1_15 = "1.15"

	DEFAULT_VERSION = GOLANG_1_14
)

type GolangVersionDetector struct {
}

func (detector GolangVersionDetector) Execute() (map[string]string, error) {
	cmd := exec.Command("bash", "-c", "go version")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile("go[[:digit:]].[[:digit:]]*")
	if err != nil {
		return nil, err
	}
	matched := re.Find(output)
	if matched == nil {
		return map[string]string{"version": DEFAULT_VERSION}, nil
	}
	return map[string]string{"version": string(matched)}, nil
}

func (detector GolangVersionDetector) Name() string {
	return "GolangVersionDetector"
}
