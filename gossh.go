package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	serversConfigFile = "~/.gossh/servers.conf"
	mountBaseDir      = "/mnt"
)

func expandPath(path string) string {
    if strings.HasPrefix(path, "~") {
        home, err := os.UserHomeDir()
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }

        path = filepath.Join(home, path[1:])
    }

    return path
}

func loadServersConfig() map[string]string {
    servers := make(map[string]string)
    file, err := os.Open(expandPath(serversConfigFile))
    if err != nil {
        return servers
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) == 2 {
            servers[parts[0]] = parts[1]
        }
    }

    return servers
}

func saveServersConfig(servers map[string]string) error {
    path := expandPath(serversConfigFile)
    err := os.MkdirAll(filepath.Dir(path), 0700)
    if err != nil {
        return err
    }

    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    for name, ip := range servers {
        _, err := fmt.Fprintf(file, "%s %s\n", name, ip)
        if err != nil {
            return err
        }
    }

    return nil
}

func sshIntoServer(ip string) {
    cmd := exec.Command("ssh", ip)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}

func mountFileSystem(serverName, ip string) error {
	mountPath := filepath.Join(mountBaseDir, serverName)
	if _, err := os.Stat(mountPath); os.IsNotExist(err) {
		err := os.Mkdir(mountPath, 0755)
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("sudo", "sshfs", "-o", "allow_other", fmt.Sprintf("%s:/", ip), mountPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func openPath(path string) error {
    cmd := exec.Command("open", path)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func addServer(name, address string, servers map[string]string) {
	servers[name] = address
	err := saveServersConfig(servers)
	if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
	}
	fmt.Printf("Server '%s' added successfully.\n", name)
}

func removeServer(name string, servers map[string]string) {
	_, ok := servers[name]
	if !ok {
			fmt.Printf("Server '%s' not found in configuration.\n", name)
			os.Exit(1)
	}
	delete(servers, name)
	err := saveServersConfig(servers)
	if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
	}
	fmt.Printf("Server '%s' removed successfully.\n", name)
}

func listServers(servers map[string]string) {
	fmt.Println("Available Servers:")
	for name, host := range servers {
		fmt.Printf("%s: %s\n", name, host)
	}
}

func main() {
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
    listCmd := flag.NewFlagSet("list", flag.ExitOnError)
    mountCmd := flag.NewFlagSet("fs", flag.ExitOnError)

    if len(os.Args) < 2 {
				fmt.Println("GoSSH allows you to SSH into servers and mount file systems by name.")
				fmt.Println();
				fmt.Println("<server_name> is any name you want to assign to the server, such as 'myserver', then you can SSH into the server with 'gossh myserver'.")
				fmt.Println();
        fmt.Println("Usage:")
        fmt.Println("  To SSH into a server: gossh <server_name>")
        fmt.Println("  To mount a server's file system: gossh fs <server_name>")
        fmt.Println("  To add a server: gossh add <name> <username@ip>")
        fmt.Println("  To remove a server: gossh remove <name>")
        fmt.Println("  To list available servers: gossh list")
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "add":
        addCmd.Parse(os.Args[2:])
        if addCmd.NArg() != 2 {
            fmt.Println("Usage: gossh add <name> <username@ip>")
            os.Exit(1)
        }
        name := addCmd.Arg(0)
        address := addCmd.Arg(1)
        servers := loadServersConfig()
        addServer(name, address, servers)
    case "remove":
        removeCmd.Parse(os.Args[2:])
        if removeCmd.NArg() != 1 {
            fmt.Println("Usage: gossh remove <name>")
            os.Exit(1)
        }
        name := removeCmd.Arg(0)
        servers := loadServersConfig()
        removeServer(name, servers)
    case "list":
        listCmd.Parse(os.Args[2:])
        servers := loadServersConfig()
        listServers(servers)
    case "fs":
        mountCmd.Parse(os.Args[2:])
        if mountCmd.NArg() != 1 {
					fmt.Println("Usage: gossh fs <server_name>")
            os.Exit(1)
        }
        serverName := mountCmd.Arg(0)
				servers := loadServersConfig()
				ip, ok := servers[serverName]
				if !ok {
					fmt.Printf("Server '%s' not found in configuration.\n", serverName)
					os.Exit(1)
				}
				err := mountFileSystem(serverName, ip)
				if err != nil {
					fmt.Println("Error mounting file system:", err)
					os.Exit(1)
				}
				fmt.Printf("File system of server '%s' mounted at %s.\n", serverName, filepath.Join(mountBaseDir, serverName))
				err = openPath(filepath.Join(mountBaseDir, serverName))
				if err != nil {
					fmt.Println("Error opening file system:", err)
					os.Exit(1)
				}
    default:
        serverName := os.Args[1]
        servers := loadServersConfig()
        ip, ok := servers[serverName]
        if !ok {
            fmt.Printf("Server '%s' not found in configuration.\n", serverName)
            os.Exit(1)
        }
        sshIntoServer(ip)
    }
}

