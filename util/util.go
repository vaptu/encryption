package util

import (
    "os"
)


func CheckPrivateKeyExists(privateFileName string) (bool, error) {
    _, err := os.Stat(privateFileName)
    if err == nil {
        return true, err
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return false, err
}
