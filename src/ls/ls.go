/*
 * File: dir.go
 * 5/3/2014
 * Edited by Wangdeqin
 * List filenames in a specific path
 */

package ls

import (
    "os"
    "errors"
    "fmt"
)

func GetFilenames(path string) ([]string, error) {
    dir, err := os.Open(path)
    checkErr(err)
    defer dir.Close()

    state, err := dir.Stat()
    checkErr(err)

    if !state.IsDir() {
        return nil, errors.New(fmt.Sprintf("%s is not a dir", path))
    }

    fis, err := dir.Readdir(-1)
    checkErr(err)

    filenames := make([]string, 0)

    for _, fi := range fis {
        filenames = append(filenames, fi.Name())
    }

    return filenames, nil
}

func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}
