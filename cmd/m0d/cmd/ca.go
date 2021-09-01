package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmos "github.com/tendermint/tendermint/libs/os"
)

var (
	Country            = []string{"CN"}
	Province           = []string{"BJ"}
	Organization       = []string{"TQ"}
	OrganizationalUnit = []string{"M0"}
)

func generateNewSerialNumber() *big.Int {
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128)) // [0, 2^128[
	if err != nil || serial.Sign() <= 0 {
		if err != nil {
			fmt.Printf("Warning : failed to generate a random serial number : %s. It will be set to 1, consider use --server-cert-serial arg.\n", err)
		}
		return big.NewInt(1)
	}
	return serial
}

func generateCA(serial *big.Int) (*x509.Certificate, *rsa.PrivateKey, error) {
	// Generate 2048bit RSA Key
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Cannot generate RSA key : %s", err))
	}

	if serial == nil {
		serial = generateNewSerialNumber()
	}
	// Generate CA Certificate
	certTemplate := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			Country:            Country,
			Province:           Province,
			Organization:       Organization,
			OrganizationalUnit: OrganizationalUnit,
			CommonName:         "RootCA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(20 * 365 * 24 * time.Hour), // 20 years
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCRLSign | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		IsCA:                  true,
		BasicConstraintsValid: true,
	}

	return &certTemplate, key, nil
}

func generateCertificate(
	serial *big.Int,
	domains []string,
	caCert *x509.Certificate) (*x509.Certificate, *rsa.PrivateKey, error) {
	if len(domains) == 0 {
		return nil, nil, errors.New(fmt.Sprintf("No domains provided for generating server certificate."))
	}

	if serial == nil {
		serial = generateNewSerialNumber()
	}
	// Generate 2048bit RSA Key
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Cannot generate RSA key : %s", err))
	}

	serverCertTemplate := x509.Certificate{
		SerialNumber: serial,
		Issuer:       caCert.Subject,
		Subject: pkix.Name{
			Country:            Country,
			Province:           Province,
			Organization:       Organization,
			OrganizationalUnit: OrganizationalUnit,
			CommonName:         domains[0],
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour), // 10 years
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		DNSNames:              domains,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	return &serverCertTemplate, key, nil
}

func parseCert(filepath string) (*x509.Certificate, error) {
	caCertFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot read %s : %s", filepath, err))
	}
	caCertPem, _ := pem.Decode(caCertFile)
	if caCertPem == nil {
		return nil, errors.New(fmt.Sprintf("Failed to PEM decode %s.", filepath))
	}
	caCert, err := x509.ParseCertificate(caCertPem.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot parse Certificate %s : %s", filepath, err))
	}

	return caCert, nil
}

func parseRsaKey(filepath string) (*rsa.PrivateKey, error) {
	caKeyFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot read %s : %s", filepath, err))
	}
	caKeyPem, _ := pem.Decode(caKeyFile)
	if caKeyPem == nil {
		return nil, errors.New(fmt.Sprintf("Failed to PEM decode %s.", filepath))
	}

	der := caKeyPem.Bytes
	caKey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot parse RSA Key %s (missing password ?) : %s", filepath, err))
	}

	return caKey, nil
}

func saveCert(certFilename string, keyFilename string, caCert *x509.Certificate, caKey *rsa.PrivateKey, cert *x509.Certificate, key *rsa.PrivateKey) error {
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, key.Public(), caKey)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot create CA certificate : %s", err))
	}

	// Write Certificate
	caCertOut, err := os.OpenFile(certFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot open file %s : %s", certFilename, err))
	}
	err = pem.Encode(caCertOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot encode CA Certificate in PEM %s : %s", certFilename, err))
	}
	err = caCertOut.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot close %s : %s", certFilename, err))
	}

	// write Key
	caKeyOut, err := os.OpenFile(keyFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot open file %s : %s", keyFilename, err))
	}
	err = pem.Encode(caKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot encode RSA key in PEM %s : %s", keyFilename, err))
	}
	err = caKeyOut.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot close %s : %s", keyFilename, err))
	}
	return nil
}

const (
	flagkey    = "key-file"
	flagCert   = "cert-file"
	flagSerial = "cert-serial"
)

func GenCACommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-ca",
		Short: "generate a ca certificate, (either newly generated or recovered), and save to disk",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := filepath.Join(viper.GetString(flags.FlagHome), "ca")
			if err := tmos.EnsureDir(dir, 0755); err != nil {
				return err
			}

			keyFile := filepath.Join(dir, fmt.Sprintf("%s.key", "ca"))
			certFile := filepath.Join(dir, fmt.Sprintf("%s.cert", "ca"))
			if tmos.FileExists(keyFile) || tmos.FileExists(certFile) {
				return fmt.Errorf("certificate %s already exist", certFile)
			}

			rKeyFile := viper.GetString(flagkey)
			rCertFile := viper.GetString(flagCert)
			if len(rKeyFile) == 0 || len(rCertFile) == 0 {
				cert, key, err := generateCA(nil)
				if err != nil {
					return err
				}

				if err := saveCert(certFile, keyFile, cert, key, cert, key); err != nil {
					return fmt.Errorf("cannot save custom ca certificate : %s", err)
				}
			} else {
				cert, err := parseCert(rCertFile)
				if err != nil {
					return fmt.Errorf("cannot parse custom ca certificate : %s", err)
				}

				key, err := parseRsaKey(rKeyFile)
				if err != nil {
					return fmt.Errorf("cannot parse custom ca rsa key : %s", err)
				}

				if err := saveCert(certFile, keyFile, cert, key, cert, key); err != nil {
					return fmt.Errorf("cannot save custom ca certificate : %s", err)
				}
			}
			return nil
		},
	}

	cmd.Flags().String(flagkey, "", "path to existing certificate pem, recover instead of creating.")
	cmd.Flags().String(flagCert, "", "path to existing rsa key pem, recover instead of creating")
	cmd.Flags().Int(flagSerial, 0, "custom serial number for the certificate.")

	cmd.SetOut(cmd.OutOrStdout())
	cmd.SetErr(cmd.ErrOrStderr())

	return cmd
}

func GenCertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-cert [node-id]",
		Short: "generate node certificate, (either newly generated or recovered), and save to disk",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			nodeid := args[0]
			dir := filepath.Join(viper.GetString(flags.FlagHome), "ca")
			if err := tmos.EnsureDir(dir, 0755); err != nil {
				return err
			}
			keyFile := filepath.Join(dir, fmt.Sprintf("%s.key", nodeid))
			certFile := filepath.Join(dir, fmt.Sprintf("%s.cert", nodeid))
			if tmos.FileExists(keyFile) || tmos.FileExists(certFile) {
				return fmt.Errorf("certificate already exist")
			}

			caKeyFile := filepath.Join(dir, fmt.Sprintf("%s.key", "ca"))
			caCertFile := filepath.Join(dir, fmt.Sprintf("%s.cert", "ca"))
			if !tmos.FileExists(caKeyFile) || !tmos.FileExists(caCertFile) {
				return fmt.Errorf("ca certificate not exist")
			}

			caCert, err := parseCert(caCertFile)
			if err != nil {
				return fmt.Errorf("cannot parse custom ca certificate : %s", err)
			}

			caKey, err := parseRsaKey(caKeyFile)
			if err != nil {
				return fmt.Errorf("cannot parse custom ca rsa key : %s", err)
			}

			rKeyFile := viper.GetString(flagkey)
			rCertFile := viper.GetString(flagCert)
			if len(rKeyFile) == 0 || len(rCertFile) == 0 {
				cert, key, err := generateCertificate(nil, []string{nodeid}, caCert)
				if err != nil {
					return err
				}

				if err := saveCert(certFile, keyFile, caCert, caKey, cert, key); err != nil {
					return fmt.Errorf("cannot save custom ca certificate : %s", err)
				}
				fmt.Println(fmt.Sprintf("%s %s", cert.Issuer, cert.SerialNumber))
			} else {
				cert, err := parseCert(rCertFile)
				if err != nil {
					return fmt.Errorf("cannot parse custom ca certificate : %s", err)
				}

				key, err := parseRsaKey(rKeyFile)
				if err != nil {
					return fmt.Errorf("cannot parse custom ca rsa key : %s", err)
				}

				if err := saveCert(certFile, keyFile, caCert, caKey, cert, key); err != nil {
					return fmt.Errorf("cannot save custom certificate : %s", err)
				}
				fmt.Println(fmt.Sprintf("%s %s", cert.Issuer, cert.SerialNumber))
			}
			return nil
		},
	}

	cmd.Flags().String(flagkey, "", "path to existing certificate pem, recover instead of creating.")
	cmd.Flags().String(flagCert, "", "path to existing rsa key pem, recover instead of creating")
	cmd.Flags().Int(flagSerial, 0, "custom serial number for the certificate.")

	cmd.SetOut(cmd.OutOrStdout())
	cmd.SetErr(cmd.ErrOrStderr())

	return cmd
}

func caCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca-server",
		Short: "Manage your application's certificate",
	}

	cmd.AddCommand(
		GenCACommand(),
		GenCertCommand(),
	)

	cmd.PersistentFlags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}
