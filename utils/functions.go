package utils
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strings"
	//"net"
    //"io"
    "os"
    "bufio"
    //"io/ioutil"
    //"os/signal"
    //"syscall"
    //"github.com/takama/daemon"
    //"time"
    //"strconv"
    //"sort"
    //"math"
//    "encoding/json"
    //"github.com/gin-gonic/gin"
    "net/http"
    //"golang.org/x/sync/errgroup"
    //https://golang.org/pkg/encoding/base64/
    //"encoding/base64"
)


const (
    StdPadding rune = '=' // Standard padding character
    NoPadding  rune = -1  // No padding
)

func HttpPost(url string, data string) {
    resp, err := http.Post(url,
        "application/x-www-form-urlencoded",
        strings.NewReader(data))

    if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
}

func Try(fun func(), handler func(interface{})) {
        defer func() {
                if err := recover(); err != nil {
                        handler(err)
                }
        }()
        fun()
}
 
func Append_file( filePath string  ,content string) {
 
    file, err := os.OpenFile(filePath, os.O_WRONLY |os.O_CREATE| os.O_APPEND, 0666)
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

func Initial(){

}
	
func Check(e error) {
    if e != nil {
        panic(e)
    }
}
  
func HttpGet(url string) {
    resp, err :=   http.Get(url)
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    /*
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
	*/
    //fmt.Println(string(body))
}
   
 func Get_database_rows(sql1 string,config string) map[int][]string {
 
	fmt.Println("Get_record_setting_mysql")
	//config: "root:@"+Getconfig("mysqlip")+"/server"
	db, err := sql.Open("mysql",config)
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()
 
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
 
	rows, err := db.Query(sql1)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	out := make(map[int][]string)
	count :=0
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		mysql_row_1 := []string {}
		for i := 0; i < len(values); i++ {
			//mysql_row_1.data[uint8(i)]  = string(values[i])
			mysql_row_1 = append(mysql_row_1,string(values[i]))
		}
		out[count] = mysql_row_1
		count ++

	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return out
 }
 