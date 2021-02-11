package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	version = "dev"
	commit  = "none"
)

// DownloadFiles ダウンロード対象のファイル情報
type DownloadFiles struct {
	baseURL   string
	fileNames []string
}

func main() {

	if len(commit) > 7 {
		commit = commit[:7]
	}
	fmt.Printf("mvndl v%s (%s)\n", version, commit)

	var repo string
	var groupID string
	var artifactID string
	var version string
	var baseDir string
	var help bool

	flag.StringVar(&repo, "r", "", "Repository URL")
	flag.StringVar(&groupID, "g", "", "Group ID")
	flag.StringVar(&artifactID, "a", "", "Artifact ID")
	flag.StringVar(&version, "v", "", "Version")
	flag.StringVar(&baseDir, "d", "", "Save Directory")
	flag.BoolVar(&help, "h", false, "Help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if repo == "" || groupID == "" || artifactID == "" || version == "" || baseDir == "" {
		flag.Usage()
		os.Exit(1)
	}

}

func download(repo string, groupID string, artifactID string, version string, baseDir string) error {

	saveDir := filepath.Join(baseDir, strings.ReplaceAll(groupID, ".", "/"), artifactID, version)

	err := os.MkdirAll(saveDir, 0777)
	if err != nil {
		return err
	}

	downloadURLs, err := createDownloadURLs(repo, groupID, artifactID, version)

	if err != nil {
		return err
	}

	for _, fileName := range downloadURLs.fileNames {

		resp, err := http.Get(path.Join(downloadURLs.baseURL, fileName))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		out, err := os.Create(filepath.Join(saveDir, fileName))
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func createDownloadURLs(repo string, groupID string, artifactID string, version string) (DownloadFiles, error) {

	baseURL := path.Join(repo, strings.ReplaceAll(groupID, ".", "/"), artifactID, version)
	baseName := artifactID + "-" + version

	return DownloadFiles{
			baseURL: baseURL,
			fileNames: []string{
				baseName + ".pom",
				baseName + ".jar",
				baseName + "-sources.jar",
				baseName + "-javadoc.jar"}},
		nil
}
