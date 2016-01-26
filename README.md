# Directory Monitor

Watches a directory tree or single file and executes the specified command in response to changes.

`dmon` will react to the following changes
* nodes being added to a directory tree
* nodes being removed from a directory tree
* node mtimes changing

## Usage

    $ dmon <dir> <cmd> [args ...]

* dir: the directory to watch - can be absolute or relative to the current directory
* cmd: the command to execute
* args: any arguments required by cmd - these are just passed as you would to the command itself

## Examples

To watch the current directory (recursively) and run tests on change issue

    $ dmon . go test

To watch a single file and auto commit changes issue

    $ dmon working.txt git commit -am "edit"

## Installation

To install simply issue

    $ go get github.com/toshaf/dmon

If your `$GOBIN` is on your path, you should be in business
