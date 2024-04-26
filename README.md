# GoSSH

1. Clone the repository or download the binary

Create a directory named '.gossh' on your home folder
  ```bash
  mkdir ~/.gossh
  ```

2. Move the binary into the folder
  ```bash
  mv gossh ~/.gossh
  ```
3. Add it to your `$PATH`
  Open bashrc
  ```bash
  sudo nano ~/.bashrc
  ```
  add the following line
  ```bash
  export PATH="$HOME/.gossh:$PATH"
  ```
  update the bashrc
  ```bash
  source ~/.bashrc
  ```
4. Run gossh
  ```bash
  gossh add <name> <username@ip>
  ```

Linux only.

It is necessary to have `sshfs` installed in order to mount file systems.

Usage:

  To SSH into a server: gossh <server_name>

  To mount a server's file system: gossh fs <server_name>
  
  To add a server: gossh add <name> <username@ip>
  
  To remove a server: gossh remove <name>
  
  To list available servers: gossh list
