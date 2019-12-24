package main

import (
    "fmt"
    "math"
    "log"
    "net/http"
    "strconv"
    "os"
    "github.com/gorilla/mux"
)


// Function to implement SieveOfEratosthenes algorithm to generate prime number.
// Arguments:
//      f -> boolean array of size of input number
//      value -> input number 
func SieveOfEratosthenes(f []bool,value int) []bool {
   // Function to determine prime numbers  
   f[0],f[1] = false,false 
   for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == true {
            for j := i * i; j <= value; j += i {
                f[j] = false
            }
        }
    }
    return f
}


// Function to check if input number is right truncated prime or not
// Arguments:
//      f -> boolean array of size of input number
//      value -> input number 
func rightPrime(f []bool, value int) bool {

    for (value!=0) {
        if (!f[value]) {
            return false
        }
        value = value/10
     }
     return true
}


// Function to check if input number is left truncated prime or not
// Arguments:
//      f -> boolean array of size of input number
//      value -> input number 
func leftPrime(f []bool, value int) bool {
    temp :=value
    cnt :=0
    for (temp!=0) {
        cnt = cnt + 1
        temp1 := temp % 10
        if (temp1 == 0) {
            return false
        }
        temp = temp/10
    }

    for i := cnt; i > 0; i-- {
        mod := int(math.Pow(10, float64(i)))
       if (!f[value % mod]) {
           return false
       }
   }
   return true

}


// Function to check if input number is Two-sided prime or not
// Arguments:
//      w -> http Response object
//      r -> http request object 
func isTwoSidedPrime(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    number := vars["number"]
    //fmt.Fprintf(w, input_number)
    input, err := strconv.Atoi(number)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }

    f := make([]bool, input+1)
    for i,_ := range f {
        f[i] =true
    }

    f = SieveOfEratosthenes(f,input)

    if (leftPrime(f, input) && rightPrime(f, input)) {
        fmt.Fprintf(w, "True")
    } else {
        fmt.Fprintf(w, "False")
    }

}



// Request handler to route Http requests
 func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/isTwoSidedPrime/{number}", isTwoSidedPrime)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}


// Main function 
func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}
