package main

import "fmt"

type newPrintS struct{}

type basePrintS struct{}

type adapter struct {
	adapted basePrintS
}

func (a *adapter) AdaptedPrint() string {
	return a.adapted.BasePrint("base hello")
}

func (b *basePrintS) BasePrint(str string) string {
	return "base print " + str
}

func (s *newPrintS) NewPrint() string {
	return "new print"
}

func main() {
	str := &basePrintS{}
	newStr := &newPrintS{}
	adapt := &adapter{*str}
	fmt.Println(newStr.NewPrint())
	fmt.Println(adapt.AdaptedPrint())
}
