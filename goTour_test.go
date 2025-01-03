package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/wc"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	tt := time.Now()
	switch {
	case tt.Hour() < 12:
		fmt.Println("早上好！")
	case tt.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}

func TestSwitch1(t *testing.T) {
	defer fmt.Println("world")
	fmt.Println("hello")
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TestReflect(t *testing.T) {
	for v, i := range pow {
		fmt.Println(v)
		fmt.Println(i)
	}
}
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := range img {
		img[y] = make([]uint8, dx)
		for x := range img[y] {
			img[y][x] = uint8(x % (y + 1))
		}
	}
	return img
}
func TestPic(t *testing.T) {
	pic.Show(Pic)
	//base64Str := "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w4oA3rz9NvOKAN688w4jiwDevPsuU4sA3rz3HoOLAN68/z6ziwDefPAB44sA3nz4IROMAN589BFDjADefPwxc4wA3nzyCaOMAN58+inTjADefPYZA40A3nz+OTONAN588QVjjQDefPklk40A3nz1FcONAN58/TXzjQDefPMNI44A3nz7LVOOAN589x2DjgDefP89s44A3vzwA+OOAN78+CMTjwDe/PQTQ48A3vz8M3OPAN788gujjwDe/Por048A3vz2GwIggDe//44A4M0ffyAAeLN4MQKAN0uWIAB4s3QpAoA3y5YhAHizfDkCgDcrViAAeLNyJQKAN6tWIQB4s3o1AoA3a9YgAHizdi0CgDfr1iEAeLN+PQKANxs2IAB4s3EjAoA3mzYhAHizeTMCgDdbtiAAeLN1KwKAN9u2IQB4s307AoA3O3YgAHizcycCgDcVFQgA3lRWIgB4U1WFAOBNdTUCgDc1NQgA3tTWIgB4U1eHAOBNfT0CgDcNDQgA3jQ2IgB409SEAOBNczMCgDctLQgA3rS2IgB409aGAOBNezsCgDcdHQgA3nR2IgB409WFAODNrl0IAN7s3o0A4M2ePQgA3uzdiwDgzb59CADe7N+PAODNgQMIAN4cPIgA4M2hQwgA3hw+jADgzZEjCADeHD2KAODNsWMIAN4cP44A4M2JEwgA3pw8iQDgzalTCADenD6NAODNmTMIAN6cPYsA4M25cwgA3pw/jwDgzYULCADeXLyIAODNpUsIAN5cvowA4M2VKwgA3ly9igDgzbVrCADeXL+OAODNjRsIAN7cvIkA4M2tWwgA3ty+jQDgzZ07CADe3L2LAODNvXsIAN7cv48A4M2DBwgA3vz5JwKAN93dCADe9PTo/wEAAP//EjOHlO2qZmIAAAAASUVORK5CYII="
	//decodeData, err := base64.StdEncoding.DecodeString(base64Str)
	//if err != nil {
	//	fmt.Println("fail to decode")
	//}
	//filename := `C:\Users\mamingyang\Desktop\md\output.png`
	//err = ioutil.WriteFile(filename, decodeData, 0644)
	//if err != nil {
	//	fmt.Println("fail to save")
	//}
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func TestMap(t *testing.T) {
	m = make(map[string]Vertex) // m初始值是nil并没有分配内存 使用make初始化分配内存
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

var mapTest map[string]int

func WordCount(s string) map[string]int {
	split := strings.Split(s, " ")
	mapTest = make(map[string]int) //初始化map
	for _, v := range split {
		mapTest[v]++
	}
	return mapTest
}

func TestWordCount(t *testing.T) {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func TestFunction(t *testing.T) {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunc(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		j := pos(i)
		k := neg(-2 * i)
		fmt.Println(j, k)
	}
}

func foo(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	//for _, val := range values {
	//	fs = append(fs, func() {
	//		fmt.Printf("foo7 val = %d\n", x+val)
	//	})
	//}
	var val int
	for i := 0; i < len(values); i++ {
		val = values[i]
		fs = append(fs, func() {
			fmt.Printf("foo8 val = %d\n", x+val)
		})
	}
	return fs
}
func TestClosure(t *testing.T) {
	for _, f7 := range foo(11) {
		f7()
	}
}

// fibonacci 是返回一个「返回一个 int 的函数」的函数
func fibonacci() func() int {
	var callTimes, a, b = 0, 0, 1
	return func() int {
		if callTimes == 0 || callTimes == 1 {
			callTimes++
			return callTimes - 1
		} else {
			res := a + b
			a, b = b, res
			return res
		}
	}
}

func TestFibo(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

type Entity struct {
	X, Y int
}
type A interface {
	calculate(entity Entity)
}
type B int

func (b B) calculate(entity Entity) {
	res := entity.X + entity.Y + int(b)
	fmt.Println(res)
}

type C int

func (c C) calculate(entity Entity) {
	var res int
	res = entity.X * entity.Y * int(c)
	fmt.Println(res)
}

func TestMethod(t *testing.T) {

	e := Entity{5, 9}
	ab := A(B(7))
	ab.calculate(e)
	var b A
	b = B(3)
	b.calculate(e)
	b = C(2)
	b.calculate(e)

	var c interface{} = 23
	fmt.Printf("%T, %v\n", c, c)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func TestNilInterface(tt *testing.T) {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Z interface{}

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func TestNoMethodInterface(tt *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		sqrt := math.Sqrt(x)
		return sqrt, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}
func TestError(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func TestReader(t *testing.T) {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

// TODO: 为 MyReader 添加一个 Read([]byte) (int, error) 方法。
func (mr MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func TestReaderA(t *testing.T) {
	reader.Validate(MyReader{})
}

// ---------------------------------------------------------------------

// rot13代换
func rot13(out byte) byte { //字母转换
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}
	return out
}

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b) // 把rr里面的io.Reader读取的内容写入b
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func TestRot13Reader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type Person struct {
	Name string
	Age  int
}

func (pp *Person) setName(p Person, name string) {
	p.Name = name
	fmt.Printf("%p\n", &p)
}

func setNameP(p *Person, name string) {
	p.Name = name
}
func TestPPP(t *testing.T) {
	p := Person{Name: "Mike", Age: 20}
	fmt.Printf("%p\n", &p)
	(p).setName(p, "Jack")
	fmt.Println(p)
	setNameP(&p, "Marry")
	fmt.Println(p)
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	m.Bounds()
}

type Image struct {
	w int
	h int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.w, m.h)
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x), G: uint8(y), B: uint8(255), A: uint8(255)}
}

func TestImage(t *testing.T) {
	m := Image{w: 200, h: 200}
	pic.ShowImage(m)
}

// Index 返回 x 在 s 中的下标，未找到则返回 -1。
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 的类型为 T，它拥有 comparable 可比较的约束，
		// 因此我们可以使用 ==。
		if v == x {
			return i
		}
	}
	return -1
}

func TestGenerics(t *testing.T) {
	// Index 可以在整数切片上使用
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index 也可以在字符串切片上使用
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

// List 表示一个可以保存任何类型的值的单链表。
type List[T any] struct {
	next *List[T]
	val  T
}

func TestGenerics1(t *testing.T) {
	l := List[string]{
		next: &List[string]{
			next: nil,
			val:  "saaass",
		},
		val: "sss",
	}
	next := l.next
	fmt.Println(next.val)
}
