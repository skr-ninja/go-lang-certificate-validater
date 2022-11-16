package models

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	Domainname string    `gorm:"size:255;not null;unique" json:"domainname"`
	Expiry     time.Time `gorm:"size:255;not null;unique" json:"expiry"`
	Issuedate  time.Time `gorm:"size:255;not null;unique" json:"issuedate"`
}

func (u *Certificate) SaveCertificate() (*Certificate, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Certificate{}, err
	}
	return u, nil
}

func ExtrctCert() (*Certificate, error) {

	bs := googleCert

	block, _ := pem.Decode(bs)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the public key")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("failed parse x509 public cert: %v", err)
	}

	c := Certificate{}
	c.Domainname = cert.DNSNames[0]
	c.Expiry = cert.NotAfter
	c.Issuedate = cert.NotBefore

	cr, err := c.SaveCertificate()
	if err != nil {
		return &Certificate{}, err
	}

	log.Printf("Subject:   %q", cert.Subject)
	log.Printf("Domain names: %+v", cert.DNSNames)
	log.Printf("Before: %+v", cert.NotBefore)
	log.Printf("Expiry: %+v", cert.NotAfter)
	return cr, nil
}

var googleCert = []byte(`-----BEGIN CERTIFICATE-----
// Please Paste Your Cert
-----END CERTIFICATE-----`)

func UsingFile(filename string) (*Certificate, error) {

	//bs, err := os.ReadFile("./temp/google.txt")
	bs, err := os.ReadFile("./temp/" + filename)
	if err != nil {
		log.Fatalf("failed read cert file: %v", err)
	}
	block, _ := pem.Decode(bs)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the public key")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("failed parse x509 public cert: %v", err)
	}

	c := Certificate{}
	c.Domainname = cert.DNSNames[0]
	c.Expiry = cert.NotAfter
	c.Issuedate = cert.NotBefore

	cr, err := c.SaveCertificate()
	if err != nil {
		return &Certificate{}, err
	}

	log.Printf("Subject:   %q", cert.Subject)
	log.Printf("Domain names: %+v", cert.DNSNames)
	log.Printf("Before: %+v", cert.NotBefore)
	log.Printf("Expiry: %+v", cert.NotAfter)
	return cr, nil
}
