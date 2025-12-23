package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"mine_sync/src/curse"
	"mine_sync/src/types/manifest"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var GAME_VERSION = "1.20.1"
var MOD_LOADER_TYPE = "1"
var DEFAUL_MODS_FOLDER = "/home/bene/workspace/bene/mine-sync/mods"

var curseForge = curse.CurseForge{
	API_KEY: "$2a$10$H3mU24im8aLckaz47zeHgOof82pJlmnRgo.GooHwtCTpnVJr5bfWS",
	Client: &http.Client{
		Timeout: time.Second * 10,
	},
	GameId: 432,
}

func main() {

	MODS_FOLDER := flag.String("out", "", "mods folder")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var flag")

	flag.Parse()

	// Access the values (dereference pointers for flag.String, flag.Int, flag.Bool)
	fmt.Println("MODS_FOLDER:", *MODS_FOLDER)
	if *MODS_FOLDER == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Caminho da pasta nçao definipor por argumento -out")
		fmt.Print("Caminho da pasta sua mods(~/Documents/curseforge/minecraft/Instances/mockpackname/mods): ")

		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		MODS_FOLDER = &text
	}

	log.Println(MODS_FOLDER)
	manifestPath, _ := filepath.Abs(fmt.Sprintf("./manifest.json"))

	bytes, err := os.ReadFile(manifestPath)
	if err != nil {
		log.Fatalln(err)
	}
	var manifest manifest.Manifest
	err = json.Unmarshal(bytes, &manifest)
	if err != nil {
		log.Fatalln(err)
	}

	for _, mod := range manifest.Files {
		modFile, err := curseForge.GetFileByID(mod.ProjectID, mod.FileID)
		if err != nil {
			log.Fatalln(err)
		}
		absPath, _ := filepath.Abs(fmt.Sprintf("%s/%s", *MODS_FOLDER, modFile.FileName))
		err = curseForge.FetchFile(modFile.DownloadUrl, absPath)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

type ModFile struct {
	FileId int
}

// func searchMod(name string) {
// 	var getModFileResponse types.GetModFileResponse
// 	log.Printf("fetching mod files: %d", m.ID)

// 	fullUrl := base_url()
// 	fullUrl.Path = fmt.Sprintf("/v1/mods/%d/files/%d", m.ID, fileID)
// 	err := innerFetch(fullUrl.String(), "GET", nil, &getModFileResponse)

// 	if err != nil {
// 		return types.File{}, fmt.Errorf("fail to find mod: %v", err)
// 	}

// 	return getModFileResponse.Data, nil
// }
