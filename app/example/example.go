package example

import "fmt"

var (
	varWithExplicitType string = "anystring"
	varWithImplicitType        = false
)

func privateWithoutParam() {
	fmt.Println("Functions with lower-case names are automatically private")
}

func privateWithOneParamAndOneReturn(myparameter int) string {
	var localVarWithExplicitType string = "anystring"
	fmt.Println(myparameter)
	fmt.Println(localVarWithExplicitType)
	return "private example"
}

func PublicWithOneParamAndOneReturn(myparameter bool) string {
	var localVarWithImplicitType = "anystring"
	fmt.Println(myparameter)
	fmt.Println(localVarWithImplicitType)
	return "public example"
}

func PublicMultiParamAndOneReturn(first float32, second float64) string {
	localVarShortSyntax := "anystring"
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(localVarShortSyntax)
	return "private example"
}

func PublicMultiParamWithMultiReturns(first int, second string) (string, string) {
	fmt.Println(varWithExplicitType)
	fmt.Println(varWithImplicitType)
	fmt.Println(first)
	fmt.Println(second)
	return "first return", "second return"
}
