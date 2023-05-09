package models

import "minecraft/cmd/internal/config"

type Mod struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Link         string `json:"link"`
	Version      string `json:"version"`
	ImgLink      string `json:"imgLink"`
	DownloadLink string `json:"downloadLink"`
}

func NewMode(title, description, link, version, imgLink, downloadLink string) *Mod {
	config.LastID++
	return &Mod{ID: config.LastID, Title: title, Description: description, Link: link, Version: version, ImgLink: imgLink, DownloadLink: downloadLink}
}
