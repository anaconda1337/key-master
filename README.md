## GO key-master - developed by [Anaconda1337](https://github.com/anaconda1337)
_key-master is a simple command-line tool for managing SSH keys and associated Git configurations. It allows you to generate SSH keys, delete them, and set Git user information, all from the command line._

<hr>

### Installation
Before using key-master, ensure that you have Go (Golang) installed on your system. You can download and install Go from the official website: [Go Downloads](https://go.dev/dl/)

1. Clone the key-master repository to your local machine:
```bash
git clone github.com/anaconda1337/key-master
```

2. Change to the key-master directory:
```bash
cd key-master
```

3. Build the key-master binary:
```bash
go build
```

4. Now you can run key-master from the command line:
```bash
./key-master
```

### Usage

#### Generate SSH Keys
You can use key-master to generate SSH keys, either for a single profile or for all profiles defined in a configuration file. To generate keys, use the following command:
```bash
./key-master generate
```
- To generate keys for all profiles:
```bash
./key-master generate all
```
- To generate keys for a single profile:
```bash
./key-master generate <key-name>
```

#### Delete SSH Keys
You can use key-master to delete SSH keys, either for a single profile or for all profiles defined in a configuration file. To delete keys, use the following command:
```bash
./key-master delete
```
- To delete keys for all profiles:
```bash
./key-master delete all
```
- To delete keys for a single profile:
```bash
./key-master delete <key-name>
```

#### Set Git User Information
You can use key-master to set Git user information, either for a single profile or for all profiles defined in a configuration file. To set Git user information, use the following command:
```bash
./key-master config <key-name>
```

### Configuration
key-master reads configuration details from a "config.yml" file, which should be present in the same directory as the key-master executable. The configuration file should define SSH key profiles, each with a unique name, description, Git username, and Git email.

Here's an example of the configuration format in "config.yml":
```yaml
# Description: Configuration file for ssh-key-manager

ssh_keys:
  - name: ssh_key_1
    description: My primary SSH key
    git_config_username: my_github_username_0
    git_config_email: my_github_email_0

  - name: ssh_key_2
    description: Secondary SSH key
    git_config_username: my_github_username_1
    git_config_email: my_github_email_1

```

### License
This module is open-source and available under the [MIT License](https://opensource.org/license/mit/).

<hr>

- Please feel free to contribute to this project. I am happy about every contribution. :smiley:
- Give me a star if you like the project. :star:
- Give me a follow if you want to see more projects from me. :heart:
- Provide feedback if you have any suggestions. :speech_balloon:
- Provide ideas if you have any. :bulb: