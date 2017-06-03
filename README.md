# mcdemo
```
Store and retrieve large files from memcached

Usage:
  mcdemo [command]

Available Commands:
  get         Get a stored file.
  help        Help about any command
  put         Store a file to memcached.

Flags:
      --config string   config file (default is .mcdemo.yaml)
  -h, --help            help for mcdemo

Use "mcdemo [command] --help" for more information about a command.
```
### mcdemo put
```
Store a file to memcached. One positional argument - the filename

Usage:
  mcdemo put [flags]

Flags:
  -h, --help   help for put

Global Flags:
      --config string   config file (default is .mcdemo.yaml)
```

### mcdemo get
```
Get a stored file. Two positions arguments - the original filename and the output file to write to

Usage:
  mcdemo get [flags]

Flags:
  -h, --help   help for get

Global Flags:
      --config string   config file (default is .mcdemo.yaml)

```
