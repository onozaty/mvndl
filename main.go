package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {

	if len(commit) > 7 {
		commit = commit[:7]
	}
	fmt.Printf("mvndl v%s (%s)\n", version, commit)

	var repo string
	var groupID string
	var artifactID string
	var version string
	var help bool

	flag.StringVar(&repo, "r", "", "repository url")
	flag.StringVar(&groupID, "g", "", "Group ID")
	flag.StringVar(&artifactID, "a", "", "Artifact ID")
	flag.StringVar(&version, "v", "", "version")
	flag.BoolVar(&help, "h", false, "Help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if repo == "" || groupID == "" || artifactID == "" || version == "" {
		flag.Usage()
		os.Exit(1)
	}

}

func createDownloadURLs(repo string, groupID string, artifactID string, version string) ([]string, error) {

	baseURL := path.Join(repo, strings.ReplaceAll(groupID, ".", "/"), artifactID, version)
	baseName := artifactID + "-" + version

	return []string{
		path.Join(baseURL, baseName+".pom"),
		path.Join(baseURL, baseName+".jar"),
		path.Join(baseURL, baseName+"-sources.jar"),
		path.Join(baseURL, baseName+"-javadoc.jar")}, nil
}
