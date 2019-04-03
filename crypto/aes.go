package crypto

import (
    "Desktop/compression"
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "io/ioutil"
    "log"
    "math/rand"
    "os"
    filepath2 "path/filepath"
    "time"
)

func EncryptFile(filepath string, savepath string, keyFileName string) error {
    loadKey, err := ioutil.ReadFile(keyFileName)
    if err != nil {
        return err
    }
    fileData, err := ioutil.ReadFile(filepath)
    if err != nil {
        return err
    }
    data, err := AesEncrypt(fileData, loadKey)
    if err != nil {
        log.Fatal(err)
    }


    // 进行压缩
    cData, err := compression.Comporession(data)
    if err != nil {
        return err
    }

    absPath := filepath2.Dir(savepath)
    if err := os.MkdirAll(absPath, 0755); err != nil {

    }

    if err := ioutil.WriteFile(savepath, *cData, 0644); err != nil {
        return err
    }
    return nil
}

func DecryptFile(filepath string, savepath string, keyFileName string) error {
    loadKey, err := ioutil.ReadFile(keyFileName)
    if err != nil {
        return err
    }
    fileData, err := ioutil.ReadFile(filepath)
    if err != nil {
        return err
    }

    // 进行解压
    uData, err := compression.UnComporession(fileData)
    if err != nil {
        log.Fatal(uData)
        return err
    }

    data, err := AesDecrypt(*uData, loadKey)
    if err != nil {
        log.Fatal("Error: 私钥不正确.", err)
    }

    absPath := filepath2.Dir(savepath)
    if err := os.MkdirAll(absPath, 0755); err != nil {
    }

    if err := ioutil.WriteFile(savepath, []byte(data), 0644); err != nil {
        return err
    }

    return nil
}

func GenAesKey(bits int, keyFileName string) error {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < bits; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }

    if err := ioutil.WriteFile(keyFileName, result, 0644); err != nil {
        return err
    }

    return nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    origData = PKCS5Padding(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)
    return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)
    return origData, nil
}
