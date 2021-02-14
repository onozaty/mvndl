package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	flag.StringVar(&repo, "r", "", "Repository (\"jcenter\" or \"central\" or Specify by url)")
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

	err := downloadFiles(repo, groupID, artifactID, version, baseDir)
	if err != nil {
		fmt.Println("\nError: ", err)
		os.Exit(1)
	}
}

func downloadFiles(repo string, groupID string, artifactID string, version string, baseDir string) error {

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

		url, err := joinURL(downloadURLs.baseURL, fileName)
		if err != nil {
			return err
		}

		err = downloadFile(url, filepath.Join(saveDir, fileName))
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(url string, savePath string) error {

	fmt.Printf("%s -> ", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("skipped (%s)\n", resp.Status)
		return nil
	}

	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Print("saved\n")

	return nil
}

func createDownloadURLs(repo string, groupID string, artifactID string, version string) (*DownloadFiles, error) {

	baseURL, err := joinURL(repoURL(repo), strings.ReplaceAll(groupID, ".", "/"), artifactID, version)
	if err != nil {
		return nil, err
	}

	baseName := artifactID + "-" + version

	return &DownloadFiles{
			baseURL: baseURL,
			fileNames: []string{
				baseName + ".pom",
				baseName + ".jar",
				baseName + "-sources.jar",
				baseName + "-javadoc.jar"}},
		nil
}

func joinURL(elem ...string) (string, error) {

	baseURL, err := url.Parse(elem[0])
	if err != nil {
		return "", err
	}

	for _, elm := range elem[1:] {
		baseURL.Path = path.Join(baseURL.Path, elm)
	}

	return baseURL.String(), nil
}

func repoURL(repo string) string {

	switch repo {
	case "jcenter":
		return "https://jcenter.bintray.com/"
	case "central":
		return "https://repo1.maven.org/maven2/"
	default:
		return repo
	}
}
