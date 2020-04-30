package auth

import (
	"context"
	"fmt"

	"github.com/go-ldap/ldap"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/authboss"
)

// ServerStorerLdap contains state for the ldap connection
type ServerStorerLdap struct {
	conn *ldap.Conn
}

// NewServerStorerLdap will create a new ServerStorer
func NewServerStorerLdap(conn *ldap.Conn) *ServerStorerLdap {
	return &ServerStorerLdap{
		conn: conn,
	}
}

// Save the user to LDAP DB
func (ssldap ServerStorerLdap) Save(_ context.Context, user authboss.User) error {
	// Assert authboss.User is a pointer to an LDAP transfer User
	//u := user.(*User)
	//ssldap.Users[u.Email] = *u

	log.Debug(fmt.Printf("Saved user: %s", user.GetPID))
	return nil
}

// Load the user
func (ssldap ServerStorerLdap) Load(_ context.Context, key string) (user authboss.User, err error) {
	// Get the user from the DB using the key
	//u, ok := ssldap.Users[key]
	//if !ok {
	//	return nil, authboss.ErrUserNotFound
	//}

	log.Debug(fmt.Printf("Loaded user: %s", key))
	return nil, nil
}
