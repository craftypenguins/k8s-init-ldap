package main

import (
  "gopkg.in/ldap.v2"
  "github.com/namsral/flag"
  "fmt"
  "os"
)

func main() {

  hostPtr := flag.String("ldaphost", "localhost", "LDAP Host")
  portPtr := flag.Int("ldapport", 389, "LDAP Port")
  dnPtr := flag.String("dn", "", "Bind DN (required)")
  basednPtr := flag.String("basedn", "", "Query BaseDN (optional)")
  queryPtr := flag.String("query", "(objectClass=*)", "LDAP Query (used with basedn)")
  passwordPtr := flag.String("password", "", "Bind Password (required)")
  flag.Parse()

  l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", *hostPtr, *portPtr))
  checkErr(err)
  defer l.Close()

  err = l.Bind(*dnPtr, *passwordPtr)
  checkErr(err)

  fmt.Println("LDAP connection established")

	if *basednPtr != "" {
		searchRequest := ldap.NewSearchRequest( *basednPtr,
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
			0,
			0,
		  false,
			*queryPtr,
			[]string{"dn", "cn"},
			nil,
		)
		sr,err := l.Search(searchRequest)
		checkErr(err)
		fmt.Printf("Query of %s against %s gave %d results\n", *queryPtr, *basednPtr, len(sr.Entries))
    if len(sr.Entries) == 0 {
      fmt.Println("Exiting with error - result set empty")
      os.Exit(1)
    }
  }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
