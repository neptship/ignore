package version

import (
	"encoding/json"
	"errors"
	"net/http"
)

var Version string

func Latest() (version string, err error) {
	if err != nil {
		return "", err
	}

	resp, err := http.Get("https://api.github.com/repos/neptunsk1y/ignore/releases/latest")
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return
	}

	if release.TagName == "" {
		err = errors.New("empty tag name")
		return
	}
	Version = release.TagName[1:]
	return
}
