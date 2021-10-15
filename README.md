##### Introduction
just have a try to delete harbor's image by using
a tool which developed by golang

##### How to use
cautions: In my opinion，this is just a trial run.so if you want to use
this tool to delete your harbor's image,you should do many tests first. 
In addition，You can follow the steps below to delete images：
1. clone the code by following command

`git clone https://github.com/sunsharing-note/harbor.git` 

2. compile

`go build -o harbor .`

3. get help

```
./harbor -h
a registry to store and pull or push image

Usage:
  harbor [flags]
  harbor [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  project     to operator project
  repo        to operator repository
  tag         to operator image

Flags:
  -h, --help   help for harbor

Use "harbor [command] --help" for more information about a command.
```
##### Env
| component  | version  |
|  ----  | ----  |
| golang  | 1.14.1  |
| harbor  | v2.3.1 |