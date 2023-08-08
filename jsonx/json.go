package jsonx

import (
	json "github.com/json-iterator/go"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

func Encode(v interface{}) string {
	if str, err := MarshalToString(v); err == nil {
		return str
	} else {
		panic(err)
	}
}
