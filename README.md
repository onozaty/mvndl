# mvndl

Download the files from the maven repository.

## Usage

```
$ mvndl -r jcenter -g com.github.onozaty -a postgresql-copy-helper -v 1.0.0 -d local-repo
```

The arguments are as follows.

```
Usage of mvndl:
  -r, --repository string   Maven repository ("jcenter" or "central" or Specify by url)
  -g, --group string        Group ID
  -a, --artifact string     Artifact ID
  -v, --version string      Version
  -d, --dest string         Destination directory for download files
  -h, --help                Help
```

The downloaded files will be saved in the same structure as the maven repository.

```
$ ./mvndl -r jcenter -g com.github.onozaty -a postgresql-copy-helper -v 1.0.0 -d local-repo
mvndl vX.X.X (xxxxxxx)
https://jcenter.bintray.com/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0.pom -> saved
https://jcenter.bintray.com/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0.jar -> saved
https://jcenter.bintray.com/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0-sources.jar -> saved
https://jcenter.bintray.com/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0-javadoc.jar -> saved

$ find local-repo/ -type f
local-repo/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0-javadoc.jar
local-repo/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0-sources.jar
local-repo/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0.jar
local-repo/com/github/onozaty/postgresql-copy-helper/1.0.0/postgresql-copy-helper-1.0.0.pom
```

## Install

You can download the binary from the following.

* https://github.com/onozaty/mvndl/releases/latest

## License

MIT

## Author

[onozaty](https://github.com/onozaty)
