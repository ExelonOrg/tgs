# tgs (terragrunt scaffolder)

This repository is for a command line tool that helps users build a terragrunt project quickly in a semi-automated fashion.

![release](https://img.shields.io/github/v/release/ExelonOrg/tgs)  ![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ExelonOrg/tgs/Go/main)


## Quick Start

1. Download release binary and place it in your $PATH
2. Rename binary e.g. (tgs.exe)
3. Create workspace directory e.g. ```mkdir workspace && cd workspace```
4. Inside workspace directory run ```tgs init```
5. Edit ```.tgs/config.json``` to match your desired scaffolding; the key names of the ```stacks``` property will be the names of each of your stacks or patterns
6. Create your stacks/patterns by running ```tgs stack create <pattern-name>```; You should create a stack for every item listed in ```config.json```.
7. Edit each of your stack files to your desired configuration
8. Run ```tgs scaffold create``` to create your scaffolding