package passwork

import "encoding/json"

type PasswordResponse struct {
	Status string
	Code   string // passwordNull, accessDenied
	Data   PasswordResponseData
}

type PasswordSearchResponse struct {
	Status string
	Code   string
	Data   []PasswordResponseData
}

// UnmarshalJSON implements custom unmarshaling to support both API v1 and v4 formats
// API v1 uses "items" field, API v4 uses "Data" field
func (p *PasswordSearchResponse) UnmarshalJSON(data []byte) error {
	// Try API v1 format first (items field)
	var v1Format struct {
		Items []PasswordResponseData `json:"items"`
	}
	if err := json.Unmarshal(data, &v1Format); err == nil && v1Format.Items != nil {
		p.Data = v1Format.Items
		return nil
	}
	
	// Try API v4 format (Status, Code, Data fields)
	type Alias PasswordSearchResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	return json.Unmarshal(data, aux)
}

type PasswordResponseData struct {
	VaultId            string
	FolderId           string
	Custom             []PasswordCustomData
	Id                 string
	Name               string
	Login              string
	CryptedPassword    string
	CryptedKey         string
	Description        string
	Url                string
	Color              int
	Attachments        []PasswordAttachmentData
	Tags               []string
	Path               []PathData
	Access             string
	AccessCode         int
	Shortcut           PasswordShortcutData
	LastPasswordUpdate int
	UpdatedAt          string
	IsFavorite         bool
}

type PasswordSearchRequest struct {
	Query         string   `json:"query"`
	VaultId       string   `json:"vaultId"`
	Colors        []int    `json:"colors,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	IncludeShared bool     `json:"includeShared,omitempty"`
}

type PasswordRequest struct {
	Name            string                   `json:"name"`
	Login           string                   `json:"login,omitempty"`
	CryptedPassword string                   `json:"cryptedPassword,omitempty"`
	Url             string                   `json:"url,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Custom          []PasswordCustomData     `json:"custom,omitempty"`
	Color           int                      `json:"color,omitempty"`
	Attachments     []PasswordAttachmentData `json:"attachments,omitempty"`
	Tags            []string                 `json:"tags,omitempty"`
	MasterHash      string                   `json:"masterHash,omitempty"`
	VaultId         string                   `json:"vaultId"`
	FolderId        string                   `json:"folderId,omitempty"`
	ShortcutId      string                   `json:"shortcutId,omitempty"`
}

type PasswordCustomData struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

type PasswordAttachmentData struct {
	Name          string `json:"name,omitempty"`
	Id            string `json:"id,omitempty"`
	EncryptedKey  string `json:"encryptedKey,omitempty"`
	Hash          string `json:"hash,omitempty"`
	EncryptedData string `json:"encryptedData,omitempty"`
}

type PasswordShortcutData struct {
	Id         string
	PasswordId string
	VaultId    string
	FolderId   string
	Access     string
	AccessCode int
	CryptedKey string
}
