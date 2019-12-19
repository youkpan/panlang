package main
import (
	"fmt"
    // 	"strings"
    "os"
    "./utils"
    //     "bufio"
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
    for variable_61 := 0 ; variable_61 <  len(variable_2); variable_61 = variable_61 + 1 {
        fmt.Println(variable_2 [ variable_61 ] )
    }
    fmt.Println("")
    fmt.Println("方法2："  )
    for variable_61 ,variable_7 := range  variable_2 {
        fmt.Println("索引 值：",variable_61,",结果：",variable_7 )
    }
    fmt.Println("")
    fmt.Println("方法3："  )
    variable_4 := make(map [string  ] string)
    variable_4 ["天王星"  ] = " 比海王星近"
    variable_4 ["地球"  ] = " 我们.家园"

    for variable_61 ,variable_7 := range  variable_4 {
        if variable_61 == "地球" {
            fmt.Println("发现地球"  )
            fmt.Println("索引 值：",variable_61,",结果：",variable_7  )
        }else{
            fmt.Println("索引 值：",variable_61,",结果：",variable_7 )
        }
    }

}
