package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mine_sync/src/curse"
	"mine_sync/src/types/manifest"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

var GAME_VERSION = "1.20.1"
var MOD_LOADER_TYPE = "1"
var DEFAUL_MODS_FOLDER = "/home/bene/workspace/bene/mine-sync/mods"
var DEFAULT_MANIFEST_LINK = "https://raw.githubusercontent.com/Carlos-Benedetti/simple-curseforge-sync/refs/heads/master/manifest.json"

var curseForge = curse.NewCurseForge(
	"$2a$10$H3mU24im8aLckaz47zeHgOof82pJlmnRgo.GooHwtCTpnVJr5bfWS",
	&http.Client{
		Timeout: time.Second * 10,
	},
	432,
)

func main() {

	logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)

	defer logFile.Close()

	// Set the standard logger's output to the multi-writer.
	log.SetOutput(mw)

	// Optional: set log flags (e.g., date, time, short file name)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	MODS_FOLDER := flag.String("out", "", "mods folder")
	MANIFEST_LINK := flag.String("manifest", DEFAULT_MANIFEST_LINK, "mods folder")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var flag")

	flag.Parse()

	if *MODS_FOLDER == "" {
		text, err := curseForge.AskInstancePath()
		if err != nil {
			log.Fatalln(err)
		}
		text = filepath.Clean(text)
		if err != nil {
			panic(err)
		}

		MODS_FOLDER = &text

	}

	manifest, err := downloadManifest(*MANIFEST_LINK)
	if err != nil {
		log.Fatalln(err)
	}

	clearModsFolder(*MODS_FOLDER)
	log.Printf("Total de mods a baixar: %d", len(manifest.Files))
	for _, mod := range manifest.Files {
		modFile, err := curseForge.GetFileByID(mod.ProjectID, mod.FileID)
		if err != nil {
			log.Fatalln(err)
		}
		absPath, _ := filepath.Abs(fmt.Sprintf("%s/%s", *MODS_FOLDER, modFile.FileName))
		log.Printf("Downloading mod: %s to %s", modFile.FileName, absPath)
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

func clearModsFolder(folder string) {

	if len(folder) < 6 {
		log.Fatalln("Caminho é muito curto, medo de ser um diretorio root, por motivos de bug, não deletarei nada, flws")
	}
	entries, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("Falha ao apagar mods antigos da pasta %s", folder)
	}
	log.Printf("Total de mods anteriormente: %d", len(entries))
	for _, d := range entries {
		os.RemoveAll(path.Join([]string{folder, d.Name()}...))
	}

}
