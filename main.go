package testcase

import (
	"fmt"
	"log"
	"os"
)

type TestCase struct {
	in  *[]interface{}
	out *[]interface{}
}

func New() TestCase {
	return TestCase{
		in:  new([]interface{}),
		out: new([]interface{}),
	}
}

func (T *TestCase) Write(path string) {
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

	for i, input := range *T.in {
		go export(fmt.Sprintf("%s/in/input%d.txt", path, i+1), input.([]interface{}))
	}

	for i, output := range *T.out {
		go export(fmt.Sprintf("%s/out/output%d.txt", path, i+1), output.([]interface{}))
	}
}

func export(path string, data []interface{}) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("cannot create directory: %s", err.Error())
	}

	defer f.Close()

	for i := 0; i < len(data); i++ {
		_, err = f.WriteString(fmt.Sprintf("%v\n", data[i]))
		if err != nil {
			log.Panicf("error while creating %s: %s", path, err.Error())
		}
	}
}
