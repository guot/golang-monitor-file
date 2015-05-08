// monitor

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	filename := flag.String("name", "README.md", "input monito file name")
	cmd := flag.String("cmd ", "./family_start.sh", "input execute  file name")
	flag.Parse()
	log.Printf("wait monti file:%s", *filename)
	log.Printf("wait cmd  :%s", *cmd)
	modi := make(chan int)
	var sleepFun = func() {
		time.Sleep(time.Duration(1) * time.Second)
	}
	go func() {
		for {

			mflag := checkFile(*filename, sleepFun)
			if mflag {
				modi <- 1
			}
		}
	}()

	for {
		data := <-modi
		log.Printf("File is modi file %#v", data)

		out, err := exec.Command("/bin/bash", *cmd).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The runner result  is: %s\n", out)

	}

}

func checkFile(fname string, fn func()) bool {
	fi1, err := os.Stat(fname)
	fn()
	fi2, _ := os.Stat(fname)
	return err == nil && !fi1.ModTime().Equal(fi2.ModTime())
}
