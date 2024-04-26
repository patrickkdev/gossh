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
x
    source ~/.bashrc
    ```

7. Run gossh
    ```bash
    gossh
    ```

It is necessary to have `sshfs` installed in order to mount file systems.

# Usage

Add a server

    ```bash
    gossh add <name> <username@ip>
    ```

List available servers
    ```bash
    gossh list
    ```

SSH into a server
    ```bash
    gossh <server_name>
    ```

Mount a server's file system
    ```bash
    gossh fs <server_name>
    ```

Remove a server from the list:
    ```bash
    gossh remove <name>
    ```
