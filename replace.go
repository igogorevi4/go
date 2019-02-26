// Code creates copy of specific file where it changes some pattern to another pattern.
//
// Usage:
// echo /path/to/file currentPattern newPatterm | go run replace.go
//

package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
)

var TargetFile string
var Pattern string
var PatternToReplace string
var Writer string

func main() {
    fmt.Scan(&TargetFile, &Pattern, &PatternToReplace)
    //fmt.Println(TargetFile, Pattern, PatternToReplace)

    inFile, _ := os.Open(TargetFile)
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
        scanner.Split(bufio.ScanLines) 
    
    i := 1 
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        matched, _ := regexp.MatchString(Pattern, scanner.Text()) 
        if matched {
            // to newFile
            fmt.Println("OK")
        } else {
            //scanner.Text() -> to newFile
            fmt.Println("Not OK")
        }
        fmt.Println(i)
        i++
    }
}