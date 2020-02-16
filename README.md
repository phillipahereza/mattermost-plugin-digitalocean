# Mattermost Digital Ocean Plugin

## Table of Contents
- [1. Features](#1-features)
- [2. Configuration](#2-configuration)
- [3. Development](#3-development)

## 1. Features
### 1.1 Help
Run ```/do help``` and get help on all commands.

### 1.2 Available Commands
- ```/do help``` - Run 'test' to see if you're configured to run do commands
#### Authentication
- ```/do connect <access token>``` - Associates your DO team personal token with your mattermost account
- ```/do disconnect``` - Remove your DO team personal token with your mattermost account
- ```/do token``` - Provides instructions on getting a personal access token for the configured DigitalOcean team
- ```/do show-configured-token``` - Display your configured access token
#### Droplets
- ```/do list-droplets``` - List all Droplets in your team
- ```/do rename-droplet <dropletID> <name>``` - Rename a droplet
- ```/do reboot-droplet <dropletID>``` - Reboot a droplet
- ```/do shutdown-droplet <dropletID>``` - Shutdown a droplet
- ```/do powercycle-droplet <dropletID>``` - action is similar to pushing the reset button on a physical machine, it's similar to booting from scratch
#### Domains
- ```/do list-domains``` - Retrieve a list of all of the domains in your team
#### SSH Keys
- ```/do list-keys``` - Retrieve a list of all of SSH keys in your team
- ```/do retrieve-key <keyID>``` - Retrieve a single key by its ID
- ```/do delete-key <keyID>``` - Delete single key by its ID
- ```/do create-key <name> <publicKey>``` - Add an SSH key to your team

## 2. Configuration
Configure the plugin in Mattermost by going to ```System Console > Plugins > DigitalOcean```. Enable the plugin if it's not enabled.

## 3. Development
- Fork this repo
- Clone your fork and make changes on your branch
- Run ```$ make``` at the root of this project
- Install the generated tar on your server to see your changes

NB: This will add images to the github CDN. We will not merge this PR.