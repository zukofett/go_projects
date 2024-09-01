package main

import (
	"cli/utils"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
)

type flags_ty struct {
    isAll       bool
    isLong      bool
    isInode     bool
    isClassify  bool
    isOnlyFiles bool
    isOnlyDirs  bool
    isDecor     bool
}

type file_ty struct {
    inode     string
    fileType  string
    perms     string
    links     string
    owner     string
    group     string
    size      string
    timestamp string
    name      string
}

func (f file_ty) printFile() {
    permissions := f.fileType + f.perms
    fmt.Print(f.inode, permissions, f.links, f.owner, f.group, f.size, f.timestamp, f.name)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


func main() {
    app := cli.NewApp() 
    app.Name = "ls"
    app.Usage = "ls [flags] [command][args]"
    app.Author = "Idan Baker"
    app.Version = "0.1.0"
    app.Compiled = time.Now()
    app.Copyright = "MIT Licensed"
    app.EnableBashCompletion = true
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "a, all",
            Usage: "do not ignore entries starting with .",
        },
        cli.BoolFlag{
            Name: "l, long",
            Usage: "use a long listing format",
        },
        cli.BoolFlag{
            Name: "i, inode",
            Usage: "print the index number of each file",
        },
        cli.BoolFlag{
            Name: "F, classify",
            Usage: "append indicator (one of */=>@|) to entries",
        },
        cli.BoolFlag{
            Name: "f, files",
            Usage: "include only regular files",
        },
        cli.BoolFlag{
            Name: "d, dirs",
            Usage: "include only directories",
        },
    }
    app.Action = func(c *cli.Context) error {
        decor := c.Bool("a") || c.Bool("l") || c.Bool("i") || c.Bool("F") || c.Bool("f") || c.Bool("d")
        flags := flags_ty{
            c.Bool("a"), c.Bool("l"), c.Bool("i"), c.Bool("F"), c.Bool("f"), c.Bool("d"), decor,
        }
        walk_path := "."
        if c.NArg() > 0 {
            walk_path = c.Args()[0]
        }
        parseDir(walk_path, flags) 
        return nil
    }

    app.Run(os.Args)
}


func parseDir(parse_path string, flags flags_ty) {
    err := filepath.WalkDir(parse_path, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
            return err
        }
        if d.IsDir() && flags.isOnlyFiles {
            return nil
        }
        if !d.IsDir() && flags.isOnlyDirs {
            return nil
        }
        if !flags.isAll && path[0] == '.' {
            return nil
        }
        
        file := file_ty{}
        if flags.isDecor {
            file.name = colorFileName(d.Name(), d.Type())
        }
        if flags.isClassify {
            // add file classify
        }

        if flags.isInode {
            // add inode
        }
        if flags.isLong {
            // add other fields
        }

        // print file

        return nil

    }) 
    checkErr(err)
}

func colorFileName(fileName string, fileType fs.FileMode) string {
    output := ""
    switch fileType.String() {
    case "d":
        if fileName
        output = utils.BLUE + fileName + utils.DEFAULT + "/"
    case "L":
        output = utils.CYAN + fileName + utils.DEFAULT + "@"
    case "D":
    case "c":
        output = utils.BLACKBG + utils.BOLDYELLOW + fileName + utils.DEFAULT
    case "p":
        output = utils.BLACKBG + utils.YELLOW + fileName + "|"
    case "S":
        output = utils.MAGENTA + fileName + utils.DEFAULT + "="
    } 

}
