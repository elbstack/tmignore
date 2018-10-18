package main

import (
	"github.com/iralution/tmignore/cmd"
)

/*
TODO cp original plist file to ~/.plistbackup/...
TODO add cobra for commandline interaction
TODO add blob logic to ignore all directories (not just node_modules)
TODO add filelist to ~/.plist to know which paths were added earlier
TODO flags: tmignore filepattern --debug --test --rootDir
TODO tmignore restore
 */


func main() {
	cmd.Execute()
}