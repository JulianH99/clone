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

## Ussage

There are several commands in `clone`:

- `clone hosts list`: will list all the hosts found in the `~/.ssh/config` file
- `clone workspaces list`: will list all workspaces available in the
`~/.config/clone.yaml` file
- `clone workspaces add`: will show a form to create a new workspace pointing to
  a path and using a host from the list of hosts in `~/.ssh/config`
- `clone worksaces delete`: will let you select a workspace from your config
file and remove it

### get command
Now, there's also the `clone get` command, which lets you pass an repository's
ssh url to clone from and lets you choose between two options:
1. Custom Configuration: you can specify where you want to clone to and a host
   (optional)
2. Saved configuration: you can specify a workspace from your saved workspaces
   to use as path to clone to and a host

additionaly, you can specify a sub folder when clonning to both custom and saved
configuration with the `-s` flag when running the command. For example, running
`clone get git@github.com:username/repo.git -s sub_folder` will clone to
whatever path you choose (from custom or saved config) and append `sub_folder`
at the end.

The goal is not to replace a simple `git clone` entirely but to aid when there
are too many ssh configurations to remember (which is my case) and ease the
proccess of cloning to defined paths where the repositories belonging to those
configurations will likely point to.

## Contributing
if there's any feature you'd like to see implemented, feel free to open an
issue. This project is still under development but I think it's stable enough

## Planned features
- [ ] Support for themes
- [ ] Full TUI besides command support
