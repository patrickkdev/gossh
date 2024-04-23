# GoSSH


Linux only.

It is necessary to have `sshfs` installed in order to mount file systems.

Usage:

  To SSH into a server: gossh <server_name>

  To mount a server's file system: gossh fs <server_name>
  
  To add a server: gossh add <name> <username@ip>
  
  To remove a server: gossh remove <name>
  
  To list available servers: gossh list
