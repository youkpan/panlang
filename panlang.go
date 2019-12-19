package main
import (
//	"database/sql"
	_ "github.com/go-sql-driver/mysql"
//	"log"
	"fmt"
	"strings"
	//"net"
    "os"
    "bufio"
 //   "io/ioutil"
    //"os/signal"
    //"syscall"
    //"github.com/takama/daemon"
 //   "time"
 //   "strconv"
    "sort"
    //"math"
//    "encoding/json"
 //   "github.com/gin-gonic/gin"
  //  "net/http"
 //   "golang.org/x/sync/errgroup"
    //https://golang.org/pkg/encoding/base64/
  //  "encoding/base64"
)

var variables_replace map[string]string
var reserved_word_order []string
var reserved_word map[string]string

func main(){
    if len(os.Args)<2{
        fmt.Println("请把 你编写的 .pan.go 程序文件拖进来")
        return
    }
    
    filename := os.Args[1]

    fmt.Println("文件名：",filename)

    variables_replace = make(map[string]string)
    reserved_word = make(map[string]string)    
    init_reserved_word()
    gen_reserved_word_order()


    file,err := os.Open(filename)
    if err!=nil{
        return
    }
    reader := bufio.NewReader(file)
    //scanner := bufio.NewScanner(file)
    file_out_content:=""
    lines :=[]string{}
    //for scanner.Scan() {
    for {
        //line, prefix, err := reader.ReadLine();
        //line := scanner.Text()
        line, err := reader.ReadString('\n')
        line1 := string(line)
        line1 = strings.ReplaceAll(line1,"\r\n","")
        line1 = strings.ReplaceAll(line1,"\n","")
        lines = append(lines,line1)
        //fmt.Println(string(line))
        if err != nil  {
            break
        }
     }

     find_variables_replace(lines)

     for i := 0; i < len(lines); i++ {
        line2 := replace_key_word(lines[i])
        file_out_content += line2+"\r\n"
     }

     //path := strings.ReplaceAll(filename,".pan.go",".go")
     path := filename+".run.go"
     Write_file(path ,file_out_content)
}

var variables_index = 0
func remove_reserved_word_array(lines []string)[]string{
    for i := 0; i < len(lines); i++ {
        for index := 0; index < len(reserved_word_order); index++ {
            reserved_word := reserved_word_order[index]
            if len(reserved_word)>0{
                lines[i] = strings.ReplaceAll(lines[i]+" ",reserved_word,"")
                lines[i] = strings.ReplaceAll(" "+lines[i],reserved_word,"")
                lines[i] = strings.ReplaceAll(lines[i],reserved_word,"")
                lines[i] = strings.ReplaceAll(" ",reserved_word,"")
            }
        }
    }

    return lines

}

func remove_reserved_word(line string,rep map[string]string)string{

        for reserved_word,_ := range rep {
            if len(reserved_word)>0{
                line = strings.ReplaceAll(line,reserved_word,"")
            }
        }
    return line
}

func remove_strings(line string) string{
    string_start := -1
    string_end := -1
    for strings.Contains(line,"“"){
        string_start =strings.Index(line,"“")
        if string_start >0{
            string_end =strings.Index(line,"”")
            //lineslice := strings.Split(line,"“")
            //lineslice2 := strings.Split(line,"“")
            //remove string
            if string_end >0{
                line = line[0:string_start] + line[string_end+1:len(line)]
            }else{
                break
            }
        }        
    }


    for strings.Contains(line,"\""){
        string_start =strings.Index(line,"\"")
        if string_start >0{

            string_end =strings.Index(line[string_start+1:],"\"")+string_start+1
            if string_end >0{
                line = line[0:string_start] + line[string_end+1:len(line)]
            }else{
                break
            }
        }        
    }

    return line
}

func find_variables_replace(lines [] string){
    //string_start := -1
    temp_reserved_word := make(map[string]string)
    comment_stat :=0

    for k,v := range reserved_word {
        temp_reserved_word[k] = ""
        temp_reserved_word[v] = ""
    }
    
    for i := 0; i < len(lines); i++ {
        line := lines[i]
        fmt.Println("第 ",i ,"行：",line)

        line = remove_strings(line)
        
        if strings.Contains(line,"注释开始") {
            fmt.Println("注释开始")
            comment_stat = 1
            continue
        }

        if strings.Contains(line,"注释结束") && comment_stat == 1 {
            fmt.Println("注释结束")
            comment_stat = 0
        }else if strings.Contains(line,"注释")  {
            continue
        }

        if comment_stat==1{
            continue
        }

        line = remove_reserved_word(line,temp_reserved_word)
        lineslice := strings.Split(line," ")
        
        for j := 0; j < len(lineslice); j++ {
            variable_name := lineslice[j]
            variable_name = strings.ReplaceAll(variable_name,"�","")
            variable_name = strings.ReplaceAll(variable_name," ","")
            variable_name = strings.ReplaceAll(variable_name,"\t","")
            if len(variable_name)>0 {
                _,isset := variables_replace[variable_name]
                if isset{
                    continue
                }
                _,isset2 := reserved_word[variable_name]
                if isset2{
                    continue
                }
                variables_index_s:=fmt.Sprintf("%d",variables_index)
                variables_replace[variable_name] = "variable_"+variables_index_s
                variables_index ++
                fmt.Println(" 变量名:[",variable_name,"] ,配置为 :","pan_variable_"+variables_index_s)
                //fmt.Println(" find variable name:",variable_name," ,config to :","variable_"+variables_index_s)
            }
        }
 
     }
}

func replace_with_reserved(line string)string{

    for _, key := range reserved_word_order{
        line = strings.ReplaceAll(line,key,reserved_word[key])
    }
    return line
}

func replace_with_wordmap(line string,wordmap map[string]string)string{

    for old,new := range wordmap{
        line = strings.ReplaceAll(line,old,new)
    }
    return line
}

func replace_key_word(line string)string{
    line = replace_with_reserved(line)

    line = replace_with_wordmap(line,variables_replace)
    return line
}



func Write_file(filePath string ,content string){

    file, err := os.OpenFile(filePath, os.O_WRONLY |os.O_TRUNC|os.O_CREATE, 0666)
    if err != nil{
        fmt.Println("open file err",err)
        return
    }
 
    //及时关闭file句柄
    defer file.Close()
    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(file)
    
    write.WriteString(content)
    
    //Flush将缓存的文件真正写入到文件中
    write.Flush()
}

func set_reserved_word(in string,out string){
    reserved_word[in] = out
    
}

func init_reserved_word(){
    set_reserved_word("变量","var")
    set_reserved_word("包名","package")
    set_reserved_word("主程序","main")
    set_reserved_word("导入包","import")
    set_reserved_word("字典","map")
    set_reserved_word("数组","[]")
    set_reserved_word("生成","make")
    set_reserved_word("循环","for")
    set_reserved_word("启动循环：","for")
    set_reserved_word("若","if")
    set_reserved_word("如果","if")
    set_reserved_word("函数","func")
    set_reserved_word("功能","func")
    set_reserved_word("推迟执行","defer")
    set_reserved_word("跳出循环","break")
    set_reserved_word("默认","default")
    set_reserved_word("选择","select")
    set_reserved_word("此外","else")
    set_reserved_word("常数","const")
    set_reserved_word("往下执行","fallthrough")
    set_reserved_word("继续","continue")
    set_reserved_word("返回","return")
    set_reserved_word("选择执行","switch")
    set_reserved_word("当","case")
    set_reserved_word("当它为","case")

    set_reserved_word("等待队列执行完毕","<-")
    set_reserved_word("队列","chan")
    set_reserved_word("运行线程","go")
    set_reserved_word("长度","len")
    set_reserved_word("添加","append")
    set_reserved_word("关闭","close")

    set_reserved_word("生成范围","range")
    //package
    set_reserved_word("运行库","runtime")
    set_reserved_word("格式","fmt")
    set_reserved_word("打印","Println")
    set_reserved_word("系统","os")
    set_reserved_word("传入参数","Args")
    set_reserved_word("工具集","utils")
    set_reserved_word("初始化函数","Initial")

    set_reserved_word("显示数","%d")
    set_reserved_word("显示字符串","%s")
    set_reserved_word("显示浮点","%f")

    set_reserved_word("显示","Println")
    set_reserved_word("生成","make")

    set_reserved_word("返回","return")
    set_reserved_word("定义","type")
    set_reserved_word("结构体","struct")
    set_reserved_word("定义","type")
    set_reserved_word("接口","interface")
    set_reserved_word("的类型是","")
 
    set_reserved_word("为","=")
    set_reserved_word("设置","=")
    set_reserved_word("等于","=")
    set_reserved_word("取引用","&")
    set_reserved_word("引用","*")
    set_reserved_word("与运算","&")
    set_reserved_word("或运算","|")

    set_reserved_word("或者","||")
    set_reserved_word("与","|")

    set_reserved_word("不相等","!=")
    set_reserved_word("相等于","==")
    set_reserved_word("相等","==")
    set_reserved_word("不","!")
    set_reserved_word("的",".")
    //set_reserved_word("。",".")
    set_reserved_word("，",",")
    set_reserved_word("模块调用",".")
    
    set_reserved_word("注释","//")
    set_reserved_word("注释开始","/*")
    set_reserved_word("注释结束","*/")
    set_reserved_word("使用索引：","[")
    set_reserved_word("结束索引","]")

    set_reserved_word("初始化",":=")
    set_reserved_word("初始化为",":=")
    set_reserved_word("删除","delete")
    set_reserved_word("字节","byte")
    set_reserved_word("语句","string")
    set_reserved_word("字符串","string")
    set_reserved_word("整数","int")
    set_reserved_word("长整数","int64")
    set_reserved_word("浮点数","float64")
    set_reserved_word("布尔","bool")
    set_reserved_word("“","\"")
    set_reserved_word("”","\"")
    set_reserved_word("右移",">>")
    set_reserved_word("左移","<<")
    set_reserved_word("》",">")
    set_reserved_word("《","<")
    set_reserved_word("大于",">")
    set_reserved_word("小于","<")
    set_reserved_word("（","(")
    set_reserved_word("（","(")
    set_reserved_word("）",")")
    set_reserved_word("『","{")
    set_reserved_word("』","}")
    set_reserved_word("；",";")
    set_reserved_word("加","+")
    set_reserved_word("减","-")
    set_reserved_word("乘","*")
    set_reserved_word("除","/")

    set_reserved_word("一","1")
    set_reserved_word("二","2")
    set_reserved_word("三","3")
    set_reserved_word("四","4")
    set_reserved_word("五","5")
    set_reserved_word("六","6")
    set_reserved_word("七","7")
    set_reserved_word("八","8")
    set_reserved_word("九","9")
    set_reserved_word("零","0")


}
 
type string_array [] string

func (s string_array) Len() int { return len(s) }
func (s string_array) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s string_array) Less(i, j int) bool { return len(s[i]) > len(s[j]) }

func gen_reserved_word_order(){
    var word_order string_array

    for k,_ := range reserved_word{
        word_order = append(word_order,k)
    }
    sort.Stable(word_order)
    reserved_word_order = word_order
}

  