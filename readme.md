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
clone  [user/repo]  
```
This will ask the user to choose a host and a workspace, with the posibility to
choose no workspace. It is important to note that a host is required, otherwise
this program makes no sense and you can just use git clone as normal.  
This command offers the following options as well:
```
clone [owner/repo] -w [worspkace] -s [host] -p [customPath]
```
- `-w` will take the workspace name and will not ask for the worskpace again in
  the form

- `-t` will take the host name and will not ask for the host again in the form.
  You can type the hole form or just the hostname (e.g. `github.com-personal` or
just `personal`)
- `-p` will be a custom path to git clone to and will take precedence over the
workspace's path

### host names
You can list available host names in your ~/.ssh/config file
```
clone hosts 
```
This will show a list of all the hosts inside the ~/.ssh/config file, the same
list appears of course when you run `clone [owner/repo]` without additional
arguments


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
- [ ] origin assignment, for cases like `git remote add origin
git@github.com-domain:owner/repo.git`, with an interface like `clone [owner/repo] -r [remoteName]`
