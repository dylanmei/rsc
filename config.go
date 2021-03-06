package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Basic configuration settings required by all clients
type ClientConfig struct {
	Account   int    // RightScale account ID
	LoginHost string // RightScale API login host, e.g. "us-3.rightscale.com"
	Email     string // RightScale API login email
	Password  string // RightScale API login password
}

// LoadConfig loads the client configuration from disk
func LoadConfig(path string) (*ClientConfig, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config ClientConfig
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	config.Password, err = Decrypt(config.Password)
	return &config, err
}

// Save config encrypts the password and persists the config to file
func (cfg *ClientConfig) Save(path string) error {
	encrypted, err := Encrypt(cfg.Password)
	if err != nil {
		return fmt.Errorf("Failed to encrypt password: %s", err)
	}
	cfg.Password = encrypted
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Failed to serialize config: %s", err)
	}
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write config file: %s", err)
	}
	return nil
}

// Create configuration file and save it to file at given path
func CreateConfig(path string) error {
	config, _ := LoadConfig(path)
	var emailDef, passwordDef, accountDef, hostDef string
	if config != nil {
		yn := PromptConfirmation("Found existing configuration file %v, overwrite? (y/N): ", path)
		if yn != "y" {
			PrintSuccess("Exiting")
			return nil
		}
		emailDef = fmt.Sprintf(" (%v)", config.Email)
		accountDef = fmt.Sprintf(" (%v)", config.Account)
		passwordDef = " (leave blank to leave unchanged)"
		if config.LoginHost == "" {
			config.LoginHost = "my.rightscale.com"
		}
		hostDef = fmt.Sprintf(" (%v)", config.LoginHost)
	} else {
		config = &ClientConfig{}
	}

	fmt.Fprintf(out, "Account id%v: ", accountDef)
	var newAccount string
	fmt.Fscanln(in, &newAccount)
	if newAccount != "" {
		a, err := strconv.Atoi(newAccount)
		if err != nil {
			return fmt.Errorf("Account ID must be an integer, got '%s'.", newAccount)
		}
		config.Account = a
	}

	fmt.Fprintf(out, "Login email%v: ", emailDef)
	var newEmail string
	fmt.Fscanln(in, &newEmail)
	if newEmail != "" {
		config.Email = newEmail
	}

	fmt.Fprintf(out, "Login password%v: ", passwordDef)
	var newPassword string
	fmt.Fscanln(in, &newPassword)
	if newPassword != "" {
		config.Password = newPassword
	}

	fmt.Fprintf(out, "API Login host%v: ", hostDef)
	var newLoginHost string
	fmt.Fscanln(in, &newLoginHost)
	if newLoginHost != "" {
		config.LoginHost = newLoginHost
	}

	err := config.Save(path)
	if err != nil {
		return fmt.Errorf("Failed to save config: %s", err)
	}

	return nil
}
