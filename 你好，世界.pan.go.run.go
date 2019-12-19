package main
import (
	"fmt"
    // 	"strings"
    "os"
    "time"
    "./utils"
    "bufio"
    //     "sort"
    
)
// 这是1个示例程序
/*
    欢迎使用 pan 语言,详细使用说明请查看：
    https://github.com/youkpan/panlang

*/

var variable_2 = [] string {"世界 你好","你在哪里？","我来了 pan 语言"}
var variable_4  map [string  ] string

func main(){
    if len(os.Args) < 2 {
        // fmt.Println("请传入更多参数")
    }
    utils.Initial()

    fmt.Println(" 欢迎 使用 pan 语言,详细使用说明请查看：  https://github.com/youkpan/panlang")
    fmt.Println("")
    fmt.Println("方法1："  )
    for variable_51 := 0 ; variable_51 <  len(variable_2); variable_51 = variable_51 + 1 {
        fmt.Println(variable_2 [ variable_51 ] )
    }
    fmt.Println("")
    fmt.Println("方法2："  )
    for variable_51 ,variable_6 := range  variable_2 {
        fmt.Println("索引 值：",variable_51,",结果：",variable_6 )
    }
    fmt.Println("")
    fmt.Println("方法3："  )
    variable_4 := make(map [string  ] string)
    variable_4 ["天王星"  ] = " 比海王星近"
    variable_4 ["地球"  ] = " 我们.家园"
    variable_4 = variable_13(variable_4)
    for variable_51 ,variable_6 := range  variable_4 {
        if variable_51 == "地球" {
            fmt.Println("发现地球"  )
            fmt.Println("索引 值：",variable_51,",结果：",variable_6  )
        }else{
            fmt.Println("索引 值：",variable_51,",结果：",variable_6 )
        }
    }
    fmt.Println("")
    fmt.Println("variable_15")
    variable_15("README.md")

    for {
        variable_11 := time . Now() . Format(utils.TIME_LAYOUT)
        fmt.Println("现在是：",variable_11)
        time. Sleep( time.Duration(10)* time.Second)
    }
}

func variable_13 (variable_141 map [string  ] string)map [string  ] string{
    variable_141 ["火星" ] = "要去吗"
    return variable_141
}

func variable_15 (variable_16 string) string{
    variable_191,variable_18 := os . Open(variable_16)
    if variable_18 != nil{
        return ""
    }

    reader1 := bufio . NewReader(variable_191)

    variable_20 := [] string {}

    for {
        variable_22, variable_182 := reader1 . ReadString('\n')
        variable_24string := string(variable_22)

        fmt.Println(variable_24string)
        if variable_182 != nil{
            break
        }
        variable_20 =  append (variable_20 ,variable_24string)
        return variable_24string
    }

    return ""
}
