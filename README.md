# multipkg
A tool that removes the differences between package installers from linux distro's

## Status
I just started working on this app and it's a quite big idea so it might take some time before it suports the major linux distrobutions and maybe windows some day  
:heavy_plus_sign: = Working on, :heavy_check_mark: = Mostly dune, :x: No development yet
- (:heavy_plus_sign:) The base app
- (:heavy_check_mark:) Solus
- (:x:) Fedora
- (:heavy_plus_sign:) Arch linux
- (:x:) Ubuntu 
- (:x:) Windows *chocolatey*
For a more extensive status read: [docs/supported.md](./docs/supported.md)

## Setup:
```
$ go get github.com/mjarkk/multipkg
```

## Usage:
```
$ multipkg --help

  Usage: multipkg [options] [command]

  Options:

    -f, --force        Force command
    -y, --yes          Automaticly input yes for next
    --help             Help menu
    --version          App info
    --debug            Log get debug data 

  Commands:

    install|in|i       <program>  Install a program
    reinstall|rein|ri  <program>  Re-install a program
    remove|re|r        <program>  Remove a program from the system
    update|up|u        *<program> Update a program or the complete system
    search|se|s        <program>  Search for programs
    info|inf           <program>  Get info about a specific programs
```
