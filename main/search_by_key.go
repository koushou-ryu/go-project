package main

import (
   "fmt"
   "io/ioutil"
   "os"
   "path/filepath"
)

func dirwalk(dir string) []string {
   files, err := ioutil.ReadDir(dir)
   if err != nil {
       panic(err)
   }
   var paths []string
   for _, file := range files {
       if file.IsDir() {
           paths =  dirwalk(filepath.Join(dir, file.Name()))
           continue
       }
       paths = append(paths, filepath.Join(dir, file.Name()))
   }
   return paths
}

func main() {
   var files = dirwalk("./resouces")

    for i,v := range files {
       fmt.Println(i,v)
       readFile(v)
    }
}

func readFile(path string) {
    fmt.Println("ファイル読み取り処理を開始します")
    // ファイルをOpenする
    f, err := os.Open(path)
    if err != nil{
        fmt.Println("error")
    }
    defer f.Close()

    // 一気に全部読み取り
    b, err := ioutil.ReadAll(f)
    // 出力
    fmt.Println(string(b))

}