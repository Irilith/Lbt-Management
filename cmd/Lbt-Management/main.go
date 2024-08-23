package main

import (
	"Lbt-Management/pkg/qbittorrent/requests"
	"Lbt-Management/pkg/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func int() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	username := os.Getenv("QBIT_USERNAME")
	password := os.Getenv("QBIT_PASSWORD")
	sid, err := requests.GetAuth(username, password)
	if err != nil {
		panic(err)
	}

	info, err := requests.GetTorrentsInfos(sid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SID: %s\n", sid)
	for _, item := range info {
		files := requests.GetFiles(sid, item.Hash)

		for _, file := range files {
			if strings.Contains(file.Name, ".epub") || strings.Contains(file.Name, ".pdf") || strings.Contains(file.Name, ".mobi") || strings.Contains(file.Name, ".cbz") {
				parentName, fileName := utils.SplitFolder(file.Name)
				if parentName == "" {
					parentName = fileName
				}
				parentNameSmart := utils.GenerateFolder(parentName)
				base := utils.GetParent(item.ContentPath)
				if parentNameSmart != base && fileName != parentNameSmart && fileName != base {
					// if the parentName contains '.' which is a file
					if !strings.Contains(parentName, ".") {
						newPath := fmt.Sprintf("%s", strings.TrimSpace(parentNameSmart))
						err := requests.RenameFolder(sid, item.Hash, parentName, newPath)
						if err != nil {
							fmt.Println(fmt.Sprintf("Error renaming %s to %s, base: %s\nParent: %s", fileName, newPath, base, parentName))
							panic(err)
						}
					} else if strings.TrimSpace(parentName) == strings.TrimSpace(file.Name) && fileName == parentName && strings.Contains(parentName, ".") {
						newPath := fmt.Sprintf("%s/%s", strings.TrimSpace(parentNameSmart), strings.TrimSpace(fileName))
						err := requests.RenameFile(sid, item.Hash, fileName, newPath)
						if err != nil {
							fmt.Println(fmt.Sprintf("Error renaming %s to %s, base: %s\nParent: %s", fileName, newPath, base, parentName))
							panic(err)
						}
					}
				}
			}
		}
	}
}
