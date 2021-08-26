package main

// reads from a file that has the hex output from showing certs in a table (gets output as a hexidecimal.)
import (
	"bytes"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func isHexChar(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	default:
		return false
	}
}
func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	var b []byte
	var err error

	b, err = ioutil.ReadFile("/Users/mkeffele/test.crt")

	//fmt.Println(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	var bb bytes.Buffer
	// Remove 0x prefix
	if len(b) >= 2 && bytes.EqualFold(b[:2], []byte("0x")) {
		b = b[2:]
	}
	for _, c := range b {
		if isHexChar(c) {
			bb.WriteByte(c)
		}
	}
	out := make([]byte, bb.Len()/2)
	if _, err = hex.Decode(out, bb.Bytes()); err != nil {
		fatal(err)
	}
	//fmt.Print(hex.Dump(out))
	//fmt.Println(b)
	//content2, err := hex.DecodeString(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	fmt.Println(content2)
	//	finalContent := []byte(content2)
	cert, err := x509.ParseCertificate(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cert.NotBefore)
	fmt.Println(cert.NotAfter)
	fmt.Println(cert.Subject)
	fmt.Println(cert.Subject.Names)
	fmt.Println(cert.Subject.Names[0])
	fmt.Println(cert.Subject.Names[0].Type)
	fmt.Println(cert.Subject.Names[0].Value)
	fmt.Println(cert.Issuer)
	fmt.Println(cert.Issuer.Organization)
	fmt.Println(cert.SerialNumber)
	return
}
