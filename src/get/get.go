/* 
Package main implements the main procedure of impg.

The syntax of the user input is:
    
    $./get <si|st|it> <params>
           <si|st>  :   <kwd> <dl_path>
           <it>     :   <img_path> [dl_path]
*/

package main

import (
    "fmt"
    "os"
    "strconv"
)

func usage(progName string) {
    fmt.Printf("Usage: %s <si|st|it|li|lt> <params>\n", progName)
    fmt.Println("          <si|st> : <kwd>      [dl_path]")
    fmt.Println("          <it>    : <img_path> [dl_path]")
    fmt.Println("          <li>    : <cnt>      [dl_path]")
    fmt.Println("          <lt>    : <cnt>      [dl_path]")
    os.Exit(0)
}

func main() {
    if len(os.Args) < 3 {
        usage(os.Args[0])
    }

    var dl_path string
    switch (os.Args[1]) {
    case "si":
        if len(os.Args) < 4 {
            dl_path = "." // working dierctory
        } else {
            dl_path = os.Args[3] // specified by user
        }
        sit(os.Args[2], dl_path, "img")
    case "st":
        if len(os.Args) < 4 {
            dl_path = "." // working dierctory
        } else {
            dl_path = os.Args[3] // specified by user
        }
        sit(os.Args[2], dl_path, "tor")
    case "it":
        var dl_path string
        if len(os.Args) == 3 {
            dl_path = os.Args[2]
        } else {
            dl_path = os.Args[3]
        }
        it(os.Args[2], dl_path)
    case "li":
        if len(os.Args) < 4 {
            dl_path = "." // working dierctory
        } else {
            dl_path = os.Args[3] // specified by user
        }
        i, err := strconv.ParseInt(os.Args[2], 10, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Wrong <cnt>: %s\n", err.Error())
            os.Exit(1)
        }
        lit(int(i), dl_path, "img")
    case "lt":
        if len(os.Args) < 4 {
            dl_path = "." // working dierctory
        } else {
            dl_path = os.Args[3] // specified by user
        }
        i, err := strconv.ParseInt(os.Args[2], 10, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Wrong <cnt>: %s\n", err.Error())
            os.Exit(1)
        }
        lit(int(i), dl_path, "tor")
    default:
        usage(os.Args[0])
    }
}
