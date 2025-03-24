package util

import (
    "bytes"
    "io/ioutil"
    "path/filepath"
)

func ReadFile(file string) string {
    data, err := ioutil.ReadFile(filepath.FromSlash(file))
    Check(err)
    // replace CR LF \r\n (windows) with LF \n (unix)
    data = bytes.Replace(data, []byte{13, 10}, []byte{10}, -1)
    // replace CF \r (mac) with LF \n (unix)
    data = bytes.Replace(data, []byte{13}, []byte{10}, -1)
    return string(data)
}

func WriteFile(file string, text string) {
    // Only the owner can read and write. Everyone else can only read. No one can execute the file.
    err := ioutil.WriteFile(filepath.FromSlash(file), []byte(text), 0644)
    Check(err)
}
