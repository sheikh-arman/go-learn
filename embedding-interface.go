package main

import "io"
import "os"
import "fmt"

type echoer struct{
    io.Reader //Embed an interface
}

func (e * echoer) Read(p []byte) (int, error) {
    amount, err := e.Reader.Read(p)
    if err == nil {
        fmt.Printf("Overridden read %d bytes:%s\n",amount,p[:amount])
    }
    return amount,err
}

func readUpToN(r io.Reader, n int) {
    buf := make([]byte,n)
    amount, err := r.Read(buf[:])
    if err == nil {
        fmt.Printf("Read %d bytes:%s\n",amount,buf[:amount])
    }
}

func main(){
    readUpToN(os.Stdin,3)

    var replacement echoer
    replacement.Reader = os.Stdin

    readUpToN(&replacement,3)
}