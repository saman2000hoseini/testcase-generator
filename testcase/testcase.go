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

func (T *TestCase) Write(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path+"/in", 0755)
	if err != nil {
		return err
	}

	for i, input := range *T.in {
		export(fmt.Sprintf("%s/in/input%d.txt", path, i+1), input.([]interface{}))
	}

	err = os.Mkdir(path+"/out", 0755)
	if err != nil {
		return err
	}

	for i, output := range *T.out {
		export(fmt.Sprintf("%s/out/output%d.txt", path, i+1), output.([]interface{}))
	}
	return nil
}

func (T TestCase) AddInput(input []interface{}) {
	*T.in = append(*T.in, input)
}

func (T TestCase) AddOutput(output []interface{}) {
	*T.out = append(*T.out, output)
}

func export(path string, data []interface{}) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("cannot create directory: %s", err.Error())
	}

	defer f.Close()

	for i := 0; i < len(data); i++ {
		_, err = f.Write([]byte(fmt.Sprintf("%v\n", data[i])))
		if err != nil {
			log.Panicf("error while creating %s: %s", path, err.Error())
		}
	}
}
