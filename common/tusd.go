package common

import (
	b64 "encoding/base64"
	"errors"
	"strings"
)

func GetFileType(metadata string) (string, error) {
	headerUpload := strings.Split(metadata, ",")
	if headerUpload[1] == "" {
		return "", errors.New("no header uplaod detected")
	}

	fileType := strings.Split(headerUpload[1], " ")
	if fileType[1] == "" {
		return "", errors.New("no filetype detected")
	}

	decodedFiletype, err := b64.StdEncoding.DecodeString(fileType[1])
	if err != nil {
		return "", err
	}

	return string(decodedFiletype), nil
}


func IsFileTypeVideo(filetype string) bool {
	return strings.Contains(filetype, "video")
}