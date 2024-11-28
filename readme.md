# What is `clone`

clone is an app to manage cloning of different repositories using a predefined
host and path to clone to. It's not a big deal, but it's useful for my workflow
of clonning different repositories from different sources using different ssh
keys so I don't have to remember all the configs or `cat` the file everytime.

`clone` is useful for listing hosts, and managing "workspaces" as well as
clonning repositories under those workspaces. Now, what is a workspace: A
workspace is a record of some host and path to clone to. Prehaps you have seen
ssh config files with the following contents:
```
Host github.com-personal
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_rsa

Host github.com-work
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_rsa_work
```
Of course, with many more configurations that eventually become hard to remember
each time you need to clone a repo. A host is then the string
`gibhub.com-personal`. The path can be any valid path you can write to in your
os

`clone` will write a file into your `.config` folder called `clone.yaml` where
it will store all the workspaces so you can check there as well if you want to
have a look.


## Building and installing

As of now, the only way to use `clone` locally is to build it and install it.
Luckily, it's very simple:
1. Clone the repo with `git clone https://github.com/JulianH99/clone.git`
2. `cd clone && go mod download`
3. `go install`

if you have your `$GOPATH` set up correctly, you should be able to run `clone`
without issues

## Usage

### Clone basics
Command shape:
```sh
clone get [domainName] [user/repo]  
```
by default it'll clone in the current directory
Examples
```
clone endrock Endrock/luseta
```

you can use a saved workspace
```sh
clone [domainName] [giturl] -w work
```

each workspace is composed by:
- name
- path

you can also clone to a custom path with

```
clone [domainName] [gitUrl] -p /path/to/project
```

In both cases (`-w` and `-p` ) the project name will be appended to the path.
`-w` will take precedence over the `-p` flag.

### Domain names
You can list available domain names in your ~/.ssh/config file
```
clone domains list
```
This will show a list of all the domains inside the ~/.ssh/config file

and, if you want to see all the available configurations, you can write
```
clone domains list --full
```
which will essentially list all the configuration options inside the ~/.ssh/config file

### Workspaces
you can create, list, edit, and delete workspaces

#### list
```
clone workspaces list
```
#### edit
```
clone workspaces edit [workspaceName]
```

#### delete
```
clone workspaces delete [workspaceName]
```

#### create
```
clone workspace create
```

## Contributing
if there's any feature you'd like to see implemented, feel free to open an
issue. This project is still under development but I think it's stable enough

## Planned features
- [ ] Support for themes
- [ ] Full TUI besides command support
- [ ] Workspace and domain linking
