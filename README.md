# GoSSH
Allows you to SSH into servers and mount file systems by name.

### Install
Linux only.

1. Clone the repository or download the binary
    ```bash
    git clone https://github.com/patrickkdev/GoSSH.git
    cd GoSSH
    ```

2. Create a directory named '.gossh' on your home folder
    ```bash
    mkdir ~/.gossh
    ```

3. Move the binary into the folder
    ```bash
    mv gossh ~/.gossh
    ```

4. Open bashrc with your text editor
    ```bash
    sudo nano ~/.bashrc
    ```

5. Add the following line, save and close the editor
    ```bash
    export PATH="$HOME/.gossh:$PATH"
    ```
6. Update bash
    ```bash
    source ~/.bashrc
    ```

8. Run gossh
    ```bash
    gossh
    ```

It is necessary to have `sshfs` installed in order to mount file systems.

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
