package archive

import (
	"io"
	"path/filepath"

	zip "github.com/CloudSilk/pkg/archive/zip"
	"github.com/CloudSilk/pkg/utils/log"
)

type Archive interface {
	PackageFolder(path, pwd string) (string, error)
	UnPackage(r io.Reader, pwd, path string) error
}

var DefaultArch Archive

func init() {
	if DefaultArch == nil {
		DefaultArch = zip.NewArchive()
	}
}

func NewUnArchiver(file string) Archive {
	switch filepath.Ext(file) {
	case zip.ZipExtString:
		log.Println(nil, "find zip decoder")
		return zip.NewArchive()
	}
	return nil
}

func PackageFolder(path, pwd string) (string, error) {
	file, err := DefaultArch.PackageFolder(path, pwd)
	if err != nil {
		log.Println(nil, "PackageFolder failed:", err)
		return "", err
	}
	return file, nil
}

func UnPackageFolder(file io.Reader, pwd, path string) error {
	if err := DefaultArch.UnPackage(file, pwd, path); err != nil {
		log.Println(nil, "UnPackageFolder failed:", err)
		return err
	}
	return nil
}
