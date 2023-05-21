package main

import (
    "strings"
    "fmt"
)

var (
    opNames = `"Skiplist","add","add","add","add","add","add","add","add","add","erase","search","add","erase","erase","erase","add","search","search","search","erase","search","add","add","add","erase","search","add","search","erase","search","search","erase","erase","add","erase","search","erase","erase","search","add","add","erase","erase","erase","add","erase","add","erase","erase","add","add","add","search","search","add","erase","search","add","add","search","add","search","erase","erase","search","search","erase","search","add","erase","search","erase","search","erase","erase","search","search","add","add","add","add","search","search","search","search","search","search","search","search","search"`
    opArgs = `[],[16],[5],[14],[13],[0],[3],[12],[9],[12],[3],[6],[7],[0],[1],[10],[5],[12],[7],[16],[7],[0],[9],[16],[3],[2],[17],[2],[17],[0],[9],[14],[1],[6],[1],[16],[9],[10],[9],[2],[3],[16],[15],[12],[7],[4],[3],[2],[1],[14],[13],[12],[3],[6],[17],[2],[3],[14],[11],[0],[13],[2],[1],[10],[17],[0],[5],[8],[9],[8],[11],[10],[11],[10],[9],[8],[15],[14],[1],[6],[17],[16],[13],[4],[5],[4],[17],[16],[7],[14],[1]`
    rets = `null,null,null,null,null,null,null,null,null,null,true,false,null,true,false,false,null,true,true,true,true,false,null,null,null,false,false,null,false,false,true,true,false,false,null,true,true,false,true,true,null,null,false,true,false,null,true,null,true,true,null,null,null,false,false,null,true,false,null,null,true,null,false,false,false,true,true,false,false,null,true,false,false,false,false,true,false,false,null,null,null,null,true,true,true,true,true,true,false,false,true`
    wants = `null,null,null,null,null,null,null,null,null,null,true,false,null,true,false,false,null,true,true,true,true,false,null,null,null,false,false,null,false,false,true,true,false,false,null,true,true,false,true,true,null,null,false,true,false,null,true,null,true,true,null,null,null,false,false,null,true,false,null,null,true,null,false,false,false,true,true,false,true,null,true,false,false,false,true,true,false,false,null,null,null,null,true,true,true,true,true,true,false,false,true`
)

type Call struct {
    opName string
    opArg string
    ret string
    want string
}

func main() {
    arrOPNames := strings.Split(strings.ReplaceAll(opNames, "\"", ""), ",")
    arrOPArgs := strings.Split(strings.ReplaceAll(strings.ReplaceAll(opArgs, "[", ""), "]", ""), ",")
    arrRets := strings.Split(rets, ",")
    arrWants := strings.Split(wants, ",")

    calls := make([]Call, len(arrOPNames))
    for i := range calls {
        calls[i].opName = arrOPNames[i]
        calls[i].opArg = arrOPArgs[i]
        calls[i].ret = arrRets[i]
        calls[i].want = arrWants[i]
    }

    for _, call := range calls {
        if call.opName == "Skiplist" {
            fmt.Println("tree := Constructor()")
        } else if call.opName == "add" {
            fmt.Printf("tree.Add(%s)\n", call.opArg)
        } else if call.opName == "search" {
            if call.ret != call.want {
                fmt.Printf("tree.Search(%s)", call.opArg)
                fmt.Printf("// want: %s", call.want)
                fmt.Println()
            }
        } else {
            fmt.Printf("tree.Erase(%s)", call.opArg)
            if call.ret != call.want {
                fmt.Printf("// want: %s", call.want)
            }
            fmt.Println()
        }
    }
}

