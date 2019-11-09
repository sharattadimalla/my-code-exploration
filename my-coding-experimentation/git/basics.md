# Git basics

## Fundamentals
* git is open source  version control system
* created by Linus Travolds (creater of Linux)
* git project consists of 3 main directories - `the working directory`, `the staging directory` and `the git directory`.
* **the working directory** is where you add, edit and delete files
* The changes are staged into **the staging area**
* After you commit your changes, the snapshot of the files are copied into **the git directory**

## Basic git commands

```
# initialize local git repo
git init

# clone a remote repository
git clone https://path/to/remote-repository

# add files to staging area
git add filename.txt

# set user specific config values
git config --global user.name name

# display the list of changed files together with the files yet to be committed
git status

# send local commits to remote branch
git push origin master

# create a new branch and automatically switch to it
git checkout -b branch_name

# view all remote repositories
git remote -v

# add local repo to a remote
git remote add origin https://path/to/remote-repo

# remove a connection to the specified remote repo
git remote rm origin

# list all branches
git branch

# remove a local branch
git branch -d branch_name

# merge a branch into an active one
git merge branch_name

# list conflict agains the base file
git diff --base filename

# view conflicts between branches
git diff source_branch target_branch

# list all present conflicts
git diff

# tag a specific commit
git tag commit_id

# view repo's history along with few commit details
git log

# reset the index and working directory to the last commit state
git reset --hard HEAD

# Remove files from the index and the working directory
git rm filename.txt

# Temporarily save changes that are not ready to be comitted
git stash

# view info about any git object
git show

# fetch all objects from the remote repository that don't exist in the local repo
git fetch origin branch

# clean unnecessary files and optimize the local repo
git gc

# create a zip or tar file with the contents of a single repo tree
git archive --format=tar master

# deletes objects that don't have incoming pointers
git prune

# apply certain changes from one branch to another
git rebase master

```

## git config

* git provides 3 levels of configuration - system, user (global) and repo
* system wide config can be found `/etc/gitconfig` file
* user level config can be found `~/.gitconfig`
* repo level config can be found `.git/config`

### git config commands

```
# list git global configs
git config --global --list

# set user name
git config --global user.name git_user_id

# set user email
git config --global user.email batman@gotham_city.gov

# set default editor
git config --global core.editor vim
```
