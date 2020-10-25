package main

import (
   "fmt"
   "io/ioutil"
   "os"
   "path/filepath"
   "bufio"
   "strings"
)

// マイン関数
func main() {
   var files = getFileList("./resouces")
   
   fmt.Print("検索キーワードを入力してください：")
   var key = StrStdin()

   for _,v := range files {
      //fmt.Println(i,v)
      readFile(key , v)     
   }
}


// ファイルリスト取得関数
func getFileList(dir string) []string {
   files, err := ioutil.ReadDir(dir)
   if err != nil {
       panic(err)
   }
   var paths []string
   for _, file := range files {
       if file.IsDir() {
           paths =  getFileList(filepath.Join(dir, file.Name()))
           continue
       }
       paths = append(paths, filepath.Join(dir, file.Name()))
   }
   return paths
}

//ファイルを読み込み関数
func readFile(key string, path string) {

    // ファイルをOpenする
    f, err := os.Open(path)
    if err != nil{
        fmt.Println("error")
    }
    defer f.Close()

    // 一気に全部読み取り
    b, err := ioutil.ReadAll(f)
    // 出力
    fmt.Print("検索対象ファイル:" + path)
    fmt.Print("    検索結果件数：")
    fmt.Println(strings.Count(string(b),key))

}

// 検索キーワード入力関数
func StrStdin() (stringInput string) {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    stringInput = scanner.Text()

    stringInput = strings.TrimSpace(stringInput)
    return
}