package main

import ( "net/http" 
         "log" 
        "github.com/gorilla/mux" 
         "fmt"
         "strconv"

)

func Fibonacci(n uint) uint64 {
        if n <= 1 {
                return uint64(n)
        }

        var n2, n1 uint64 = 0, 1

        for i := uint(2); i < n; i++ {
                n2, n1 = n1, n1+n2
        }

        return n2 + n1
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Fibonacci!\n"))


for i:= uint(1); i<200;i++{

        fmt.Fprintf(w,strconv.FormatUint(Fibonacci(i),10)+"\n")

        }
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}

