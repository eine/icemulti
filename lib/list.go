package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type t_meta struct {
	Board       string `json:"board"`
	Compression string `json:"compression"`
	Device      string `json:"device"`
	UID         string `json:"uid"`
	URL         string `json:"url"`
}

func EmptyMeta() (m t_meta) {
	return m
}

func NewMeta(i, u, d, b string) (m t_meta) {
	m.UID = i
	m.URL = u
	m.Device = d
	m.Board = b
	return m
}

type t_binFile struct {
	Path    string    `json:"path"`
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modtime"`
	Meta    t_meta    `json:"meta"`
	Bin     []byte    `json:"bin"`
}

func NewBinFile(p, n string, s int64, t time.Time, m t_meta, b []byte) (f t_binFile) {
	f.Path = p
	f.Name = n
	f.Size = s
	f.ModTime = t
	f.Meta = m
	f.Bin = b
	return f
}

func EmptyBinFileFromPathAndInfo(p string, i os.FileInfo) (f t_binFile) {
  return NewBinFile(
	  p,
	  strings.TrimSuffix(i.Name(), filepath.Ext(p)),
	  i.Size(),
	  i.ModTime(),
	  EmptyMeta(),
	  []byte{},
  )
}

type t_binFileList map[string]t_binFile

func (l *t_binFileList) Add(f t_binFile) {
	(*l)[f.Name] = f;
}

func List(args []string, rec bool) (files t_binFileList, err error) {

	files = make(t_binFileList)

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println(path, "does not exists. Skipping...")
			} else {
				fmt.Println("Skipping...", err)
			}
			return nil
		} else {
			if info.IsDir() {
				return nil
			} else {
				switch ext := filepath.Ext(path); ext {
				case ".bin", ".bit":
          bin, err := ReadBIN(path)
					if err != nil {
						return err
					}
					if bin != nil {
						f := EmptyBinFileFromPathAndInfo(path, info)
						f.Bin = bin
						files.Add(f)
					}
				case ".tgz", ".tar.gz":
					bin, meta, err := ReadTGZ(path)
					if err != nil {
						return err
					}
					if bin != nil {
						f := EmptyBinFileFromPathAndInfo(path, info)
						f.Bin = bin
						f.Meta = meta
					  files.Add(f)
					}
				case ".json", ".icem":
					fmt.Println(path)
//        default:
				}
			}
			return nil
		}
	}

	for _, a := range args {
		var err error
		if rec {
			err = filepath.Walk(a, func(path string, info os.FileInfo, err error) error {
				return walkFunc(path, info, err)
			})
		} else {
			info, err := os.Stat(a)
			err = walkFunc(a, info, err)
		}
		if err != nil {
			fmt.Printf("walk error [%v]\n", err)
			return files, err
		}
	}
	return files, nil
}
