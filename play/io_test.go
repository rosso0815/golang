package play

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createTestFile(f string) {
	log.Println("@@@ createTestFile", f)
	d1 := []byte("hello andreas\n")
	err := ioutil.WriteFile(f, d1, 0644)
	check(err)
}

func TestReadFile(t *testing.T) {

	log.Println("@@@ TestReadFile")

	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	log.Print(string(dat))

	//	var v float64
	////	v = Average([]float64{1, 2})
	//	if v != 1.5 {
	//		t.Error("Expected 1.5, got ", v)
	//	}
}

func TestMain(m *testing.M) {

	log.Println("@@@ test start")

	file := "/tmp/dat"
	createTestFile(file)

	// call flag.Parse() here if TestMain uses flags
	//os.Exit(m.Run())
	m.Run()

	err := os.Remove(file)
	check(err)

	log.Println("@@@ test finished")
}
