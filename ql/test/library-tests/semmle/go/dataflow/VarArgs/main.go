package main

func source() string {
	return "untrusted data"
}

func sink(string) {
}

type A struct {
	f string
}

func functionWithStringArrayParameter(s []string) {
	t := s[0]
	sink(t) // $taintflow $f-:dataflow
}

func functionWithVarArgsOfStructsParameter(s ...A) {
	t := s[0]
	sink(t.f) // $f-:taintflow $f-:dataflow
}

func functionWithArrayOfStructsParameter(s []A) {
	t := s[0]
	sink(t.f) // $f-:taintflow $f-:dataflow
}

func main() {
	stringSlice := []string{source()}
	functionWithStringArrayParameter(stringSlice)

	a := A{f: source()}
	functionWithVarArgsOfStructsParameter(a)

	arrayOfStructs := []A{{f: source()}}
	functionWithArrayOfStructsParameter(arrayOfStructs)
}
