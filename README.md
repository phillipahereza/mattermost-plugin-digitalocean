# Mattermost Digital Ocean Plugin
[![phillipahereza](https://circleci.com/gh/phillipahereza/mattermost-plugin-digitalocean.svg?style=svg)](https://github.com/phillipahereza/mattermost-plugin-digitalocean)
[![Go Report Card](https://goreportcard.com/badge/github.com/phillipahereza/mattermost-plugin-digitalocean)](https://goreportcard.com/badge/github.com/phillipahereza/mattermost-plugin-digitalocean)

## Table of Contents
- [1. Configuration](#1-configuration)
- [2. Features](#2-features)
- [3. Development](#3-development)

## 1. Configuration
- Configure the plugin in Mattermost by going to ```System Console > Plugins > DigitalOcean```. Enable the plugin if it's not enabled.
- Go to Digital Ocean, copy and then add your Digital Ocean team identifier (ID) to the settings under *"Unique DigitalOcean Team Identifier"*

![Screen Shot 2020-02-16 at 09 51 09](https://user-images.githubusercontent.com/28563179/74600387-f8a8f300-50a1-11ea-99ee-b913c3d68fa7.png)

NB: This can be found as query param on any Digital Ocean Url. From this URL *https://cloud.digitalocean.com/projects/9ae5693a-1573-4dc1-a55d-3ebf87579XXX/resources?i=6dXcXX*, It would be *6dXcXX*
- Save your settings

## 2. Features
### 2.1 Help
Run ```/do help``` and get help on all commands.

### 2.2 Available Commands
- ```/do help``` - Run 'test' to see if you're configured to run do commands
#### Authentication
![token_instructions](https://user-images.githubusercontent.com/13383422/75652191-7efd3180-5c6b-11ea-826f-59ac26b93a0a.gif)
- ```/do connect <access token>``` - Associates your DO team personal token with your mattermost account
- ```/do disconnect``` - Remove your DO team personal token with your mattermost account
- ```/do token``` - Provides instructions on getting a personal access token for the configured DigitalOcean team
- ```/do show-configured-token``` - Display your configured access token

#### Droplets
- ```/do create-droplet``` - Easily create a new droplet from within Mattermost

[![drop](https://user-images.githubusercontent.com/28563179/75614604-888f7800-5b4b-11ea-8c9e-0222ce1b6eec.gif)](https://drive.google.com/file/d/1ccWofd3eUX5Mn61wsxzGqdglubm9vCqh/view?usp=sharing)


Create process alerts team members of the newly created resource

![MM SS](https://user-images.githubusercontent.com/28563179/75112995-8d8b8d80-565a-11ea-96a1-709f7b543ad1.png)


- ```/do list-droplets``` - List all Droplets in your team
- ```/do rename-droplet <dropletID> <name>``` - Rename a droplet
- ```/do reboot-droplet <dropletID>``` - Reboot a droplet
- ```/do shutdown-droplet <dropletID>``` - Shutdown a droplet
- ```/do powercycle-droplet <dropletID>``` - action is similar to pushing the reset button on a physical machine, it's similar to booting from scratch
#### Domains
![domains](https://user-images.githubusercontent.com/13383422/75648899-8966fd80-5c62-11ea-806d-cb47fbe1d469.gif)
- ```/do list-domains``` - Retrieve a list of all of the domains in your team
#### SSH Keys
![SSH](https://user-images.githubusercontent.com/13383422/75677559-cf40b780-5c9c-11ea-9d09-16e8678ed665.gif)
- ```/do list-keys``` - Retrieve a list of all of SSH keys in your team
- ```/do retrieve-key <keyID>``` - Retrieve a single key by its ID
- ```/do delete-key <keyID>``` - Delete single key by its ID
- ```/do create-key <name> <publicKey>``` - Add an SSH key to your team
#### Databases
- ```/do list-clusters``` - Retrieve a list of all Database Clusters set up in your team
- ```/do list-cluster-backups <id>``` - Retrieve a list of all backups of a Database Cluster
- ```/do add-cluster-user <clusterID> <userName>``` - Add a database user to a cluster
- ```/do list-cluster-users <clusterID>``` - List database cluster users
- ```/do delete-cluster-user <clusterID> <userName>``` - Delete a database user to a cluster
- ```/do list-cluster-dbs <clusterID>``` - List databases in the cluster
#### Kubernetes
## List Kubernetes clusters
![list-k8s](https://user-images.githubusercontent.com/13383422/75651335-3f354a80-5c69-11ea-922b-b94e987ee7ff.gif)

## List Kubernetes Cluster Nodes
![list-nodes](https://user-images.githubusercontent.com/13383422/75651365-4e1bfd00-5c69-11ea-94c6-08fe925c6685.gif)

## Get cluster kubeconfig 
![k8s-config](https://user-images.githubusercontent.com/13383422/75651382-5a07bf00-5c69-11ea-9856-2a28ff740278.gif)

- ```/do list-k8s-clusters``` - List all Kubernetes Clusters in your team
- ```/do list-k8s-cluster-nodepools <clusterID>``` - List Nodepools in a Kubernetes cluster
- ```/do list-k8s-cluster-nodes <clusterID>``` - List Nodes in a Kubernetes cluster
- ```/do list-k8s-cluster-upgrades <clusterID>``` - Retrieve a list of available upgrades for a Kubernetes cluster
- ```/do get-k8s-config <clusterID>``` - Retrieve kubeconfig file in YAML format
- ```/do upgrade-k8s-cluster <clusterID> <versionSlug>``` - Upgrade a Kubernetes cluster to a newer patch release of Kubernetes

### 2.3 Droplet updates
Plugin bot posts a message about the active status of droplets at configured intervals.
Running ```/do subcribe``` in a central channel will set it up to receive these regular updates from the bot.
Channels can unsubcribe by running ```/do unsubscribe```

![Screen Shot 2020-02-21 at 17 11 49](https://user-images.githubusercontent.com/28563179/75041362-7c614600-54cd-11ea-8611-741984efdf7d.png)

![Screen Shot 2020-02-21 at 17 21 26](https://user-images.githubusercontent.com/28563179/75041999-c39c0680-54ce-11ea-849f-9079fcb973aa.png)

![config](https://user-images.githubusercontent.com/13383422/75653372-85d97380-5c6e-11ea-984f-89443118dfa1.gif)

## 3. Development
- Fork this repo
- Clone your fork and make changes on your branch
- Run ```$ make``` at the root of this project
- Install the generated tar on your server to see your changes


## 4. TODO
- [ ] Create Kubernetes Clusters
- [ ] Scaling Kubernetes Clusters
- [ ] Configuring LoadBalancers
- [ ] Getting detailed Droplet healthcheck updates
- [ ] Working with different projects within an account
- [ ] Block Storage
- [ ] CDNs
- [ ] SSL Certificates

