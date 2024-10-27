# gossh
Allows you to SSH into servers and mount file systems by name.

### Install
Linux only.

1. Install directly from the repo using Go:
    ```bash
    go install https://github.com/patrickkdev/gossh@latest
    cd GoSSH
    ```

2. Ensure the go bin directory is in your path
    ```bash
    export PATH="$HOME/go/bin:$PATH"
    ```
    
8. Run gossh
    ```bash
    gossh
    ```

OBS: In order to mount file systems it is necessary to have `sshfs` installed

# Usage

Add a server
```bash
gossh add <server_name> <username@ip>
# Ex: gossh add my_new_vps root@100.50.100.30
```

List available servers
```bash
gossh list
```

SSH into a server
```bash
gossh <server_name>
# Ex: gossh my_new_vps
```

Mount a server's file system
```bash
gossh fs <server_name>
# Ex: gossh fs my_new_vps
```

Remove a server from the list:
```bash
gossh remove <server_name>
# Ex: gossh remove my_new_vps
```
