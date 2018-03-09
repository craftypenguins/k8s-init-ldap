package main

import (
  "gopkg.in/ldap.v2"
  "github.com/namsral/flag"
  "fmt"
)

func main() {

  hostPtr := flag.String("ldaphost", "localhost", "LDAP Host")
  portPtr := flag.Int("ldapport", 389, "LDAP Port")
  dnPtr := flag.String("dn", "", "Bind DN (required)")
  passwordPtr := flag.String("password", "", "Bind Password (required)")
  flag.Parse()

  ldap, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", *hostPtr, *portPtr))
  checkErr(err)
  defer ldap.Close()

  err = ldap.Bind(*dnPtr, *passwordPtr)
  checkErr(err)

  fmt.Println("LDAP connection established")

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
