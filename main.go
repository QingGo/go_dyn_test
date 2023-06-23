package main

import (
	"fmt"
	"os"
	"plugin"

	"golang.org/x/xerrors"
)

var calc func(int) int

func loadCalcPlugin(name string) (f func(int) int, err error) {
	p, err := plugin.Open(name)
	if err != nil {
		err = xerrors.Errorf("Error loading plugin: %s", err)
		return nil, err
	}

	fRaw, err := p.Lookup("Calc")
	if err != nil {
		fmt.Println("Error looking up Greet symbol:", err)
		return nil, err
	}
	f, ok := fRaw.(func(int) int)
	if !ok {
		err = xerrors.Errorf("Error asserting Calc as function")
		return nil, err
	}
	return f, nil
}

func loadCalcPlugin2(name string) (err error) {
	p, err := plugin.Open(name)
	if err != nil {
		err = xerrors.Errorf("Error loading plugin: %s", err)
		return err
	}

	fRaw, err := p.Lookup("Calc")
	if err != nil {
		fmt.Println("Error looking up Greet symbol:", err)
		return
	}
	f, ok := fRaw.(func(int) int)
	if !ok {
		err = xerrors.Errorf("Error asserting Calc as function")
		return err
	}
	calc = f
	return
}

func main() {
	// get command line args
	pluginName := os.Args[1]
	// f, err := loadCalcPlugin(pluginName)
	err := loadCalcPlugin2(pluginName)
	if err != nil {
		fmt.Printf("get error: %+v\n", err)
		return
	}
	fmt.Println(calc(1))
}
