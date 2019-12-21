package main

// TODO: template directory structure & files instead of hardcoding

import (
	"fmt"
	"log"
	"os"

	"github.com/mholt/archiver/v3"
	"github.com/urfave/cli/v2"
)

const (
	dirMode  = int(0755)
	fileMode = int(0660)
)

func main() {
	// $ htbctl new --name boxname
	// $ htbctl archive --name boxname
	// $ htbctl restore --name boxname

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	htbDir := fmt.Sprintf("%s/hackthebox", homeDir)
	htbBackupDir := fmt.Sprintf("%s/backup/hackthebox", homeDir)
	dirs := []string{htbDir, htbBackupDir}
	createDirs(dirs)

	app := &cli.App{
		Name:    "htbctl",
		Version: "1.0.0",
		Usage:   "Hack The Planet!",
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Create a new HTB project",
				Action: func(c *cli.Context) error {
					if err := newProject(htbDir, c.String("name")); err != nil {
						return err
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "name of the HTB project",
					},
				},
			},
			{
				Name:    "archive",
				Aliases: []string{"a"},
				Usage:   "Archive a HTB project",
				Action: func(c *cli.Context) error {
					if err := archiveProject(htbDir, htbBackupDir, c.String("name")); err != nil {
						return err
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "name of the HTB project",
					},
				},
			},
			{
				Name:    "restore",
				Aliases: []string{"r"},
				Usage:   "Restore a HTB project",
				Action: func(c *cli.Context) error {
					if err := restoreProject(htbDir, htbBackupDir, c.String("name")); err != nil {
						return err
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "name of the HTB project",
					},
				},
			},
		},
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Unrcognized command %s.\n", command)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// checkDirs recursively creates directories if they doesn't exist
func createDirs(dirs []string) {
	for _, d := range dirs {
		if _, err := os.Stat(d); os.IsNotExist(err) {
			os.MkdirAll(d, os.FileMode(dirMode))
		}
	}
}

// newProject creates a new skeleton HTB project directory & corresponding files
func newProject(h, n string) error {
	fmt.Println("Creating new project...")
	d := fmt.Sprintf("%s/%s/nmap", h, n)
	if err := os.MkdirAll(d, os.FileMode(dirMode)); err != nil {
		return err
	}

	notesFile, err := os.Create(fmt.Sprintf("%s/%s/NOTES.md", h, n))
	if err != nil {
		log.Fatal(err)
	}
	notesFile.Close()
	fmt.Printf("New project available at %s/%s\n", h, n)
	return nil
}

func archiveProject(h, d, n string) error {
	// ARCHIVE BOXES
	// Compress $HOME/hackthebox/boxname directory to $HOME/backup/hackthebox/boxname.{tar.gz,zip}
	src := fmt.Sprintf("%s/%s", h, n)
	dst := fmt.Sprintf("%s/%s.tgz", d, n)
	if err := archiver.Archive([]string{src}, dst); err != nil {
		return err
	}
	if err := os.Chmod(dst, os.FileMode(fileMode)); err != nil {
		return err
	}
	// Delete $HOME/hackthebox/boxname
	if err := os.RemoveAll(src); err != nil {
		return err
	}
	return nil
}

func restoreProject(h, d, n string) error {
	// RESTORE BOXES
	// Uncompress $HOME/backup/hackthebox/boxname to $HOME/hackthebox/boxname
	dst := fmt.Sprintf("%s", h)
	src := fmt.Sprintf("%s/%s.tgz", d, n)
	if err := archiver.Unarchive(src, dst); err != nil {
		return err
	}
	return nil
}
