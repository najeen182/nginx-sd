package pkg

import (
	"crypto"
	"fmt"
	"os"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type User struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
}

func (u *User) GetEmail() string                        { return u.Email }
func (u *User) GetRegistration() *registration.Resource { return u.Registration }
func (u *User) GetPrivateKey() crypto.PrivateKey        { return u.Key }

func GenerateCertificate(email, domain string) error {
	user := &User{Email: email}

	config := lego.NewConfig(user)
	client, err := lego.NewClient(config)
	if err != nil {
		return fmt.Errorf("failed to create LEGO client: %w", err)
	}

	request := certificate.ObtainRequest{
		Domains: []string{domain},
		Bundle:  true,
	}

	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		return fmt.Errorf("failed to obtain certificate: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("/etc/letsencrypt/live/%s/fullchain.pem", domain), certificates.Certificate, 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("/etc/letsencrypt/live/%s/privkey.pem", domain), certificates.PrivateKey, 0644)
	return err
}
