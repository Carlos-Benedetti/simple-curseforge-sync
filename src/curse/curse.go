package curse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mine_sync/src/types"
	"net/http"
	"net/url"
	"os"
)

type CurseForge struct {
	API_KEY string
	Client  *http.Client
	GameId  int
}

func (c CurseForge) SearchMod(searchFilter string) (types.Data, error) {

	fullUrl := c.base_url()
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

	fullUrl := m.base_url()
	fullUrl.Path = fmt.Sprintf("/v1/mods/%d", modID)
	err := m.innerFetch(fullUrl.String(), "GET", nil, &getModResponse)

	if err != nil {
		return getModResponse, fmt.Errorf("fail to find mod: %v", err)
	}

	return getModResponse, nil
}

func (c *CurseForge) GetFileByID(modID int, fileID int) (types.File, error) {
	var getModFileResponse types.GetModFileResponse
	log.Printf("fetching mod files: %d", modID)

	fullUrl := c.base_url()
	fullUrl.Path = fmt.Sprintf("/v1/mods/%d/files/%d", modID, fileID)
	err := c.innerFetch(fullUrl.String(), "GET", nil, &getModFileResponse)

	if err != nil {
		return types.File{}, fmt.Errorf("fail to find mod: %v", err)
	}

	return getModFileResponse.Data, nil
}

func (c CurseForge) base_url() *url.URL {
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
