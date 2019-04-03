package compression

import (
    "bytes"
    "compress/zlib"
    "io/ioutil"
    "log"
)

func Comporession(data []byte) (*[]byte, error) {
    var buf bytes.Buffer
    w := zlib.NewWriter(&buf)
    if _, err := w.Write(data); err != nil {
        return nil, err
    }
    if err := w.Close(); err != nil {
        return nil, err
    }
    cData := buf.Bytes()

    return &cData, nil
}

func UnComporession(data []byte) (*[]byte, error) {
    b := bytes.NewReader(data)
    r, err := zlib.NewReader(b)
    if err != nil {
        return nil, err
    }

    uData, err := ioutil.ReadAll(r)
    if err != nil {
        log.Fatal(data)
        return nil, err
    }
    r.Close()

    return &uData, nil
}
