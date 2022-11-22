package pkg2

import (
	pkg1 "github.com/gucio321/pkg2/sample-module1"
)

func Bar() {
	pkg1.Foo()
}
