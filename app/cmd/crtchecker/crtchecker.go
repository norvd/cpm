package main

import (
	"fmt"
	"log"
	"os"

	"github.com/norvd/cpm/internal/crtchecker"
)

func main() {
	f := crtchecker.New("https://crt.sh")
	certs, err := f.Find(os.Args[1])
	if err != nil {
		log.Fatalf("failed to find certs: %v", err)
	}
	fmt.Println(len(certs))
	for i, cert := range certs {
		fmt.Println(i, cert.CommonName)
	}

}
