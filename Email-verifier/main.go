package main

import (
	"fmt"
	"net"
	"net/mail"
	"os"
	"strings"
	"net/smtp"
)
func validateSyntax(email string)bool {
	_,err := mail.ParseAddress(email)
	return err==nil
}
func validateDomain(email string)bool {
	addr,err:= mail.ParseAddress(email)
	if err !=nil{
		return false 
	}
	domain:=addr.Address
	if atIdx := len(domain) - len(email) + 1 ; atIdx>0{
		domain= domain[atIdx:]
	}else{
		return false
	}
	mxRecords,err:=net.LookupAddr(domain)
	return err==nil && len(mxRecords) >0
}
func validateSMTP(email string) bool {
    addr, _ := mail.ParseAddress(email)  // Assume already validated
    domain := addr.Address[strings.LastIndex(addr.Address, "@")+1:]

    mxRecords, _ := net.LookupMX(domain)
    if len(mxRecords) == 0 {
        return false
    }

    client, err := smtp.Dial(mxRecords[0].Host + ":25")
    if err != nil {
        return false
    }
    defer client.Close()

    if err = client.Hello("localhost"); err != nil {
        return false
    }
    if err = client.Mail("test@local.com"); err != nil {
        return false
    }
    if err = client.Rcpt(email); err != nil {
        return false
    }
    return true
}
func  main()  {
	if len(os.Args) <2{
		fmt.Println("Usage: go run main.go <email>")
		return
	}
	email:=os.Args[1]
	if validateSyntax(email) {
        if validateDomain(email) {
            fmt.Printf("%s has a valid domain.\n", email)
        } else {
            fmt.Printf("%s has an invalid or non-existent domain.\n", email)
        }
    }
}