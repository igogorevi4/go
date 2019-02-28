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
    "strings"
)

var TargetFile, NewFilePath, Pattern, PatternToReplace, CurrentLine, NewLine string

func main() {
    fmt.Scan(&TargetFile, &Pattern, &PatternToReplace)
    inFile, _ := os.Open(TargetFile)
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
        scanner.Split(bufio.ScanLines) 
    
    // new file creating
    NewFilePath := fmt.Sprintf(TargetFile + ".new")
    var NewFile, Err = os.Create(NewFilePath)
    _ = NewFile
    _ = Err

    for scanner.Scan() {
        CurrentLine := scanner.Text()
        matched, _ := regexp.MatchString(Pattern, scanner.Text()) 
        if matched {
            NewLine := strings.Replace(CurrentLine, Pattern, PatternToReplace, -1)
            file, Err := os.OpenFile(NewFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
            _ = Err
            defer file.Close()
            _, Err = file.WriteString(NewLine + "\n")
        } else {
            file, Err := os.OpenFile(NewFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
            _ = Err
            defer file.Close()
            _, Err = file.WriteString(CurrentLine + "\n")
        }
    }
}