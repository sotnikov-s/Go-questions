package main

import (
	"io/ioutil"
	"fmt"
)

func unique(src string, dst string) error {

	text, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("error!", err)
	}

	ioutil.WriteFile(dst, text, 0644)
	return nil
}


func main() {

	err := unique("unique/src.md", "unique/out.md")
	fmt.Println(err)

}


