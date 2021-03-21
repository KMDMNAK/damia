package vscode

import (
	"encoding/hex"
	"fmt"
)

func convertHex(str string) string {
	var hexString = ""
	for _, v := range str {
		hexString += fmt.Sprintf("%x", int(v))
	}
	return hexString
}

func convertString(src []byte) string {
	return hex.EncodeToString(src)
}
