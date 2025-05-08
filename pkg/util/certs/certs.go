package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

type certConfig struct {
	OrganisationName  string
	ServiceNames      []string
	ServiceNameSpaces []string
	CertsPath         string
	CaCert            string
	ServerCert        string
	ServerKey         string
}

func newSerialNumber() (*big.Int, error) {
	klog.V(3).Infof("generating new serial number")
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return nil, errors.New("failed to generate serial number: " + err.Error())
	}

	return serialNumber, nil
}

func main() {
	organisationName := os.Getenv("ORGANISATION_NAME")
	serviceNamesString := os.Getenv("SERVICE_NAMES")
	serviceNameSpacesString := os.Getenv("SERVICE_NAMESPACES")
	certsPath := os.Getenv("CERTS_PATH")
	caCert := os.Getenv("CA_CERT")
	serverCert := os.Getenv("SERVER_CERT")
	serverKey := os.Getenv("SERVER_KEY")

	serviceNames := strings.Split(serviceNamesString, ",")
	serviceNameSpaces := strings.Split(serviceNameSpacesString, ",")

	if len(serviceNames) == 1 && serviceNames[0] == "" {
		serviceNames = []string{}
	}

	if len(serviceNameSpaces) == 1 && serviceNameSpaces[0] == "" {
		serviceNameSpaces = []string{}
	}

	for i, serviceName := range serviceNames {
		serviceNames[i] = strings.TrimSpace(serviceName)
	}

	for i, serviceNameSpace := range serviceNameSpaces {
		serviceNameSpaces[i] = strings.TrimSpace(serviceNameSpace)
	}

	myCertConfig := certConfig{
		OrganisationName:  organisationName,
		ServiceNames:      serviceNames,
		ServiceNameSpaces: serviceNameSpaces,
		CertsPath:         certsPath,
		CaCert:            caCert,
		ServerCert:        serverCert,
		ServerKey:         serverKey,
	}
	err := generateCerts(myCertConfig)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func generateCerts(config certConfig) error {
	err := config.validateCertConfig()
	if err != nil {
		return err
	}

	klog.V(1).Infof("generating certificates")

	serialNumber, err := newSerialNumber()
	if err != nil {
		return err
	}

	klog.V(4).Infof("generating config for ca certificates")
	caConfig := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{config.OrganisationName},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(100, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	klog.V(4).Infof("generating ca private key")
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	klog.V(4).Infof("generating ca certificate")
	caCert, err := x509.CreateCertificate(rand.Reader, caConfig, caConfig, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return err
	}

	klog.V(4).Infof("generating config for server certificate")

	dnsNames := []string{}
	for _, serviceName := range config.ServiceNames {
		dnsNames = append(dnsNames, serviceName)
		for _, namespace := range config.ServiceNameSpaces {
			dnsNames = append(dnsNames, serviceName+"."+namespace)
			dnsNames = append(dnsNames, serviceName, "."+namespace+".svc")
		}
	}

	serialNumber, err = newSerialNumber()
	if err != nil {
		return err
	}

	serverConfig := &x509.Certificate{
		DNSNames:     dnsNames,
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   dnsNames[2],
			Organization: []string{config.OrganisationName},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	klog.V(4).Infof("generating server private key")
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	klog.V(4).Infof("signing server certificate")
	serverCert, err := x509.CreateCertificate(rand.Reader, serverConfig, caConfig, &serverPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return err
	}

	klog.V(4).Infof("pem encoding ca certificate, service certificate and server private key")
	caPem := new(bytes.Buffer)
	err = pem.Encode(caPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCert,
	})
	if err != nil {
		return err
	}

	serverCertPem := new(bytes.Buffer)
	err = pem.Encode(serverCertPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: serverCert,
	})
	if err != nil {
		return err
	}

	serverKeyPem := new(bytes.Buffer)
	err = pem.Encode(serverKeyPem, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(serverPrivateKey),
	})
	if err != nil {
		return err
	}

	klog.V(4).Infof("writing ca certificate: %s, service certificate: %s, server private key: %s, to folder: %s", config.CaCert, config.ServerCert, config.ServerKey, config.CertsPath)
	if _, err := os.Stat(config.CertsPath); os.IsNotExist(err) {
		err = os.MkdirAll(config.CertsPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = writeFile(filepath.Join(config.CertsPath, config.CaCert), caPem)
	if err != nil {
		return err
	}

	err = writeFile(filepath.Join(config.CertsPath, config.ServerCert), serverCertPem)
	if err != nil {
		return err
	}

	err = writeFile(filepath.Join(config.CertsPath, config.ServerKey), serverKeyPem)
	if err != nil {
		return err
	}

	return nil
}

func writeFile(filepath string, sCert *bytes.Buffer) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(sCert.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (c *certConfig) validateCertConfig() error {
	if c.OrganisationName == "" {
		return getErrorWithMissingMessage("ORGANISATION_NAME")
	}

	if len(c.ServiceNames) == 0 {
		return errors.New("SERVICE_NAMES cannot be empty, at least one service required, comma separated")
	}

	if len(c.ServiceNameSpaces) == 0 {
		return errors.New("SERVICE_NAMESPACES cannot be empty, at least one service required, comma separated")
	}

	if c.CertsPath == "" {
		return getErrorWithMissingMessage("CERTS_PATH")
	}

	if c.CaCert == "" {
		return getErrorWithMissingMessage("CA_CERT")
	}

	if c.ServerCert == "" {
		return getErrorWithMissingMessage("SERVER_CERT")
	}

	if c.ServerKey == "" {
		return getErrorWithMissingMessage("SERVER_KEY")
	}
	return nil
}

func getErrorWithMissingMessage(fieldName string) error {
	return errors.New(fmt.Sprintf("Field '%s' was unexpectedly empty", fieldName))
}
