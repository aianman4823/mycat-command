package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 終了コード
const (
	ExitCodeOk = iota
	ExitCodeError
)

func readFile(fn string, num int, n *string) error {
	path := filepath.Join("/Users/harada_akito/Documents/golang/mycat", fn)
	// fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// txtファイルを探し出す
	researchErr := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".txt" {
				// fmt.Println(path)
				// 標準入力から読み込む
				scanner := bufio.NewScanner(f)
				// fmt.Println(os.Stdin)
				var i int = 1
				for scanner.Scan() {
					// 1行分を出力する
					// ここに-nオプションを付けたら行番号を各行につける処理を追加
					if *n == "order" {
						if num == 0 {
							fmt.Printf("%d: %s\n", i, scanner.Text())
						} else if num == 1 {
							fmt.Printf("%d: %s\n", i+(num*2), scanner.Text())
						}
						i++
					} else {
						fmt.Println(scanner.Text())
					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
				}
			}
			return nil
		})
	if researchErr != nil {
		return researchErr
	}
	return nil
}

func main() {
	// 左から順位オプション名、デフォルトの値、helpテキストが引数に入る
	n := flag.String("n", "", "String help message")
	// コマンドラインの引数を受け取る処理
	flag.Parse()
	// ファイルネームを受け取る
	fileNames := flag.Args()
	for i, fn := range fileNames {
		err := readFile(fn, i, n)
		if err != nil {
			log.Fatal(err)
			os.Exit(ExitCodeError)
		}
	}
	os.Exit(ExitCodeOk)
}
