package util

import (
	"os"
	"path/filepath"
)

func ReadDirNames(dir_name string) ([]string, error) {
	dp, err := os.Open(dir_name)
	if err != nil {
		return nil, err
	}

	fnames, err := dp.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	return fnames, nil
}

//return relative filename base on dir_name
func ReadDirNamesWithRelativePath(dir_name string) ([]string, error) {
	fnames, err := ReadDirNames(dir_name)
	if err != nil {
		return nil, err
	}

	dir_name = filepath.Dir(dir_name)

	for k, fname := range fnames {
		fnames[k] = dir_name + "/" + fname
	}

	return fnames, nil
}

//return absolute filename base on dir_name
func ReadDirNamesWithAbsoluePath(dir_name string) ([]string, error) {
	fnames, err := ReadDirNames(dir_name)
	if err != nil {
		return nil, err
	}

	dir_name, err = filepath.Abs(filepath.Dir(dir_name))
	if err != nil {
		return nil, err
	}

	for k, fname := range fnames {
		fnames[k] = dir_name + "/" + fname
	}

	return fnames, nil
}

func FileExists(file_name string) bool {
	_, err := os.Stat(file_name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
