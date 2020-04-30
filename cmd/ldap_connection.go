package main

import (
	"crypto/tls"
	"os"

	auth "github.com/awltux/authboss_store_ldap"
	"github.com/go-ldap/ldap"
	log "github.com/sirupsen/logrus"
)

// init should be moved to the application init
func logInit() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func ldapConnect(ldapURL string) *ldap.Conn {
	ldapConn, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}

	// Reconnect with TLS
	err = ldapConn.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	return ldapConn
}

// DefaultLdapURL is the default url for testing.
const DefaultLdapURL = "ldap://localhost:389"

func main() {
	// Get ldapUrl from properties file
	ldapUrl := DefaultLdapURL
	logInit()
	ldapConn := ldapConnect(ldapUrl)
	defer ldapConn.Close()
	ldapStore := auth.NewServerStorerLdap(ldapConn)
	log.Info(ldapStore)
}
