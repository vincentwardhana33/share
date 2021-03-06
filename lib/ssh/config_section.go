// Copyright 2020, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssh

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

const (
	// Valid values for AddKeysToAgent.
	valueAsk     = "ask"
	valueConfirm = "confirm"
	valueNo      = "no"
	valueYes     = "yes"
	valueAlways  = "always"

	// Valid values for AddressFamily.
	valueAny   = "any"
	valueInet  = "inet"
	valueInet6 = "inet6"
)

// List of default values.
const (
	defConnectionAttempts = 1
	defPort               = 22
	defStringPort         = "22"
	defXAuthLocation      = "/usr/X11R6/bin/xauth"
)

//
// ConfigSection is the type that represent SSH client Host and Match section
// in configuration.
//
type ConfigSection struct {
	AddKeysToAgent              string
	AddressFamily               string
	BindAddress                 string
	BindInterface               string
	CanonicalDomains            []string
	CanonicalizeHostname        string
	CanonicalizeMaxDots         int
	CanonicalizePermittedCNAMEs *PermittedCNAMEs
	CASignatureAlgorithms       []string
	CertificateFile             []string
	ConnectionAttempts          int
	ConnectTimeout              int

	// Environments contains system environment variables that will be
	// passed to Execute().
	// The key and value is derived from "SendEnv" and "SetEnv".
	Environments map[string]string

	Hostname                          string
	IdentityFile                      []string
	Port                              int
	User                              string
	XAuthLocation                     string
	IsBatchMode                       bool
	IsCanonicalizeFallbackLocal       bool
	IsChallengeResponseAuthentication bool
	IsCheckHostIP                     bool
	IsClearAllForwardings             bool
	UseCompression                    bool
	UseVisualHostKey                  bool

	stringPort string

	// List of SSH private keys.
	signers []ssh.Signer

	// User's home directory.
	homeDir string

	// workingDir contains the directory where the SSH client started.
	// This value is required when client want to copy file from/to
	// remote.
	// This field is optional, default to current working directory from
	// os.Getwd() or user's home directory.
	workingDir string

	// The first IdentityFile that exist and valid.
	privateKeyFile string

	// Patterns for Host section.
	patterns []*configPattern

	// Criteria for Match section.
	criteria    []*matchCriteria
	useCriteria bool

	useDefaultIdentityFile bool // Flag for the IdentityFile.
}

// newConfigSection create new Host or Match with default values.
func newConfigSection() *ConfigSection {
	return &ConfigSection{
		AddKeysToAgent: valueNo,
		AddressFamily:  valueAny,
		CASignatureAlgorithms: []string{
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
			ssh.KeyAlgoRSA,
		},
		ConnectionAttempts: defConnectionAttempts,
		Environments:       map[string]string{},
		IdentityFile: []string{
			"~/.ssh/id_dsa",
			"~/.ssh/id_ecdsa",
			"~/.ssh/id_ed25519",
			"~/.ssh/id_rsa",
		},
		Port:                              defPort,
		XAuthLocation:                     defXAuthLocation,
		useDefaultIdentityFile:            true,
		IsChallengeResponseAuthentication: true,
		IsCheckHostIP:                     true,
		stringPort:                        defStringPort,
	}
}

//
// generateSigners convert the IdentityFile to ssh.Signer for authentication
// using PublicKey.
//
func (section *ConfigSection) generateSigners() {
	section.signers = make([]ssh.Signer, 0, len(section.IdentityFile))

	for _, pkey := range section.IdentityFile {
		pkeyRaw, err := ioutil.ReadFile(pkey)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Printf("generateSigners: " + err.Error())
			}
			continue
		}

		signer, err := ssh.ParsePrivateKey(pkeyRaw)
		if err != nil {
			log.Printf("generateSigners: ParsePrivateKey: " + err.Error())
			continue
		}

		if len(section.privateKeyFile) == 0 {
			section.privateKeyFile = pkey
		}
		section.signers = append(section.signers, signer)
	}
}

//
// isMatch will return true if the string "s" match with one of Host or Match
// section.
//
func (section *ConfigSection) isMatch(s string) bool {
	if section.useCriteria {
		for _, criteria := range section.criteria {
			if criteria.isMatch(s) {
				return true
			}
		}
	} else {
		for _, pat := range section.patterns {
			if pat.isMatch(s) {
				return true
			}
		}
	}
	return false
}

//
// postConfig check, parse, and expand all of the fields values.
//
func (section *ConfigSection) postConfig(homeDir string) {
	var err error

	if len(homeDir) == 0 {
		section.homeDir, err = os.UserHomeDir()
		if err != nil {
			log.Println("ConfigSection.postConfig: " + err.Error())
		}
	} else {
		section.homeDir = homeDir
	}

	section.workingDir, err = os.Getwd()
	if err != nil {
		log.Println("ssh: cannot get working directory, default to user's home")
		section.workingDir = section.homeDir
	}

	for x, identFile := range section.IdentityFile {
		if identFile[0] == '~' {
			section.IdentityFile[x] = strings.Replace(identFile,
				"~", homeDir, 1)
		}
	}

	section.generateSigners()
}

func (section *ConfigSection) setAddKeysToAgent(val string) (err error) {
	switch val {
	case valueAsk, valueConfirm, valueNo, valueYes:
		section.AddKeysToAgent = val
	default:
		return fmt.Errorf("%s: invalid value %q", keyAddKeysToAgent,
			val)
	}
	return nil
}

func (section *ConfigSection) setAddressFamily(val string) (err error) {
	switch val {
	case valueAny, valueInet, valueInet6:
		section.AddressFamily = val
	default:
		return fmt.Errorf("%s: invalid value %q", keyAddressFamily,
			val)
	}
	return nil
}

func (section *ConfigSection) setCanonicalizeHostname(val string) (err error) {
	switch val {
	case valueNo, valueAlways, valueYes:
		section.CanonicalizeHostname = val
	default:
		return fmt.Errorf("%s: invalid value %q", keyBatchMode, val)
	}
	return nil
}

func (section *ConfigSection) setCanonicalizePermittedCNAMEs(val string) (err error) {
	sourceTarget := strings.Split(val, ":")
	if len(sourceTarget) != 2 {
		return fmt.Errorf("%s: invalid rule",
			keyCanonicalizePermittedCNAMEs)
	}

	listSource := strings.Split(sourceTarget[0], ",")
	sources := make([]*configPattern, 0, len(listSource))
	for _, domain := range listSource {
		src := newConfigPattern(domain)
		sources = append(sources, src)
	}

	listTarget := strings.Split(sourceTarget[1], ",")
	targets := make([]*configPattern, 0, len(listTarget))
	for _, domain := range listTarget {
		target := newConfigPattern(domain)
		targets = append(targets, target)
	}

	section.CanonicalizePermittedCNAMEs = &PermittedCNAMEs{
		sources: sources,
		targets: targets,
	}
	return nil
}

func (section *ConfigSection) setCASignatureAlgorithms(val string) {
	section.CASignatureAlgorithms = strings.Split(val, ",")
}

//
// setEnv set the Environments with key and value of format "KEY=VALUE".
//
func (section *ConfigSection) setEnv(env string) {
	kv := strings.SplitN(env, "=", 2)
	if len(kv) == 2 {
		section.Environments[kv[0]] = kv[1]
	}
}

func (section *ConfigSection) setSendEnv(envs map[string]string, pattern string) {
	for k, v := range envs {
		ok, _ := filepath.Match(pattern, k)
		if ok {
			section.Environments[k] = v
		}
	}
}
