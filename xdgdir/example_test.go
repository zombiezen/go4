/*
Copyright 2017 The go4 Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package xdgdir_test

import (
	"fmt"
	"io"
	"os"

	"go4.org/xdgdir"
)

func ExampleDir_Path() {
	// Print the absolute path of the current user's XDG_CONFIG_DIR.
	fmt.Println(xdgdir.Config.Path())
}

func ExampleDir_Open() {
	// Read a file from $XDG_CONFIG_DIR/myconfig.json.
	// This will search for a file named "myconfig.json" inside
	// $XDG_CONFIG_DIR and then each entry inside $XDG_CONFIG_DIRS.
	// It opens and returns the first file it finds, or returns an error.
	f, err := xdgdir.Data.Open("myconfig.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()
	// Copy to stdout.
	if _, err := io.Copy(os.Stdout, f); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func ExampleDir_Create() {
	// Write a file to $XDG_DATA_DIR/myapp/foo.txt
	f, err := xdgdir.Data.Create("myapp/foo.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	_, err = fmt.Fprintln(f, "Hello, World!")
	cerr := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if cerr != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
