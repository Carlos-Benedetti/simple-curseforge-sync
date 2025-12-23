package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
var DEFAULT_MANIFEST_LINK = "https://raw.githubusercontent.com/Carlos-Benedetti/simple-curseforge-sync/refs/heads/master/manifest.json"

var curseForge = curse.CurseForge{
	API_KEY: "$2a$10$H3mU24im8aLckaz47zeHgOof82pJlmnRgo.GooHwtCTpnVJr5bfWS",
	Client: &http.Client{
		Timeout: time.Second * 10,
	},
	GameId: 432,
}

func main() {

	MODS_FOLDER := flag.String("out", "", "mods folder")
	MANIFEST_LINK := flag.String("manifest", DEFAULT_MANIFEST_LINK, "mods folder")

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

	manifest, err := downloadManifest(*MANIFEST_LINK)
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

func downloadManifest(url string) (manifest.Manifest, error) {
	var manifest manifest.Manifest
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	resp, err := curseForge.Client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalf("API returned non-OK status: %d", resp.StatusCode)
	}

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	err = json.Unmarshal(respbody, &manifest)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %s", err)
	}

	return manifest, nil
}
