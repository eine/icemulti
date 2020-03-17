package lib

import (
	"encoding/json"
	"fmt"
)

/*
func (f *t_binFile) MarshalJSON() ([]byte, error) {
  type Alias t_binFile
	return json.Marshal(&struct {
		Bin []byte `json:"bin"`
		*Alias
	}{
		Bin: []byte{0,1,2,3,4,5},
		Alias: (*Alias)(f),
	})
}
//	"encoding/hex"
//fmt.Println(hex.EncodeToString(f.Bin[0:16]));
*/

func List2JSON(l t_binFileList) []byte {
	b, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}
