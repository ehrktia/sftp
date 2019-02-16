# SFTP program

sftp program to copy files from unix to windows shared drive
# Introduction
Main focus of the program is to copy files using sftp from a unix source to windows shared drive destination with neither inbuilt ssh client or open ssh option.

* Current version uses username and password authentication method against the source for establishing the ssh connection.
* Password is encoded and stored in a txt file.
* Option to encode password and generate text file are inbuilt.

# How to guide

1. To use this program follow the normal git clone process and get the entire source code available.
2. Once after cloning the repo , run the go main program.

*  This program takes command line arguments.
*  Please provide the following arguments to access various inbuilt options.

# Build Instructions for Windows
 - To build the executable file for windows please use the inbuilt go build command.
 >go build main.go

The build option generates the executable file for windows.

# Usage

To use the executable in Windows , follow the normal process ofexecuting an exe file.
Below are the list of options available.

`main.exe help`

To get the help options for the program.

`main.exe GenPwdFile` *password*

GenPwdFile - generates txt file with encoded pwd.
To generate the encoded pwd file from command prompt.
*password* - provide the  password in plain text.

NOTE - if you have special characters in the password please escape them using double-quotes.

`main.exe conn` *uname url srdir tgtdir*

conn option establish connection between client and the ssh server.
*uname* - user name in plain text
*url* - ssh server url
*srdir* -source dir from ssh server
*tgtdir* -target dir in ssh server

# Build&Test guidelines
## Help
`go run main.go help` 
Above command display the help option for the program.

## Encode password
`go run main.go GenPwdFile`  *password*

GenPwdFile - generates txt file with encoded pwd.
To generate the encoded pwd file from command prompt.
*password* - provide the  password in plain text.


## Connection
`go run conn` *uname* *tgtserver* *srdir* *tgtdir*

*conn* - Command Argument to start a ssh session to target server.
*uname* - username for the session
*tgtserver* - URL for the server where the connection is required.
*srdir* - source location from where the file is required to be copied.
*tgtserver* - target location to where the file is required to be moved.



## Limitations
Current version is only for copying file from source to Target which is managed in function.
### Work in progress
 - Passwordless authentication.
 - logging to file.
 - unit test for program.
 - dynamic option to do other than copy operations.

Please review and contact me with regards to comments, pull request and bugs.


