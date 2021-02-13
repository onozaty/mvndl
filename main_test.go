package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDownloadFiles(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "mvndl")
	if err != nil {
		t.Fatal("failed test\n", err)
	}
	defer os.RemoveAll(tempDir)

	repo := "https://repo1.maven.org/maven2"
	groupID := "com.github.onozaty"
	artifactID := "postgresql-copy-helper"
	version := "1.1.0"

	err = downloadFiles(repo, groupID, artifactID, version, tempDir)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	fileInfosInDir, err := ioutil.ReadDir(filepath.Join(tempDir, "com/github/onozaty/postgresql-copy-helper/1.1.0"))
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if len(fileInfosInDir) != 4 {
		t.Fatal("failed test\n", len(fileInfosInDir))
	}
}

func TestDownloadFile(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "mvndl")
	if err != nil {
		t.Fatal("failed test\n", err)
	}
	defer os.RemoveAll(tempDir)

	savePath := filepath.Join(tempDir, "downloaded")

	err = downloadFile("https://github.com/onozaty/mvndl/raw/main/README.md", savePath)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	_, err = os.Stat(savePath)
	if os.IsNotExist(err) {
		t.Fatal("failed download\n")
	}
}

func TestDownloadFileNotFound(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "mvndl")
	if err != nil {
		t.Fatal("failed test\n", err)
	}
	defer os.RemoveAll(tempDir)

	savePath := filepath.Join(tempDir, "downloaded")

	err = downloadFile("https://github.com/onozaty/mvndl/raw/main/notfound", savePath)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	_, err = os.Stat(savePath)
	if os.IsExist(err) {
		t.Fatal("failed test\n")
	}
}

func TestCreateDownloadURLs(t *testing.T) {

	repo := "http://example.com/rep"
	groupID := "com.github.onozaty"
	artifactID := "mvndl"
	version := "1.0.0"

	expect := &DownloadFiles{
		baseURL: "http://example.com/rep/com/github/onozaty/mvndl/1.0.0",
		fileNames: []string{
			"mvndl-1.0.0.pom",
			"mvndl-1.0.0.jar",
			"mvndl-1.0.0-sources.jar",
			"mvndl-1.0.0-javadoc.jar"}}

	downloadFiles, err := createDownloadURLs(repo, groupID, artifactID, version)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if !reflect.DeepEqual(downloadFiles, expect) {
		t.Fatal("failed test\n", downloadFiles)
	}
}
