package gotool

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/klauspost/compress/zstd"
)

func Compress(src *string, buf io.Writer) error {
	// tar > zstd > buf
	zr, _ := zstd.NewWriter(buf, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	filepath.Walk(*src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		//                 fmt.Println(filepath.ToSlash(file))
		//                 header.Name = filepath.ToSlash(file)
		header.Name = strings.TrimPrefix(file, "csv/")

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open("./" + file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
			data.Close()
		}
		return nil
	})

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce zstd
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}

func CompressFolder(today string) error {
	path := "csv/" + today
	zstdFile := path + ".tar.zst"
	var buf bytes.Buffer
	err := Compress(&path, &buf)
	if err != nil {
		return err
	} else {
		fileToWrite, err := os.OpenFile(zstdFile, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(fileToWrite, &buf); err != nil {
			panic(err)
		}
		fileToWrite.Close()
		os.RemoveAll(path)
	}
	return nil
}

func Uncompress(src io.Reader) error {
	zr, err := zstd.NewReader(src)
	if err != nil {
		return err
	}
	tr := tar.NewReader(zr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		target := filepath.Join(header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			fileToWrite.Close()
		}
	}

	return nil
}

func UncompressFolder(fileName *string) error {
	reg, _ := regexp.Compile("[0-9]...-[0-1][0-9]-[0-3][0-9]")
	date := reg.FindString(*fileName)
	file, err := os.Open(*fileName)
	if err == nil {
		err = Uncompress(file)
	}
	if err != nil {
		return err
	}
	return nil
}
