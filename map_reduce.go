// map_reduce
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type channel struct {
	data chan string
	ctrl chan string
}

type data struct {
	key string
	val string
}

func list_files(dir, suffix string) (files []string, err error) {
	files = make([]string, 0, 100)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func new_channel(a *channel) channel {
	a.ctrl = make(chan string)
	a.data = make(chan string)
	return a
}

func file_parse(file_name string, chs chan string) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		buf, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			chs <- buf
		}
	}
	return
}

func map_job(chs_src chan string, chs_dest data) {
	select {
	case temp_data := <-chs_src:
		{
			//TODO: parse the string
			fmt.Printf("TODO: parse the string")
			return
		}
	default:
	}
	return
}

func reduce_job(chs []chan int, chan_ctl chan string) {
	for {
		select {
		case v := <-chs[0]:
			if v == 100 {
				chan_ctl <- "reduce finished"
				return
			} else {
				fmt.Printf("v is %d in reduce\n", v)
			}
		default:
			continue
		}
	}
}

func main() {
	list_files("C:\\Go\\liteide\\share\\liteide\\liteenv", ".env")
	num_map_job := 20
	num_reduce_job := 20
	var chs channel
	new_channel(chs)
	//data := rand.Perm(10)
	go map_job(chs_data, data)
	//for i := 0; i < 2; i++ {
	go reduce_job(chs_data, chs_ctl)
	for {
		sig := <-chs_ctl
		if sig == "reduce finished" {
			fmt.Printf("It is over")
			break
		}
	}
	//}
}
