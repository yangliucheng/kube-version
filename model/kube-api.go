package model

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"io"
	// "io/ioutil"
)

type APi struct {


}

func Config(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("open file fail:",err)
	}
	reader := bufio.NewReader(file)
	for {
		line,_,err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read line fail:",err)
		}
		strA := strings.Split(string(line), " ")
		
		fmt.Println(strA[0],"--",strA[1])
	}
}