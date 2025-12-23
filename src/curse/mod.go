package curse

import (
	"encoding/json"
	"fmt"
	"log"
	"mine_sync/src/types"
	"mine_sync/src/utils"
	"path/filepath"
)

type Mod struct {
	ID         int
	curseForge CurseForge
	data       *types.Data
}

func newMod(ID int) Mod {
	return Mod{
		ID: ID,
	}
}

func (m *Mod) GetFileByID(fileID int) (types.File, error) {
	return m.curseForge.GetFileByID(m.ID, fileID)
}

func (m *Mod) GetFileId(gameVersion string) (int, error) {

	indexMapped := utils.AssossiateBy(m.data.LatestFilesIndexes, func(it types.LatestFileIndex) (string, types.LatestFileIndex) {
		return it.GameVersion, it
	})

	indexVersion, ok := indexMapped[gameVersion]
	if !ok {
		bytes, err := json.MarshalIndent(m.data.LatestFiles, "", "  ") // Indent with two spaces
		if err != nil {
			log.Fatal(err)
		}
		return 0, fmt.Errorf("Fail to download file from mod %s, game version %s not found, avaliable: \n %s", m.data.Name, gameVersion, string(bytes))
	}

	return indexVersion.FileID, nil
}

func (m *Mod) Download(gameVersion string, dirPath string) error {

	fileId, err := m.GetFileId(gameVersion)
	if err != nil {
		log.Fatalln(err)
	}

	fileVersion, err := m.GetFileByID(fileId)
	if err != nil {
		log.Fatalln(err)
	}

	absPath, _ := filepath.Abs(fmt.Sprintf("%s/%s", dirPath, fileVersion.FileName))

	log.Printf("Dowloading mod: %d", m.ID)

	err = m.curseForge.FetchFile(fileVersion.DownloadUrl, absPath)

	if err != nil {
		return fmt.Errorf("fail to find mod: %v", err)
	}

	return nil

}
