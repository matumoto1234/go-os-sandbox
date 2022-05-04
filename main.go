package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func git(arg ...string) error {
	err := exec.Command("git", arg...).Run()
	if err != nil {
		return errors.New(fmt.Sprintf("git: run the command [git %v]\n", arg))
	}
	return nil
}

func main() {
	err := os.MkdirAll("hoge/fuga/piyo", 0755)
	if err != nil {
		fmt.Println(err)
	}

	content := "hogehoge\nfugafuga\npiyopiyooooo!!!!"
	ioutil.WriteFile("hoge/a.txt", []byte(content), 0755)

	ioutil.WriteFile("hoge/fuga/piyo/b.txt", []byte(content), 0755)

	err = git("add", ".")
	if err != nil {
		fmt.Println(err)
	}
}
