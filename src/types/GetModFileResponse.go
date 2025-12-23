package types

import (
	"time"
)

type File struct {
	Id                   int           `json:"id"`
	GameId               int           `json:"gameId"`
	ModId                int           `json:"modId"`
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
	ParentProjectFileId  int           `json:"parentProjectFileId"`
	AlternateFileId      int           `json:"alternateFileId"`
	IsServerPack         bool          `json:"isServerPack"`
	ServerPackFileId     int           `json:"serverPackFileId"`
	IsEarlyAccessContent bool          `json:"isEarlyAccessContent"`
	EarlyAccessEndDate   time.Time     `json:"earlyAccessEndDate"`
	FileFingerprint      int           `json:"fileFingerprint"`
	Modules              []Module      `json:"modules"`
}

type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}

type GetModFileResponse struct {
	Data File `json:"data"`
}
