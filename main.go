package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	var urls []string
	var finalGeneratedBackups []string
	flag.Parse()
	if filename := flag.Arg(0); filename != "" {
		urls = StdinOrFile(filename)
	} else {
		urls = StdinOrFile("stdin")
	}
	backupWords := StdinOrFile("./backup.txt")

	url_backup_items := url_backup(urls, backupWords)
	url_url_backup_items := url_url_backup(urls, backupWords)
	url_dot_path_backup_item := url_dot_path_backup(urls, backupWords)
	finalGeneratedBackups = append(finalGeneratedBackups, url_backup_items...)
	finalGeneratedBackups = append(finalGeneratedBackups, url_url_backup_items...)
	finalGeneratedBackups = append(finalGeneratedBackups, url_dot_path_backup_item...)

	for _, v := range RemoveDuplicateStr(finalGeneratedBackups) {
		fmt.Println(v)

	}

}

func url_backup(urls []string, backupWords []string) []string {
	var generatedBackups []string

	for _, item := range urls {
		u, err := url.Parse(item)
		if err != nil {
			panic(err)
		}

		if len(u.Path) == 0 || u.Path == "/" || len(u.Hostname()) == 0 {
			continue
		}

		var backupItem string
		for _, word := range backupWords {

			if len(u.Scheme) != 0 {
				backupItem = u.Scheme + "://" + u.Host + u.Path + word
			} else {
				backupItem = u.Host + u.Path + word
			}

			generatedBackups = append(generatedBackups, backupItem)
		}
	}
	return generatedBackups
}
func url_url_backup(urls []string, backupWords []string) []string {
	var generatedBackups []string

	for _, item := range urls {
		u, err := url.Parse(item)
		if err != nil {
			panic(err)
		}

		if len(u.Path) == 0 || u.Path == "/" || len(u.Hostname()) == 0 {
			continue
		}

		var backupItem string
		for _, word := range backupWords {

			if len(u.Scheme) != 0 {
				backupItem = u.Scheme + "://" + u.Host + "/" + u.Host + word
			} else {
				backupItem = u.Host + "/" + u.Host + word
			}

			generatedBackups = append(generatedBackups, backupItem)
		}
	}
	return generatedBackups
}

func url_dot_path_backup(urls []string, backupWords []string) []string {
	var generatedBackups []string

	for _, item := range urls {
		u, err := url.Parse(item)
		if err != nil {
			panic(err)
		}

		if len(u.Path) == 0 || u.Path == "/" || len(u.Hostname()) == 0 {
			continue
		}

		var backupItem string
		for _, word := range backupWords {

			var dotedPath string
			var otherPath string
			if strings.Contains(u.Path, ".") {
				if strings.Count(u.Path, "/") > 1 {
					lastIndexOfSlash := strings.LastIndex(u.Path, "/")
					otherPath = (u.Path[:lastIndexOfSlash])

					dotedPath = strings.Replace(u.Path[lastIndexOfSlash:], "/", "/.", 1)

				} else {
					dotedPath = strings.Replace(u.Path, "/", "/.", 1)
				}

			}

			if len(u.Scheme) != 0 { // getting rid of / of rthe upath
				backupItem = u.Scheme + "://" + u.Host + otherPath + dotedPath + word
			} else {
				backupItem = u.Host + "/." + otherPath + dotedPath + word
			}

			generatedBackups = append(generatedBackups, backupItem)
		}
	}
	return generatedBackups
}
