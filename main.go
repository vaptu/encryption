package main

import (
    "Desktop/crypto"
    "Desktop/util"
    "flag"
    "io/ioutil"
    "log"
    "strings"
)

const keyFileName = "private.pem"

var fileDir string
var t = flag.String("t", "encode", "encode 文件加密\ndecode文件解密")
var s = flag.String("s", "", "源文件夹")
var d = flag.String("d", "./", "输出文件夹")

func main() {
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    ParseArgs()
    var isExists bool
    var err error
    if isExists, err = util.CheckPrivateKeyExists(keyFileName); err != nil {
        log.Fatal("无法读取或私钥文件: ", err)
    }
    if isExists == false {
        err := crypto.GenAesKey(32, keyFileName)
        if err != nil {
            log.Fatal("无法创建私钥文件: ", err)
        }
    }

    fileDir = *s
    if GetLastCharset(fileDir) != "/" {
        fileDir += "/"
    }

    if *t == "encode" {
        EncodeDir(fileDir)
    } else {
        DecodeDir(fileDir)
    }
}

func ParseArgs() {
    flag.Parse()

    if *s == "" {
        log.Fatal("缺少输入文件夹 --help Usage ")
    }

    if *t != "encode" && *t != "decode" {
        log.Fatal("-t 参数类型错误 --help Usage ")
    }
}

func GetLastCharset(data string) string {
    str := []rune(data)
    return string(str[len(str)-1:])
}

func EncodeDir(dir string) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal("无法读取目标目录: ", err)
    }

    for _, item := range files {
        if item.IsDir() {
            EncodeDir(dir + item.Name() + "/")
            continue
        }
        newFilePath := strings.Replace(dir, fileDir, "", 1) + item.Name()
        sourcePath := dir + item.Name()
        savePath := *d + "secret/" + newFilePath + ".secret"

        log.Println("Encodeing... ", sourcePath)
        if err := crypto.EncryptFile(sourcePath, savePath, keyFileName); err != nil {
            log.Print(err)
        }
    }
}

func DecodeDir(dir string) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal("无法读取目标目录: ", err)
    }

    for _, item := range files {
        if item.IsDir() {
            DecodeDir(dir + item.Name() + "/")
            continue
        }
        newFilePath := strings.Replace(dir, fileDir, "", 1) + item.Name()
        sourcePath := dir + item.Name()
        savePath := *d + "source/" + strings.Replace(newFilePath, ".secret", "", 1)

        log.Println("Decodeing... ", sourcePath)
        if err := crypto.DecryptFile(sourcePath, savePath, keyFileName); err != nil {
            log.Print(err)
        }
    }
}
