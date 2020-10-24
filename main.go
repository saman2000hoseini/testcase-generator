package testcase

import (
	"fmt"
	"log"
	"os"
)

func Write(in, out []interface{}, path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Fatalf("cannot create directory: %s", err.Error())
	}

	err = os.Mkdir(path+"/in", 0755)
	if err != nil {
		log.Fatalf("cannot create directory: %s", err.Error())
	}

	err = os.Mkdir(path+"/out", 0755)
	if err != nil {
		log.Fatalf("cannot create directory: %s", err.Error())
	}

	for i := 0; i < len(in); i++ {
		go export(path+"/in/input"+string(i+1)+".txt", fmt.Sprintf("%v", in[i]))
	}

	for i := 0; i < len(out); i++ {
		go export(path+"/out/output"+string(i+1)+".txt", fmt.Sprintf("%v", out[i]))
	}
}

func export(path string, data string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("cannot create directory: %s", err.Error())
	}

	defer f.Close()

	_, err = f.Write([]byte(data))
	if err != nil {
		log.Panicf("error while creating %s: %s", path, err.Error())
	}
}
