# Git basics

## Fundamentals
* git is distributed version control system
* created by Linus Travolds (creater of Linux)

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
```
