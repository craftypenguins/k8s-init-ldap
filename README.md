# Kubernetes Init Container Helper - LDAP Probe

This is a simple GO application that can be statically compiled and build into
a Scratch docker container and used as a K8S init container to check that
a LDAP service is alive and usable with provided credentials and optionally
running a SQL Query that will result in a non-empty response

## Getting Started

Clone down and run 
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o k8s-init-ldap .
```

### Prerequisites

```
go get gopkg.in/ldap.v2
go get github.com/namsral/flag
```

## Usage ##

### Command line arguments ###
```
Usage of ./k8s-init-ldap:
  -dn string
    	Bind DN (required)
  -ldaphost string
    	LDAP Host (default "localhost")
  -ldapport int
    	LDAP Port (default 389)
  -password string
    	Bind Password (required)
```

### Environment Variables ###

DN
LDAPHOST
LDAPPORT
PASSWORD


### Example Usage ###

```
docker run -t -e DN="cn=read-only-admin,dc=example,dc=com" -e PASSWORD=password \
            -e LDAPHOST=ldap.example.com -e LDAPPORT=389 \
            craftypenguins/k8s-init-ldap:latest 
```

## Authors

* **Richard Clark** - *Initial work* - [kti-richard](https://github.com/kti-richard)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

