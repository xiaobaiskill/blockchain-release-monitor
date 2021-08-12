package github_release

import (
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

func GetCurrentVersion(url string) (version string, err error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to load url(%s): %s", url, err.Error()))
	}
	node, err := htmlquery.Query(doc, "//title")
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get title(%s): %s", url, err.Error()))
	}

	version = parseReleaseTitle(htmlquery.InnerText(node))
	if version == "" {
		return "", errors.New(fmt.Sprintf("No release information is obtained from %s", url))
	}

	return version, nil
}

func parseReleaseTitle(info string) (version string) {
	data := strings.Split(info, "·")
	if len(data) < 2 {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(data[0], "Release "))
}
