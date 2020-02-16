package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/phillipahereza/mattermost-plugin-digitalocean/server/client"
	"strings"
	"text/tabwriter"
	"time"
)

func (p *Plugin) listKubernetesClustersCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	clusters, response, err := client.ListKubernetesClusters(context.TODO(), nil)

	if err != nil {
		p.API.LogError("failed to fetch kubernetes clusters", "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching kubernetes clusters list"),
			&model.AppError{Message: err.Error()}
	}

	if len(clusters) == 0 {
		return p.responsef(args, "You don't have any Kubernetes clusters"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "ID", "Name", "Region", "Endpoint", "Created At")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "------", "----", "------", "------", "------")

	for _, cluster := range clusters {

		fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", cluster.ID, cluster.Name, cluster.RegionSlug, cluster.Endpoint, cluster.CreatedAt.Format(time.RFC822))
	}

	w.Flush()
	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) listKubernetesClusterNodePoolsCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs, id string) (*model.CommandResponse, *model.AppError) {

	nodePools, response, err := client.ListKubernetesClusterNodePools(context.TODO(), id, nil)

	if err != nil {
		p.API.LogError("failed to fetch kubernetes cluster", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching kubernetes clusters %s", id),
			&model.AppError{Message: err.Error()}
	}

	if len(nodePools) == 0 {
		return p.responsef(args, "This cluster has no node pools"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "ID", "Name", "Size", "Node Count", "Tags")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "------", "----", "------", "------", "------")

	for _, pool := range nodePools {

		fmt.Fprintf(w, "\n |%s|%s|%s|%d|%s|", pool.ID, pool.Name, pool.Size, pool.Count, strings.Join(pool.Tags, ", "))
	}

	w.Flush()
	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) listKubernetesClusterNodesCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs, id string) (*model.CommandResponse, *model.AppError) {

	nodes, response, err := client.ListKubernetesClusterNodes(context.TODO(), id, nil)

	if err != nil {
		p.API.LogError("failed to fetch kubernetes cluster", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching kubernetes clusters %s", id),
			&model.AppError{Message: err.Error()}
	}

	if len(nodes) == 0 {
		return p.responsef(args, "This cluster has no nodes"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "ID", "Name", "Status", "Node Pool", "Created At")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", "------", "----", "------", "------", "------")

	for _, node := range nodes {
		fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|", node.ID, node.Name, node.Status.State, node.NodePoolName, node.CreatedAt.Format(time.RFC822))
	}

	w.Flush()
	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) retrieveAvailableUpgradesForKubernetesClusterCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs, id string) (*model.CommandResponse, *model.AppError) {

	upgrades, response, err := client.GetKubernetesClusterUpgrades(context.TODO(), id)
	if err != nil {
		p.API.LogError("failed to fetch upgrades for kubernetes cluster", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching upgrades for kubernetes cluster %s", id),
			&model.AppError{Message: err.Error()}
	}

	if len(upgrades) == 0 {
		return p.responsef(args, "This cluster has no available upgrades"), nil
	}

	upgradeList := ""
	for i, upgrade := range upgrades {
		upgradeList += fmt.Sprintf("%d. Kubernetes Version: %s", i+1, upgrade.Slug)
	}
	return p.responsef(args, upgradeList), nil
}

func (p *Plugin) upgradeKubernetesClusterCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs, id, versionSlug string) (*model.CommandResponse, *model.AppError) {

	response, err := client.UpgradeKubernetesCluster(context.TODO(), id, &godo.KubernetesClusterUpgradeRequest{VersionSlug: versionSlug})
	if err != nil {
		p.API.LogError("failed to upgrade for kubernetes cluster", "id", id, "version", versionSlug, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while attempting to upgrade kubernetes cluster %s to %s because %s", id, versionSlug, err.Error()),
			&model.AppError{Message: err.Error()}
	}
	return p.responsef(args, "Successfully upgrades Kubernetes Cluster %s to version %s", id, versionSlug), nil
}

func (p *Plugin) retrieveKubeconfigCommandFunc(client *client.DigitalOceanClient, args *model.CommandArgs, id string) (*model.CommandResponse, *model.AppError) {

	kubeconfig, response, err := client.GetKubeConfig(context.TODO(), id)

	if err != nil {
		p.API.LogError("failed to get kubeconfig for kubernetes cluster", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while attempting to get kubeconfig for kubernetes cluster %s because %s", id, err.Error()),
			&model.AppError{Message: err.Error()}
	}
	yamlString := fmt.Sprintf("```YAML\n%s\n```", string(kubeconfig.KubeconfigYAML))
	return p.responsef(args, yamlString), nil

}
