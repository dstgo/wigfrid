package assets

import (
	"embed"
	"io"
	"text/template"
)

// assets fs embed all the assets files into the binary
//
//go:embed *
var fs embed.FS

const (
	BannerFile = "banner.txt"
)

// WriteBanner write the banner into the give writer and parse it with given data.
func WriteBanner(writer io.Writer, data any) error {
	tmpl, err := template.New("banner").ParseFS(fs, BannerFile)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(writer, BannerFile, data)
}
