# htbctl

`htbctl` is a command line tool for managing your Hack The Box projects. The aim is to keep a nice organised directory of boxes that you are hacking on.

- [htbctl](#htbctl)
  - [Installation](#installation)
    - [Binary Release](#binary-release)
    - [Build From Source](#build-from-source)
  - [Create A New Project](#create-a-new-project)
  - [Archive A Project](#archive-a-project)
  - [Restore A Project](#restore-a-project)

## Installation

### Build From Source

Please ensure you meet the following prerequisites:

- Go 1.13+
- Git

```bash
$ git clone https://gitlab.com/thesoy_sauce/htbctl.git && \
cd htbctl && \
go mod download && \
go install
```

## Create A New Project

To create a new project, use the `new` subcommand and provide a name for your project (e.g. the name of the box you are hacking on) via the `--name` flag:

`htbctl new --name SwagShop` or `htbctl n --name SwagShop`

This will create a skeleton directory structure in `$HOME/hackthebox/SwapShop` like so:

```bash
$ ls -al $HOME/hackthebox/SwagShop/
total 12K
drwxr-xr-x 3 neo neo 4.0K Dec 20 23:19 .
drwxr-xr-x 3 neo neo 4.0K Dec 20 23:19 ..
drwxr-xr-x 2 neo neo 4.0K Dec 20 23:19 nmap
-rw-r--r-- 1 neo neo    0 Dec 20 23:19 NOTES.md
```

## Archive A Project

Once you have rooted a box, and are finished with your work, you can archive the project using the `archive` subcommand:

`htbctl archive --name SwagShop` or `htbctl a -n SwagShop`

This will create a gzipped tar archive of the `$HOME/hackthebox/SwagShop` directory in `$HOME/backup/hackthebox`:

```bash
$ ls -la $HOME/backup/hackthebox
total 12K
drwxr-xr-x 2 neo neo 4.0K Dec 20 23:24 .
drwxr-xr-x 3 neo neo 4.0K Dec 20 23:10 ..
-rw-rw---- 1 neo neo  174 Dec 20 23:24 SwagShop.tgz
```

**The original `$HOME/hackthebox/SwagShop` directory will be removed**:

```bash
$ ls -la $HOME/hackthebox
total 8.0K
drwxr-xr-x  2 neo neo 4.0K Dec 20 23:24 .
drwxr-xr-x 23 neo neo 4.0K Dec 20 23:20 ..
```

## Restore A Project

Shoudl you wish to revisit a box, you can restore a previously archived project:

`htbctl restore --name SwagShop` or `htbctl r -n SwahShop`

This will uncompress & extrcat the `SwagShop.tgz` file back to its original location `$HOME/hackthebox/SwagShop`:

```bash
$ ls -la $HOME/hackthebox/SwagShop
total 12K
drwxr-xr-x 3 neo neo 4.0K Dec 20 23:28 .
drwxr-xr-x 3 neo neo 4.0K Dec 20 23:28 ..
drwxr-xr-x 2 neo neo 4.0K Dec 20 23:28 nmap
-rw-r--r-- 1 neo neo    0 Dec 20 23:28 NOTES.md
```
