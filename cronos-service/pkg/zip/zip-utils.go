package ziputils

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func CreateZipFile(source, dest string) error {

	zipFile, err := os.Create(filepath.Join(dest, time.Now().Format("2006-01-02T15-04-05")+".zip"))
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if path == source {
			return nil
		}

		if !info.IsDir() {

			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			p, err := filepath.Rel(source, path)
			if err != nil {
				return err
			}

			destFile, err := zipWriter.Create(p)
			if err != nil {
				return err
			}

			_, err = io.Copy(destFile, sourceFile)
			return err
		} else {

			_, err = zipWriter.CreateHeader(&zip.FileHeader{
				Name:     filepath.Base(path) + "/",
				Method:   zip.Store,
				Modified: info.ModTime(),
			})

			return err
		}

	})

	return err
}
