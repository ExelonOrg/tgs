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
9. Complete your terragrunt configuration by doing the following (at least):
  - Add provider information to _base_modules modules
  - Add module code in main.tf files
  - Fill out variables.tf and outputs.tf for all base modules
  - add inputs terragrunt.hcl files (global.hcl, dev.hcl, terragrunt.hcl, non_production.hcl etc.)
  
  
 
## Scaffolding
this tool generates a base_modules folder:
```
.
└── workingdir/
    ├── __base_modules/                                 folder that contains all the base modules that your project will use
    │   ├── api/
    │   │   ├── api.hcl                                 environment file that defines global inputs for all api projects; defines module source
    │   │   ├── main.tf                                 top level terraform module; calls to shared tf registry or local module
    │   │   ├── variables.tf                            input file that defines variables in main.tf            
    │   │   └── outputs.tf                              output file that defines outputs that can be consumed by other modules\*
    │   └── sql/
    |       |...
    ...
 ```
This folder will contain all the *terraform* code for the building blocks of your terragrunt project. For example, if your architecuture requires a webserver, storage solution, and caching service in various configurations across different environments, you might have ```webserver/```, ```sql/```, and ```redis/``` sub folders that each contain terraform code to deploy those specific resources.
 
 Terragrunt allows you to pass information between the different modules without having to control them with one singular terraform state file. More on this later.
 
 ```
    |   ### GROUPS ###
    ├── non_production/
    |   ├── non_production.hcl                           file that contains remote state configuration
    |   | ### ENVIRONMENTS ###
    │   ├── dev/
    |   |   | ### MODULES ###
    │   │   ├── sql/
    │   │   └── api/
    |   |       | #### APPS ###
    │   │       ├── ace/
    |   |       |   |__ dev.ace.east.appsettings.tpl    appsettings template for dev/ace
    │   │       │   └── terragrunt.hcl                  configuration for dev/ace api; merges dev.env.hcl, _base_modules/api/env.hcl, terragrunt.hcl
    │   │       ├── bge/
    │   │       │   └── terragrunt.hcl
    │   │       ├── comed/
    │   │       │   └── terragrunt.hcl
    │   │       ├── dpl/
    │   │       │   └── terragrunt.hcl
    │   │       ├── peco/
    │   │       │   └── terragrunt.hcl
    │   │       ├── pepco/
    │   │       │   └── terragrunt.hcl
    |   |       |__ dev.env.hcl
    │   │       └── terragrunt.hcl                    * in this directory, you can run terragrunt run-all [init|plan|apply|destroy]
    │   ├── test/
    │   │   ├── sql/
    │   │   └── api/
    │   │       ├── ace/
    │   │       ├── bge/
    │   │       ├── ...
    │   │       └── terragrunt.hcl
    │   └── stage/
    │       └── api/
    │           ├── ace/
    |           |   |__ stage.ace.east.appsettings.tpl
    |           |   |__ stage.ace.central.appsettings.tpl 
    │           ├── bge/
    │           ├── ...
    │           └── terragrunt.hcl                     contains input variables for two regions
    └── production/
        └── prod/
 │       └── api/
    │           ├── ace/
    |           |   |__ prod.ace.east.appsettings.tpl
    |           |   |__ prod.ace.central.appsettings.tpl 
    │           ├── bge/
    │           ├── ...
    │           └── terragrunt.hcl



```
