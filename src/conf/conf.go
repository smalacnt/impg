/*
 * File: conf.go
 * 2014/11/6
 * Edited by Wangdeqin
 * Read config from a file (conf.xml, create is not existed)
 */
package conf

import (
    "fmt"
    "encoding/xml"
    "os"
    "io/ioutil"
)

type Conf struct {
    XMLName     xml.Name    `xml:"configs"`
    Configs     []Config    `xml:"config"`
}

type Config struct {
    XMLName         xml.Name    `xml:"config"`
    Version         string      `xml:"version,attr"`
    RetryTime       int         `xml:"retryTime"`
    IdChanSize      int         `xml:"idChanSize"`
    ThreadPoolSize  int         `xml:"threadPoolSize"`
}

var RETRY_TIME, ID_CHAN_SIZE, THREAD_POOL_SIZE int

func init() {
    v := &Conf{}
    f, err := os.Open("./conf.xml")
    if err != nil {
        v.Configs = append(v.Configs, Config{Version: "1", RetryTime: 5,
                    IdChanSize: 10, ThreadPoolSize: 5})
        output, err := xml.MarshalIndent(v, "", "    ")
        if err != nil {
            fmt.Printf("Marshal conf failed! \nError: %v\n", err)
            os.Exit(-1)
        }
        file, err := os.Create("./conf.xml")
        if err != nil {
            fmt.Printf("Create file ./conf.xml failed! \nError: %v\n", err)
            os.Exit(-1)
        }
        defer file.Close()
        file.Write(output)
    } else {
        data, err := ioutil.ReadAll(f)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(-1)
        }
        err = xml.Unmarshal(data, v)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(-1)
        }
    }
    RETRY_TIME = v.Configs[0].RetryTime
    ID_CHAN_SIZE = v.Configs[0].IdChanSize
    THREAD_POOL_SIZE = v.Configs[0].ThreadPoolSize
}
