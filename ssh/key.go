package key

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

type SSHKeyProfile struct {
	Name           string `yaml:"name"`
	Description    string `yaml:"description"`
	GitConfigEmail string `yaml:"git_config_email"`
	GitConfigUser  string `yaml:"git_config_username"`
}

type Config struct {
	SSHKeys []SSHKeyProfile `yaml:"ssh_keys"`
}

func ReadConfig() (Config, error) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func GenerateSSHKey(keyName string, config Config) {
	if keyName == "all" {
		fmt.Println("Generating all SSH keys in ~/.ssh")
		for _, profile := range config.SSHKeys {
			generateKey(profile.Name, profile)
		}
	} else {
		var foundProfile SSHKeyProfile
		for _, profile := range config.SSHKeys {
			if profile.Name == keyName {
				foundProfile = profile
				break
			}
		}

		if foundProfile.Name != "" {
			generateKey(foundProfile.Name, foundProfile)
		} else {
			log.Fatalf("Profile with name %s not found in the configuration", keyName)
		}
	}
}

func generateKey(keyName string, profile SSHKeyProfile) {
	GitConfigUser := profile.GitConfigUser
	GitConfigEmail := profile.GitConfigEmail

	privateKeyPath := fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"), keyName)
	publicKeyPath := fmt.Sprintf("%s/.ssh/%s.pub", os.Getenv("HOME"), keyName)

	if _, err := os.Stat(privateKeyPath); err == nil {
		fmt.Printf("SSH KEY: %s already exists at %s\n", GitConfigUser, privateKeyPath)
	} else if os.IsNotExist(err) {
		cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", GitConfigEmail, "-f", privateKeyPath)
		cmd.Stdin = strings.NewReader("\n")
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error generating SSH KEY: %s: %v", keyName, err)
		}
		fmt.Printf("Generated SSH KEY: %s at %s\n", keyName, privateKeyPath)

		if len(GitConfigUser) > 0 && len(GitConfigEmail) > 0 {
			cmd := exec.Command("ssh-keygen", "-y", "-f", privateKeyPath)
			publicKey, err := cmd.Output()
			if err != nil {
				log.Fatalf("Error generating public key for %s: %v", keyName, err)
			}
			ioutil.WriteFile(publicKeyPath, publicKey, 0644)
			fmt.Printf("Generated public key for %s at %s\n", keyName, publicKeyPath)
		}
	}
}

func DeleteSSHKey(keyName string) {
	if keyName == "all" {
		fmt.Println("Deleting all SSH keys in ~/.ssh")
		files, err := ioutil.ReadDir(fmt.Sprintf("%s/.ssh", os.Getenv("HOME")))
		if err != nil {
			log.Fatalf("Error reading ~/.ssh directory: %v", err)
			return
		}
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".pub") {
				privateKeyPath := fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"), strings.TrimSuffix(file.Name(), ".pub"))
				publicKeyPath := fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"), file.Name())
				deleteKey(privateKeyPath, publicKeyPath)
			}
		}
	} else {
		privateKeyPath := fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"), keyName)
		publicKeyPath := fmt.Sprintf("%s/.ssh/%s.pub", os.Getenv("HOME"), keyName)
		deleteKey(privateKeyPath, publicKeyPath)
	}
}

func deleteKey(privateKeyPath, publicKeyPath string) {
	if _, err := os.Stat(privateKeyPath); err == nil {
		fmt.Printf("Deleting SSH KEY: %s at %s\n", privateKeyPath, privateKeyPath)
		if err := os.Remove(privateKeyPath); err != nil {
			log.Fatalf("Error deleting SSH key %s: %v", privateKeyPath, err)
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("SSH KEY: %s does not exist at %s\n", privateKeyPath, privateKeyPath)
	}

	if _, err := os.Stat(publicKeyPath); err == nil {
		fmt.Printf("Deleting public key for %s at %s\n", privateKeyPath, publicKeyPath)
		if err := os.Remove(publicKeyPath); err != nil {
			log.Fatalf("Error deleting public key %s: %v", publicKeyPath, err)
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("Public key for %s does not exist at %s\n", privateKeyPath, publicKeyPath)
	}
}

func SetGitConfig(keyName string) {
	config, _ := ReadConfig()

	for _, profile := range config.SSHKeys {
		if profile.Name == keyName {
			cmd := exec.Command("git", "config", "--global", "user.name", profile.GitConfigUser)
			if err := cmd.Run(); err != nil {
				log.Fatalf("Error setting Git user name: %v", err)
			}
			fmt.Printf("Set Git user name to: %s\n", profile.GitConfigUser)

			cmd = exec.Command("git", "config", "--global", "user.email", profile.GitConfigEmail)
			if err := cmd.Run(); err != nil {
				log.Fatalf("Error setting Git user email: %v", err)
			}
			fmt.Printf("Set Git user email to: %s\n", profile.GitConfigEmail)

			return
		}
	}

	log.Fatalf("Profile with name %s not found in the configuration", keyName)
}
