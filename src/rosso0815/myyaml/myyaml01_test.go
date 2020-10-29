package myyaml

import (
	"fmt"
	"log"
	"testing"

	"gopkg.in/yaml.v2"
)

var data = `
---
#
hosts:
    server1:
        name: 12222
firewall_network_rules:
    rule1:
        src: blablabla-host
        dst: blabla-hostname
    rule2:
        src: r2s2
        dst: r2d2
application: iama
appid: 1234
b:
    c: 2
    d: [3, 4]
    e: [6, 7]
    further:
        - name: hans
        - name: peter

`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
// type TT struct {
// 	A string
// 	B struct {
// 		RenamedC int   `yaml:"c"`
// 		D        []int `yaml:",flow"`
// 	}
// }

// type StructA struct {
// 	A string `yaml:"a"`
// }

// type StructB struct {
// 	StructA `yaml:",inline"`
// 	B       string `yaml:"b"`
// }

func TestReadWriteYaml(tt *testing.T) {
	// t := TT{}

	// err := yaml.Unmarshal([]byte(data), &t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("---1 t:\n%v\n\n", t)

	// d, err := yaml.Marshal(&t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("---2 t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("---3 m:\n%v\n\n", m)
	fmt.Println("application =", m["application"])

	// d, err = yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("---4 m dump:\n%s\n\n", string(d))
}

// func TestGenerateYaml(t *testing.T) {

// 	fmt.Println("@@@ TestGenerateYaml")
// var data = `
// a: a string from struct A
// b: a string from struct B
// `

// var b StructB

// err := yaml.Unmarshal([]byte(data), &b)
// if err != nil {
// 	fmt.Println("[ERROR ]cannot unmarshal data")
// 	os.Exit(1)
// }
// fmt.Println(b.A)
// fmt.Println(b.B)
// }
