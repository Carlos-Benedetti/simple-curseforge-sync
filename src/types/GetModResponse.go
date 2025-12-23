package types

import (
	"time"
)

type Links struct {
	WebsiteUrl string `json:"websiteUrl"`
	WikiUrl    string `json:"wikiUrl"`
	IssuesUrl  string `json:"issuesUrl"`
	SourceUrl  string `json:"sourceUrl"`
}

type Category struct {
	ID               int       `json:"id"`
	GameID           int       `json:"gameId"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	URL              string    `json:"url"`
	IconUrl          string    `json:"iconUrl"`
	DateModified     time.Time `json:"dateModified"`
	IsClass          bool      `json:"isClass"`
	ClassID          int       `json:"classId"`
	ParentCategoryID int       `json:"parentCategoryId"`
	DisplayIndex     int       `json:"displayIndex"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Logo struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Screenshot struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Hash struct {
	Value string `json:"value"`
	Algo  int    `json:"algo"`
}

type GameVersion struct {
	GameVersionName        string    `json:"gameVersionName"`
	GameVersionPadded      string    `json:"gameVersionPadded"`
	GameVersion            string    `json:"gameVersion"`
	GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
	GameVersionTypeID      int       `json:"gameVersionTypeId"`
}

type Dependency struct {
	ModID        int `json:"modId"`
	RelationType int `json:"relationType"`
}

type Module struct {
	Name        string `json:"name"`
	Fingerprint int    `json:"fingerprint"`
}

type LatestFile struct {
	ID                   int           `json:"id"`
	GameID               int           `json:"gameId"`
	ModID                int           `json:"modId"`
	IsAvailable          bool          `json:"isAvailable"`
	DisplayName          string        `json:"displayName"`
	FileName             string        `json:"fileName"`
	ReleaseType          int           `json:"releaseType"`
	FileStatus           int           `json:"fileStatus"`
	Hashes               []Hash        `json:"hashes"`
	FileDate             time.Time     `json:"fileDate"`
	FileLength           int           `json:"fileLength"`
	DownloadCount        int           `json:"downloadCount"`
	FileSizeOnDisk       int           `json:"fileSizeOnDisk"`
	DownloadUrl          string        `json:"downloadUrl"`
	GameVersions         []string      `json:"gameVersions"`
	SortableGameVersions []GameVersion `json:"sortableGameVersions"`
	Dependencies         []Dependency  `json:"dependencies"`
	ExposeAsAlternative  bool          `json:"exposeAsAlternative"`
	ParentProjectFileID  int           `json:"parentProjectFileId"`
	AlternateFileID      int           `json:"alternateFileId"`
	IsServerPack         bool          `json:"isServerPack"`
	ServerPackFileID     int           `json:"serverPackFileId"`
	IsEarlyAccessContent bool          `json:"isEarlyAccessContent"`
	EarlyAccessEndDate   time.Time     `json:"earlyAccessEndDate"`
	FileFingerprint      int           `json:"fileFingerprint"`
	Modules              []Module      `json:"modules"`
}

type LatestFileIndex struct {
	GameVersion       string `json:"gameVersion"`
	FileID            int    `json:"fileId"`
	Filename          string `json:"filename"`
	ReleaseType       int    `json:"releaseType"`
	GameVersionTypeID int    `json:"gameVersionTypeId"`
	ModLoader         int    `json:"modLoader"`
}

type Data struct {
	ID                            int               `json:"id"`
	GameID                        int               `json:"gameId"`
	Name                          string            `json:"name"`
	Slug                          string            `json:"slug"`
	Links                         Links             `json:"links"`
	Summary                       string            `json:"summary"`
	Status                        int               `json:"status"`
	DownloadCount                 int               `json:"downloadCount"`
	IsFeatured                    bool              `json:"isFeatured"`
	PrimaryCategoryID             int               `json:"primaryCategoryId"`
	Categories                    []Category        `json:"categories"`
	ClassID                       int               `json:"classId"`
	Authors                       []Author          `json:"authors"`
	Logo                          Logo              `json:"logo"`
	Screenshots                   []Screenshot      `json:"screenshots"`
	MainFileID                    int               `json:"mainFileId"`
	LatestFiles                   []LatestFile      `json:"latestFiles"`
	LatestFilesIndexes            []LatestFileIndex `json:"latestFilesIndexes"`
	LatestEarlyAccessFilesIndexes []LatestFileIndex `json:"latestEarlyAccessFilesIndexes"`
	DateCreated                   time.Time         `json:"dateCreated"`
	DateModified                  time.Time         `json:"dateModified"`
	DateReleased                  time.Time         `json:"dateReleased"`
	AllowModDistribution          bool              `json:"allowModDistribution"`
	GamePopularityRank            int               `json:"gamePopularityRank"`
	IsAvailable                   bool              `json:"isAvailable"`
	ThumbsUpCount                 int               `json:"thumbsUpCount"`
	Rating                        int               `json:"rating"`
}

type GetModResponse struct {
	Data Data `json:"data"`
}
