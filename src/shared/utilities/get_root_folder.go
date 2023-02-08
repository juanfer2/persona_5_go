package utilities

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetRootFolder() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(b), "../../..")
}

func GetRootStaticFiles() string {
	return GetRootFolder() + "/tmp"
}

func GetUrlServer() string {
	return "http://localhost:" + GetEnv("PORT")
}

func BuildImageUrl(path string) string {
	return fmt.Sprintf("%s/images/%s", GetUrlServer(), path)
}
