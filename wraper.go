package pkg2

// #include "sample-module1/file.c"
import "C"

func Bar() {
	C.main()
}
