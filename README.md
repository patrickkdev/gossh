# GoSSH

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

6. Update the bash
    ```bash
    source ~/.bashrc
    ```

7. Run gossh
    ```bash
    gossh add <name> <username@ip>
    ```

It is necessary to have `sshfs` installed in order to mount file systems.

# Usage
    To add a server:
    ```bash
        gossh add <name> <username@ip>
    ```

    To SSH into a server: 
    ```bash
    gossh <server_name>
    ```
    
    To mount a server's file system:
    ```bash
    gossh fs <server_name>
    ```
    
    To remove a server:
    ```bash
    gossh remove <name>
    ```
    
    To list available servers:
    ```bash
    gossh list
    ```
