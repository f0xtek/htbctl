package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hack The Planet!")

	// $ htbctl new boxname
	// $ htbctl archive boxname
	// $ htbctl restore boxname

	// Recursively create /root/hackthebox if it doesn't exist
	// Recursively create /root/backups/hackthebox if it doesn't exist

	// NEW BOXES
	// Create /root/hackthebox/boxname
	// Create /root/hackthebox/boxname/nmap directory
	// Create /root/hackthebox/boxname/NOTES.md file (empty)

	// ARCHIVE BOXES
	// Compress /root/hackthebox/boxname directory to /root/backups/hackthebox/boxname.{tar.gz,zip}
	// Delete /root/hackthebox/boxname

	// RESTORE BOXES
	// Uncompress /root/backups/hackthebox/boxname to /root/hackthebox/boxname
}
