# Mile High Gophers Job Board #

In order to better share Go-related job openings in the Denver area, and to learn some new things, we are creating a Mile High Gophers Job Board.

Visit the [wiki](https://github.com/WWG-Denver-Boulder/job_board/wiki) for an overview of the project.

Each week we will break up into teams and work on issues related to the project. Contibutors can take as large or small a role as you are able to.

## Getting Set Up ##

### Install Go ###
Refer to our [Getting Started repo](https://github.com/WWG-Denver-Boulder/getting_started) for tools and helper scripts. Please note this
is still in development and suggestions are welcome!

### Create Directories ###
* `mkdir go && cd go`
* `mkdir src && cd src`
* `mkdir github.com && cd github.com`

### Clone Repo ###
* `git clone git@github.com:WWG-Denver-Boulder/job_board.git`

Your path should look like: `~/go/src/github.com/job_board`

### Create Branch ###
`git checkout -b <your_branch_name>`

Example: `git checkout -b adds_getter_function`

### Commit Changes ###
`git status` (See which files have changes)

`git add <file(s)_you_changed>`

`git commit -m "<your commit message>"`

Example:
`git add users_t.go`

`git commit -m "starts getter function for users_t"`

### Push to Remote ###
`git push origin <you_branch_name>`

Example: `git push origin adds_getter_function`