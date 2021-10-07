package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Data struct {
	A bool
	B int16
	//C string
}

func TestUnsafe(t *testing.T)  {
	var a bool
	var b int8
	var c string
	var d map[string]string
	var e *int32
	var f = Data{}
	var g = []Data{}
	var h = [3]Data{}
	var i = [3]int64{}

	fmt.Print(reflect.TypeOf(a))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(a))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(a))

	fmt.Print(reflect.TypeOf(b))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(b))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(b))

	fmt.Print(reflect.TypeOf(c))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(c))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(c))

	fmt.Print(reflect.TypeOf(d))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(d))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(d))

	fmt.Print(reflect.TypeOf(e))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(e))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(e))

	//结构的长度必须是编译器默认的对齐长度和成员中最长类型中最小的数据大小的倍数对齐 Data由 1 + 2组成   所以
	fmt.Print(reflect.TypeOf(f))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(f))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(f))

	fmt.Print(reflect.TypeOf(g))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(g))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(g))

	fmt.Print(reflect.TypeOf(h))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(h))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(h))

	fmt.Print(reflect.TypeOf(i))
	fmt.Print(":\n  sizeof:")
	fmt.Print(unsafe.Sizeof(i))
	fmt.Print("  align:")
	fmt.Println(unsafe.Alignof(i))

	type User1 struct {
		A int32 // 4
		B []int32 // 24
		C string // 16
		D bool // 1
	}

	type sss struct {
		A int16
		B int16
		C int16
		D int32
	}

	fmt.Print(unsafe.Sizeof(sss{}))


}