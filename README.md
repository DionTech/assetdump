# install

```zsh
go install
```

Recommended: set an alias.

# usage

## help

``` zsh 
assetdump -help
```

## new scan
scanning will be this and will store a file "devports.de.json" at current folder:

``` zsh
assetdump devports.de
```

## loading existing scans

### list available scans

``` zsh
assetdump -list
```

define specific path

```zsh
assetdump -list -path /var/home/
```

### pretty print complete json

loading this json is really simple, only

```zsh
assetdump -load -pretty devports.de
```

will pretty print the stored json

You can define the apth where to search for loading:

```zsh
assetdump -load -path /var/www/ -pretty devports.de
```

### define elements to output

```zsh
assetdump -load -ips devports.de
```

```zsh
assetdump -load -hosts devports.de
```

or combined:

```zsh
assetdump -load -hosts -ips devports.de
```

