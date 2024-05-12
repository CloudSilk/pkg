package zipWrapper

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const ZipExtString = ".zip"

type Archive struct {
}

func NewArchive() *Archive {
	return new(Archive)
}

func (a *Archive) PackageFolder(path, pwd string) (string, error) {
	zipFile := path + ZipExtString
	outFile, err := os.Create(zipFile)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	w := zip.NewWriter(outFile)
	defer w.Close()

	addFilesInFolder(w, path, "")

	return zipFile, nil
}

func (a *Archive) UnPackage(r io.Reader, pwd, path string) error {
	payload, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(payload)
	zr, err := zip.NewReader(reader, int64(len(payload)))
	if err != nil {
		return err
	}
	for _, f := range zr.File {
		fp := filepath.Join(path, f.Name)
		log.Printf("zip unpackage file: %+v", f)
		if f.FileInfo().IsDir() {
			log.Printf("zip unpackage file: %+v", f)
			if e := os.MkdirAll(fp, 0755); e != nil {
				log.Println("Error: cannot mkdir ", fp)
			}
			continue
		}
		log.Println("zip unpackage file paremt folder: ", filepath.Dir(fp))
		if e := os.MkdirAll(filepath.Dir(fp), 0755); e != nil {
			log.Println("Error: cannot mkdir ", fp, e)
		}

		outFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Println("cannot crete file:", fp, err)
			continue
		}
		defer outFile.Close()

		rc, er := f.Open()
		if er != nil {
			log.Println("cannot open zip file:", f.Name, err)
			continue
		}
		defer rc.Close()
		_, er = io.Copy(outFile, rc)
		if er != nil {
			log.Println("cannot read zip file:", f.Name, er.Error())
			continue
		}
	}

	return nil
}

func addFilesInFolder(w *zip.Writer, basePath, baseInZip string) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fileWithAbsPath := filepath.Join(basePath, file.Name())
		fileWithRelativePath := filepath.Join(baseInZip, file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(fileWithAbsPath)
			if err != nil {
				fmt.Println(err)
			}
			fileWithRelativePath = filepath.ToSlash(fileWithRelativePath)
			f, err := w.Create(fileWithRelativePath)
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {
			addFilesInFolder(w, fileWithAbsPath, fileWithRelativePath)
		}
	}
}
