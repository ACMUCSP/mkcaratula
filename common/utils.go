package common

import (
    "os"
    "io"
)

func CopyFile(src, dst string) (int64, error) {
    sf, err := os.Open(src)
    if err != nil {
        return 0, err
    }
    defer sf.Close()
    df, err := os.Create(dst)
    if err != nil {
        return 0, err
    }
    defer df.Close()
    return io.Copy(df, sf)
}


func MoveFile(src, dst string) error {
    err := os.Rename(src, dst)
    if err != nil {
        _, err = CopyFile(src, dst)
        if err != nil {
            return err
        }
        err = os.Remove(src)
        if err != nil {
            return err
        }
    }
    return nil
}
