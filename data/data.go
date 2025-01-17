package data

import (
	"encoding/json"
	"errors"
	"github.com/jwilder/gojq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	Path int = iota
	File
	Directory
)

// Exists determines whether a given Path, File, or Directory exists
// use constants:
//	 Path int = iota
//	 File
//	 Directory
func Exists(path string, pathType int) (bool, error) {
	fileInfo, err := os.Stat(path)

	switch pathType {
	case 0:
		if err == nil {
			return true, nil
		}
	case 1:
		if !fileInfo.IsDir() && (!os.IsNotExist(err) || os.IsExist(err)) {
			return true, nil
		}
	case 2:
		if fileInfo.IsDir() && (!os.IsNotExist(err) || os.IsExist(err)) {
			return true, nil
		}
	}

	return false, err
}

// IsStringTrue determines if a string is boolean true/false
func IsStringTrue(str string) bool {
	converted, err := strconv.ParseBool(strings.ToLower(str))
	if err == nil {
		return converted
	}
	return false
}

// QueryJson queries a json object for a given path
func QueryJson(obj string, query string) (interface{}, error) {
	parser, err := gojq.NewStringQuery(obj)
	if err != nil {
		return "", err
	}
	res, err := parser.Query(query)
	if err != nil {
		return "", err
	}
	return res, nil
}

// MapHasKey contains checks if a map[T any]string has a given key T
func MapHasKey[T comparable](item map[T]any, key T) bool {
	if _, ok := item[key]; ok {
		return true
	}
	return false
}

// GetEnvVariables returns a map of all environment variables
func GetEnvVariables() map[string]string {
	items := make(map[string]string)
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		items[splits[0]] = splits[1]
	}
	return items
}

// UnmarshalFile loads and unmarshals either a JSON or YAML file
func UnmarshalFile(filename string) (map[string]interface{}, error) {
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	filetype := filename[len(filename)-4:]
	if filetype == "json" {
		err = json.Unmarshal(byteValue, &data)
	} else if filetype == "yaml" || filetype == ".yml" {
		err = yaml.Unmarshal(byteValue, &data)
	} else {
		err = errors.New("invalid file type")
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
