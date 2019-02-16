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
## Help
>go run main.go help 

Above command display the help option for the program
## Encode password
> go run main.go GenPwdFile  password in plain text

GenPwdFile - option to geneate a text file with encoded pwd.
password - password in plain text for the user which you are planning to establish the connection.

Once excuted with the genpwdfile option it encodes plain text password to a text file and makes it available for the program to access it while execution.
## Connection
>go run conn uname tgtserver srdir tgtdir

conn - Command Argument to start a ssh session to target server.
uname - username for the session
tgtserver - URL for the server where the connection is required.
srdir - source location from where the file is required to be copied.
tgtserver - target location to where the file is required to be moved.

## Limitations
Current version is only for copying file from source to Target which is managed in function.
### Work in progress
 - Passwordless authentication.
 - logging to file.
 - unit test for program.
 - dynamic option to do other than copy operations.

Please review and contact me with regards to comments, pull request and bugs.

