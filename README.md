[![CI](https://github.com/amplify-edge/shared/actions/workflows/ci-shared.yml/badge.svg)](https://github.com/amplify-edge/shared/actions/workflows/ci-shared.yml)

# Shared

Tools for devs to build.

The tools in this repo are used by all other repos.

## SSH local setup

- REF: Setup many accounts in ssh config:  https://medium.com/@xiaolishen/use-multiple-ssh-keys-for-different-github-accounts-on-the-same-computer-7d7103ca8693

- Delete the shit in your global git config. Your leaking ..
	- ``` git config --list ```

- same but shows where the values come from.
	- ``` git config -l ```

- Make a new key
	- ``` ssh-keygen -t rsa -b 4096 -C "userXXX_github" ```


- Add ssh key 
	- ``` ssh-add ~/.ssh/userXXX_github  ```

- List added ssh 
	- ``` ssh-add -l  ```


- Add the public key to github on your github web site
	- https://github.com/settings/keys

## SSH Config file

Needs this in it:

```bash

# userXXX
# https://github.com/userXXX/dev
# e.g: git clone git@github.com-userXXX:userXXX/dev.git
Host github.com-userXXX
 HostName github.com
 User git
 UseKeychain yes
 AddKeysToAgent yes
 IdentityFile ~/.ssh/userXXX_github
 
```


## Now you can build any of the code

All the code in all the repos does lots of code gen and so you need our tools to help.

Every repo including this one follows the same pattern with:

- To Build everything:
	- ``` make all ```

- To see the infos:
	- ``` make print ```

- To install all our tools:
	- ``` make dep ```


## CI

The github actions just call the make files, so that you can run CI locally, and not go nuts 2nd guessing what the github actions are doing...

- To Build everything:
	- ``` make all ```
