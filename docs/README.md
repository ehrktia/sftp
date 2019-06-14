# SFTP program

sftp program to copy files from unix to windows shared drive
# Quick Links


* [introduction](#introduction)
* [How-to-guide](#how-to-guide)
* [Build-Instructions-for-Windows](#build-instructions-for-windows)
   * [Usage](#usage)
* [Build test guidelines](#build-test-guidelines)
   * [Help](#help)
   * [Encode Password](#encode-password)
   * [Connection](#connection)
 * [Limitations](#limitations)
 * [Work-in-Progress](#work-in-progress)
   



## Introduction
Main focus of the program is to copy files using sftp from a unix source to windows shared drive destination with neither inbuilt ssh client or open ssh option.

* Current version uses username and password authentication method against the source for establishing the ssh connection.
* Password is encoded and stored in a txt file.
* Option to encode password and generate text file are inbuilt.

## How-to-guide

1. To use this program follow the normal git clone process and get the entire source code available.
2. Once after cloning the repo , run the go main program.

*  This program takes command line arguments.
*  Please provide the following arguments to access various inbuilt options.

## Build-Instructions-for-Windows
 - To build the executable file for windows please use the inbuilt go build command.
 `go build main.go`

The build option generates the executable file for windows.

### Usage

To use the executable in Windows , follow the normal process ofexecuting an exe file.
Below are the list of options available.

`main.exe help`

To get the help options for the program.

`main.exe pwdgen --pwd=password`

pwdgen - generates txt file with encoded pwd.</br>
To generate the encoded pwd file from command prompt.</br>
*password* - provide the  password in plain text.</br>

NOTE - if you have special characters in the password please escape them using double-quotes.

`main.exe conn --serv=servername:port --uname=username --cmd=mv  --src=srcdir  --tgt=tgtdir` 

conn option establish connection between client and the ssh server.</br>
--uname : username in plain text.</br>
--serv : ssh server url with a port (generatlly the default port is 22).</br>
--cmd : command which is required to be executed in ssh remote server.</br>
--srdir : source dir from ssh server.</br>
--tgtdir:  target dir in ssh server.</br>

## Build Test guidelines
### Help
`go run main.go help` 
Above command display the help option for the program.

### Encode password
`go run main.go pwdgen --pwd=password`  

pwdgen - generates txt file with encoded pwd.</br>
To generate the encoded pwd file from command prompt.</br>
*password* - provide the  password in plain text.</br>


### Connection
`go run main.go conn --serv=servername:port --uname=username --cmd=mv  --src=srcdir  --tgt=tgtdir` 

conn option establish connection between client and the ssh server.</br>
--uname : username in plain text.</br>
--serv : ssh server url with a port (generatlly the default port is 22).</br>
--cmd : command which is required to be executed in ssh remote server.</br>
--srdir : source dir from ssh server.</br>
--tgtdir:  target dir in ssh server.</br>


## Limitations
- Passwordless authentication require key setup and key gen.
### Work in progress
 - Passwordless authentication.
 - logging to file.
 - unit test for program.
 

Please review and contact me with regards to comments, pull request and bugs.


