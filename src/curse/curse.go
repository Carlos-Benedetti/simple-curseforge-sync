package curse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mine_sync/src/config"
	"mine_sync/src/types"
	"mine_sync/src/utils"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/manifoldco/promptui"
)

type CurseForge struct {
	API_KEY    string
	Client     *http.Client
	GameId     int
	userFolder string
}

func NewCurseForge(apiKey string, Client *http.Client, GameId int) CurseForge {
	var userFolder string

	switch runtime.GOOS {
	case "windows":
		userFolder = os.Getenv("USERPROFILE")
	}

	curseForgeFolder, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s", userFolder, "curseforge", "minecraft", "Instances"))

	return CurseForge{
		API_KEY:    apiKey,
		Client:     Client,
		GameId:     GameId,
		userFolder: curseForgeFolder,
	}
}

func (c CurseForge) AskInstancePath() (string, error) {
	fmt.Println("Caminho da pasta não definipor por argumento -out")
	var path *string
	lastInstance := config.GetLastSession()
	if lastInstance != nil && *lastInstance != "" {
		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("Atualizar mods em %s", *lastInstance),
			IsConfirm: true,
			Default:   "Y",
		}

		response, no := prompt.Run()
		if no == nil && response != "n" && response != "N" {
			path = lastInstance
		}
	}
	if path == nil {

		prompt := promptui.Prompt{
			Label: "Caminho da pasta sua mods(~/Documents/curseforge/minecraft/Instances/mockpackname/mods)",
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Falha ao selecionar instancia %v\n", err)
			return "", fmt.Errorf("Falha ao selecionar instancia %v\n", err)
		}
		path = &result
	}
	config.SetLastCSession(*path)
	return *path, nil
}

func (c CurseForge) SearchInstancePath() (string, error) {

	entries, err := os.ReadDir(c.userFolder)
	// entries, err := utils.FilePathWalkDir(c.userFolder)
	nameEntries := utils.Map(entries, func(it os.DirEntry) string {
		return it.Name()
	})
	if err != nil {
		return "", fmt.Errorf("Falha ao buscar instancias do curseforge")
	}
	for _, entry := range nameEntries {
		fmt.Println(entry)
	}

	prompt := promptui.Select{
		Label: "Qual isntancia?",
		Items: nameEntries,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("Select prompt falhou %v\n", err)
	}
	return result, nil
}
func (c CurseForge) SearchMod(searchFilter string) (types.Data, error) {

	fullUrl := c.baseUrl()
	fullUrl.Path = "/v1/mods/search"
	q := fullUrl.Query()
	q.Add("gameId", fmt.Sprintf("%d", c.GameId))
	q.Add("searchFilter", searchFilter)

	fullUrl.RawQuery = q.Encode()

	var searchModsResponse types.SearchModsResponse
	err := c.innerFetch(fullUrl.String(), "GET", nil, &searchModsResponse)

	if err != nil {
		return types.Data{}, fmt.Errorf("Fail to fetch search mod on %s", fullUrl.String())
	}

	if len(searchModsResponse.Data) < 1 {
		return types.Data{}, fmt.Errorf("Mod %s not found", searchFilter)
	}

	firstMod := searchModsResponse.Data[0]

	return firstMod, nil

}

func (m *CurseForge) getData(modID int) (types.GetModResponse, error) {
	var getModResponse types.GetModResponse
	log.Printf("fetching mod Data: %d", modID)

	fullUrl := m.baseUrl()
	fullUrl.Path = fmt.Sprintf("/v1/mods/%d", modID)
	err := m.innerFetch(fullUrl.String(), "GET", nil, &getModResponse)

	if err != nil {
		return getModResponse, fmt.Errorf("fail to find mod: %v", err)
	}

	return getModResponse, nil
}

func (c *CurseForge) GetFileByID(modID int, fileID int) (types.File, error) {
	var getModFileResponse types.GetModFileResponse
	// log.Printf("fetching mod files: %d", modID)

	fullUrl := c.baseUrl()
	fullUrl.Path = fmt.Sprintf("/v1/mods/%d/files/%d", modID, fileID)
	err := c.innerFetch(fullUrl.String(), "GET", nil, &getModFileResponse)

	if err != nil {
		return types.File{}, fmt.Errorf("fail to find mod with ID %d: %v", modID, err)
	}

	return getModFileResponse.Data, nil
}

func (c CurseForge) baseUrl() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   "api.curseforge.com",
	}
}

func (c CurseForge) FetchFile(url string, filepath string) error {

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", c.API_KEY)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalf("API returned non-OK status: %d", resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c CurseForge) innerFetch(url string, method string, body map[string]string, response any) error {

	req, err := http.NewRequest(method, url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", c.API_KEY)

	resp, err := c.Client.Do(req)
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

	err = json.Unmarshal(respbody, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %s", err)
	}

	return nil
}
