package main

type MyStruct struct {
	field string
}

func (ms MyStruct) sinkFieldNonPointerReceiver() {
	sink(ms.field) // $dataflow=20 $dataflow=24
}

func (ms *MyStruct) sinkFieldPointerReceiver() {
	sink(ms.field) // $f-:dataflow=20 $dataflow=24
}

func source() string { return "untrusted data" }

func sink(_ ...interface{}) {}

func main() {
	myStruct := MyStruct{source()}
	myStruct.sinkFieldNonPointerReceiver()
	myStruct.sinkFieldPointerReceiver()

	myStructPointer := &MyStruct{source()}
	myStructPointer.sinkFieldNonPointerReceiver()
	myStructPointer.sinkFieldPointerReceiver()
}
