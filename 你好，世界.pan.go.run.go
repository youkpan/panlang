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

var variable_1 = [] string {"世界 你好","你在哪里？","我来了 pan 语言"}
var variable_3  map [string  ] string

func main(){
    if len(os.Args) < 2 {
        // fmt.Println("请传入更多参数")
    }
    utils.Initial()

    fmt.Println(" 欢迎 使用 pan 语言,详细使用说明请查看：  https://github.com/youkpan/panlang")
    fmt.Println("")
    fmt.Println("方法1："  )
    for variable_61 := 0 ; variable_61 <  len(variable_1); variable_61 = variable_61 + 1 {
        fmt.Println(variable_1 [ variable_61 ] )
    }
    fmt.Println("")
    fmt.Println("方法2："  )
    for variable_61 ,variable_7 := range  variable_1 {
        fmt.Println("索引 值：",variable_61,",结果：",variable_7 )
    }
    fmt.Println("")
    fmt.Println("方法3："  )
    variable_3 := make(map [string  ] string)
    variable_3 ["天王星"  ] = " 比海王星近"
    variable_3 ["地球"  ] = " 我们.家园"
    variable_3 = 修改variable_18(variable_3)
    for variable_61 ,variable_7 := range  variable_3 {
        if variable_61 == "地球" {
            fmt.Println("发现地球"  )
            fmt.Println("索引 值：",variable_61,",结果：",variable_7  )
        }else{
            fmt.Println("索引 值：",variable_61,",结果：",variable_7 )
        }
    }
    fmt.Println("")
    fmt.Println("读取文件")
    读取文件("README.md")

    for {
        variable_13 := time . Now() . Format(utils.TIME_LAYOUT)
        fmt.Println("现在是：",variable_13)
        time. Sleep( time.Duration(10)* time.Second)
    }
}

func 修改variable_18(variable_181 map [string  ] string)map [string  ] string{
    variable_181 ["火星" ] = "要去吗"
    return variable_181
}

func 读取文件(variable_21 string) string{
    文件1,variable_22 := os . Open(variable_21)
    if variable_22 != nil{
        return ""
    }

    reader1 := bufio . NewReader(文件1)

    variable_24 := [] string {}

    for {
        variable_25, variable_222 := reader1 . ReadString('\n')
        variable_27string := string(variable_25)

        fmt.Println(variable_27string)
        if variable_222 != nil{
            break
        }
        variable_24 =  append (variable_24 ,variable_27string)
        return variable_27string
    }

    return ""
}
