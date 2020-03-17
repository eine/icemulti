package lib

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type isICEM struct {
	hasBin  bool
	hasJson bool
}

func isValidICEM(v isICEM) bool {
	return (v.hasBin == true) && (v.hasJson == true)
}

func ReadTGZ(path string) (bin []byte, meta t_meta, err error) {

	flagValid := isICEM{false, false}

	f, err := os.Open(path)
	if err != nil {
		return nil, EmptyMeta(), err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return nil, EmptyMeta(), err
	}
	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, EmptyMeta(), err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			switch header.Name {
			case "./bin":
				flagValid.hasBin = true
				var binBuf bytes.Buffer
				io.Copy(bufio.NewWriter(&binBuf), tarReader)
				bin = binBuf.Bytes()
			case "./json":
				flagValid.hasJson = true
				var jsonBuf bytes.Buffer
				io.Copy(bufio.NewWriter(&jsonBuf), tarReader)
				json.Unmarshal(jsonBuf.Bytes(), &meta)
				//			    default:
			}
		default:
			fmt.Printf("%s : %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				header.Name,
			)
		}
	}
	if isValidICEM(flagValid) {
		return bin, meta, err
	}
	return nil, EmptyMeta(), nil
}

func ReadBIN(path string) (bin []byte, err error) {
	bin, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bin, err
}
