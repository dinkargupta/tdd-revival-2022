**Setup on WSL, Windows 11 (as on December 26th 2021) on Ubuntu 2004.2021.825.0**

Already had git installed

**installed Python 3.10**
1. sudo apt update
2. sudo apt install software-properties-common
3. sudo add-apt-repository ppa:deadsnakes/ppa
4. sudo apt update
5. sudo apt install python3.10

**installed golang**
1. sudo apt install golang

then start with guidance from the book [Learning Test Driven Development](https://www.oreilly.com/library/view/learning-test-driven-development/9781098106461/) to revive and refine TDD practices in context of current development context with languages Python, Go and Javascript 

started... 
Created a new workspce for this and then followed the instructions with three sub folder for each languages go, JS and Python

***Just a suggestion.. be careful with your nodeJS and go installations.. can eat up your 2 hours to fix this***

1. For go, the offical site is ok
2. For node, follow advise [Here](https://askubuntu.com/questions/426750/how-can-i-update-my-nodejs-to-the-latest-version)

One other time saving trick for us using Windows Subsystem for Linux(WSL) for local development and VS Code as IDE..

Use [WSL Remote VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) for this... then you can work from within IDE for WSL Ubuntu deployment.. One doesn't need to keep switching between Ubuntu console and IDE...

