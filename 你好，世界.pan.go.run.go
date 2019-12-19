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

var pan_2_ = [] string {"世界 你好","你在哪里？","我来了 pan 语言"}
var pan_4_  map [string  ] string

func main(){
    if len(os.Args) < 2 {
        // fmt.Println("请传入更多参数")
    }
    utils.Initial()

    fmt.Println(" 欢迎 使用 pan 语言，详细使用说明请查看：  https://github.com/youkpan/panlang")
    fmt.Println("")
    fmt.Println("方法一："  )
    for pan_5_1 := 0 ; pan_5_1 <  len(pan_2_); pan_5_1 = pan_5_1 + 1 {
        fmt.Println(pan_2_ [ pan_5_1 ] )
    }
    fmt.Println("")
    fmt.Println("方法二："  )
    for pan_5_1 ,pan_7_ := range  pan_2_ {
        fmt.Println("索引 值：",pan_5_1,"，结果：",pan_7_ )
    }
    fmt.Println("")
    fmt.Println("方法三："  )
    pan_4_ := make(map [string  ] string)
    pan_4_ ["天王星"  ] = " 比海王星近"
    pan_4_ ["地球"  ] = " 我们的家园"
    pan_4_ = pan_17_(pan_4_)
    for pan_5_1 ,pan_7_ := range  pan_4_ {
        if pan_5_1 == "地球" {
            fmt.Println("发现地球"  )
            fmt.Println("索引 值：",pan_5_1,"，结果：",pan_7_  )
        }else{
            fmt.Println("索引 值：",pan_5_1,"，结果：",pan_7_ )
        }
    }
    fmt.Println("")
    fmt.Println("读取文件")
    pan_19_("README.md")

    for {
        pan_13_ := time . Now() . Format(utils.TIME_LAYOUT)
        fmt.Println("现在是：",pan_13_)
        time. Sleep( time.Duration(10)* time.Second)
    }
}

func pan_17_ (pan_18_1 map [string  ] string)map [string  ] string{
    pan_18_1 ["火星" ] = "要去吗"
    return pan_18_1
}

func pan_19_ (pan_20_ string) string{
    pan_21_1 ,pan_22_ := os . Open(pan_20_)
    if pan_22_ != nil{
        return ""
    }
    defer pan_21_1 . Close()
    reader1 := bufio . NewReader(pan_21_1)

    pan_23_ := [] string {}

    for {
        pan_24_, pan_22_2 := reader1 . ReadString('\n')
        pan_25_string := string(pan_24_)

        fmt.Println(pan_25_string)
        if pan_22_2 != nil{
            break
        }
        pan_23_ =  append (pan_23_ ,pan_25_string)
        return pan_25_string
    }

    return ""
}
