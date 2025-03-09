package main
import "fmt"
type Rectangle struct {
    Length int
    Width  int
}
func NewRectangle() Rectangle { return Rectangle{Length: 0, Width: 0}}
func (r Rectangle) getArea() int {
    return r.Length * r.Width
}

func main() {
    obj := NewRectangle()
    fmt.Println(obj.getArea())  
}
