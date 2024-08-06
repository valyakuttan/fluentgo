package ioexamples


import (
    "fmt"
    "log"
    "os"
)

func ReadEntrieFile() {
    content, err := os.ReadFile("ioexamples/read_file_line_by_line.go")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))
}
