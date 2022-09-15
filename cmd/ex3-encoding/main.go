package main

import (
    "encoding/xml"
    "encoding/json"
    yaml "gopkg.in/yaml.v2"
    "fmt"
)

type cppCheckLocation struct {
        XMLName xml.Name `xml:"location"`
        File    string   `xml:"file,attr"`
        Line    int      `xml:"line,attr"`
        Info    string   `xml:"info,attr"`
}

type cppCheckError struct {
        XMLName   xml.Name            `xml:"error"`
        Id        string              `xml:"id,attr"`
        Severity  string              `xml:"severity,attr"`
        Msg       string              `xml:"msg,attr"`
        Locations []*cppCheckLocation `xml:"location"`
}

type cppCheckResults struct {
        XMLName xml.Name         `xml:"results"`
        Errors  []*cppCheckError `xml:"errors>error"`
}

type Foo struct {
    Bar1 string `json:"bar1" yaml:"bar1yaml"`
    Bar2 string `json:"bar2"`
}

func main() {
    xmlString := `
<?xml version="1.0" encoding="UTF-8"?>
<results version="2">
    <cppcheck version="1.90"/>
    <errors>
        <error id="nullPointer" severity="error" msg="Null pointer dereference: a" verbose="Null pointer dereference: a" cwe="476">
            <location file="test/demo-null-point.cpp" line="3" column="6" info="Null pointer dereference"/>
            <location file="test/demo-null-point.cpp" line="2" column="14" info="Assignment &apos;a=NULL&apos;, assigned value is 0"/>
            <symbol>a</symbol>
        </error>
    </errors>
</results>`

    var results cppCheckResults
    xml.Unmarshal([]byte(xmlString), &results)
    for i, err := range results.Errors {
        fmt.Println(i, err.Id, err.Severity, err.Msg)
    }

    foo := Foo{
        Bar1: "1",
        Bar2: "2",
    }

    fmt.Println("json")
    if data, err := json.Marshal(foo); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(data))
    }

    fmt.Println("yaml")
    if data, err := yaml.Marshal(&foo); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(data))
    }
}

