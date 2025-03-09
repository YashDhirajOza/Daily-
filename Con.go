package main
import "fmt"
type Car struct{
  Brand string 
  Model string 
  Year int }

fucn NewCar(brand,model1 string, year int) *Car{
  return &Car{Branf:brand,Model:model, Year:year}

  func main(){
    car1:= NewCar("Toyota","Coralla",2022)
    fmt.Println(car1)
    }
