package main

import {
  "fmt"
  "time"
}

// global var
// var test = "blah"

func main() {
  // fmt.Println("Hello, world.")
  // customer := getData(1)
  customers := getData()
  fmt.Println(len(customers))
  fmt.Println(customers)

  // customer = getData(3)
  // fmt.Println(customer)
}

// func getData(inputs)(outputs) {
// func getData(customerId int) (customer string) {
//   // fmt.Println(test)
//   // var firstName = "Jon"
//   // lastName := "Palacio"
//   // return "Jon Palacio"
//   // return firstName + " " + lastName
//   // return fullName
//   if customerId == 1 {
//     return "Jon Palacio"
//   } else if customerId == 2 {
//     return "Bob Smith"
//   } else {
//     return ""
//   }
// }

func getData() (customers []string) {
  customers = []string{"jon Palacio", "Bob Smith", "John Smith"}

  customers = append(customers, "Molly Molly")
  customers = append(customers, "Fred Fred")
  customers = append(customers, "Harry Harry")
  customers = append(customers, "Frank Frank")
  customers = append(customers, "Sirius Sirius")
  customers = append(customers, "George George")
  customers = append(customers, "Draco Draco")

  for {
    fmt.Println("Infinite Loop 1")
    time.sleep(time.Second)
    break
  }

  return customers
}
