package main

import (
	key "key-master/ssh"
	log "log"
	os "os"
)

func main() {
	config, err := key.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading the configuration: %v", err)
	}

	args := os.Args[1:]
	if len(args) == 2 && args[0] == "generate" {
		keyName := args[1]
		key.GenerateSSHKey(keyName, config)
		log.Printf("SSH key generated for profile: %s", keyName)
	} else if len(args) == 2 && args[0] == "delete" {
		keyName := args[1]
		key.DeleteSSHKey(keyName)
		log.Printf("SSH key %s deleted.", keyName)
	} else if len(args) == 2 && args[0] == "config" {
		keyName := args[1]
		key.SetGitConfig(keyName)
		log.Printf("SSH key %s set as GitHub config.", keyName)
	} else {
		log.Println("Usage: key-master generate|delete|config [key-name]")
	}
}
