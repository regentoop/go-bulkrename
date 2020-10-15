# go-bulkrename
A simple script to read a csv that contains list old and new file names then renames the files in the folder from the old name to the new name.
The script assumes that the files to be renamed, the compiled script and a csv containing old and new names are all in the same folder.

The file containing the old name and new name of the file must be a simple flat csv with the following format:
| old name | new name |
| :------- | --------:|
| oldname1 | newname1 |
| oldname2 | newname2 |

Note that the names must be in the first two columns since the script checks those columns for them.

HOW TO COMPILE
You can clone this repo with git clone 

```
  git clone <git-url>

```

and then compile with go build.

```
  go build
```

USAGE:

You can also run without compiling using the go run command
```
  go run ./rename.go
```
But after building all you need is the rename command.

```
  ./rename  names_file.csv   #this works for linux
    rename  names_file.csv   #this might be better for windows
```

TO DOs:
1. Add command line options to identify the columns containing old and new names
2. Add option to interactively specify file path(s) and columns containing old and new names
3. Probably add support for advanced spreadsheets like xlsx, odt etc
