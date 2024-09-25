package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/tyler-smith/go-bip39"
)

const (
	duration            = 5 * time.Second
	minutesPerViolation = 10
	minutesPerHour      = 60
	hoursPerDay         = 24
	daysPerYear         = 365.25 // Account for leap years
)

var totalViolations uint64
var totalProcessingYears float64

func incrementViolations() {
	atomic.AddUint64(&totalViolations, 1)
}

func calculateProcessingYears(violations uint64) float64 {
	return float64(violations) * minutesPerViolation / minutesPerHour / hoursPerDay / daysPerYear
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	incrementViolations()
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func generateBIP39Seed() (string, error) {
	incrementViolations()
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func generateGPGKey() error {
	incrementViolations()
	_, err := openpgp.NewEntity("", "", "", nil)
	return err
}

func generateSSHKey() error {
	incrementViolations()
	_, err := rsa.GenerateKey(rand.Reader, 2048)
	return err
}

func generateSSLCert() error {
	incrementViolations()
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	// In a real scenario, you'd save these to files
	// But for our benchmark, we'll just discard them
	_ = derBytes
	return nil
}

func benchmark(name string, f func() error) {
	start := time.Now()
	count := 0
	initialViolations := atomic.LoadUint64(&totalViolations)

	for time.Since(start) < duration {
		err := f()
		if err != nil {
			fmt.Printf("Error in %s: %v\n", name, err)
			return
		}
		count++
	}

	rate := float64(count) / duration.Seconds()
	currentViolations := atomic.LoadUint64(&totalViolations)
	newViolations := currentViolations - initialViolations
	newProcessingYears := calculateProcessingYears(newViolations)

	totalProcessingYears += newProcessingYears

	fmt.Printf("%s: %.2f per second\n", name, rate)
	fmt.Printf("New 'violations' in this benchmark: %d\n", newViolations)
	fmt.Printf("Time to process new violations: %.4f years\n", newProcessingYears)
	fmt.Printf("Total 'violations' so far: %d\n", currentViolations)
	fmt.Printf("Cumulative time to process all violations: %.4f years\n", totalProcessingYears)
	fmt.Println()
}

func main() {
	fmt.Println("Welcome to the 'Definitely Not a Wallet' Key Generator")
	fmt.Println("Warning: Each key generation is a 'violation'. The Empire is watching.")
	fmt.Printf("Running benchmarks for %v each...\n\n", duration)

	benchmark("ECDSA Private Keys", func() error {
		_, err := generatePrivateKey()
		return err
	})

	benchmark("BIP39 Seeds", func() error {
		_, err := generateBIP39Seed()
		return err
	})

	benchmark("GPG Keys (Go implementation)", func() error {
		return generateGPGKey()
	})

	benchmark("SSH Keys", func() error {
		return generateSSHKey()
	})

	benchmark("SSL Certificates", func() error {
		return generateSSLCert()
	})

	violations := atomic.LoadUint64(&totalViolations)
	fmt.Printf("\nTotal 'violations' committed: %d\n", violations)
	fmt.Printf("If each violation takes %d minutes of bureaucratic processing...\n", minutesPerViolation)
	fmt.Printf("Time to process all violations: %.4f years\n", totalProcessingYears)
	fmt.Println("\nRemember: The Empire's regulations are definitely enforceable and not at all absurd.")
}
