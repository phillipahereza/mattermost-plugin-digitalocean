package client

import (
	"context"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

//go:generate mockgen -destination=../mocks/mock_client.go -package=mocks github.com/phillipahereza/mattermost-plugin-digitalocean/server/client DigitalOceanService

// DigitalOceanClient manages communication with DigitalOcean V2 API.
type DigitalOceanClient struct {
	Client *godo.Client
}

// K8sNode represents a node in a Kubernetes cluster with the node pool name
type K8sNode struct {
	*godo.KubernetesNode
	NodePoolName string
}

// TokenSource contains the access token used for authentication
type TokenSource struct {
	AccessToken string
}

// Token returns the access token
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// DigitalOceanService is an interface for interfacing with the DigitalOcean API endpoints
type DigitalOceanService interface {
	CreateDatabaseUser(ctx context.Context, userName string, databaseCreateUserRequest *godo.DatabaseCreateUserRequest) (*godo.DatabaseUser, *godo.Response, error)
	DeleteDatabaseClusterUser(ctx context.Context, clusterID string, userName string) (*godo.Response, error)
	ListDatabaseClusterUsers(ctx context.Context, clusterID string, listOptions *godo.ListOptions) ([]godo.DatabaseUser, *godo.Response, error)
	ListDatabasesInCluster(ctx context.Context, clusterID string, listOptions *godo.ListOptions) ([]godo.DatabaseDB, *godo.Response, error)
	ListDatabaseClusters(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Database, *godo.Response, error)
	ListDatabaseClusterBackups(ctx context.Context, clusterID string, listOptions *godo.ListOptions) ([]godo.DatabaseBackup, *godo.Response, error)

	ListDomains(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Domain, *godo.Response, error)

	CreateDroplet(ctx context.Context, createRequest *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error)
	ListDroplets(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Droplet, *godo.Response, error)
	PowerCycleDroplet(ctx context.Context, dropletID int) (*godo.Action, *godo.Response, error)
	RebootDroplet(ctx context.Context, dropletID int) (*godo.Action, *godo.Response, error)
	RenameDroplet(ctx context.Context, dropletID int, newName string) (*godo.Action, *godo.Response, error)
	ShutdownDroplet(ctx context.Context, dropletID int) (*godo.Action, *godo.Response, error)

	CreateSSHKey(ctx context.Context, keyCreateRequest *godo.KeyCreateRequest) (*godo.Key, *godo.Response, error)
	DeleteSSHKeyByID(ctx context.Context, keyID int) (*godo.Response, error)
	ListSSHKeys(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Key, *godo.Response, error)
	GetSSHKeyByID(ctx context.Context, keyID int) (*godo.Key, *godo.Response, error)

	ListKubernetesClusters(ctx context.Context, listOptions *godo.ListOptions) ([]*godo.KubernetesCluster, *godo.Response, error)
	ListKubernetesClusterNodePools(ctx context.Context, clusterID string, opts *godo.ListOptions) ([]*godo.KubernetesNodePool, *godo.Response, error)
	ListKubernetesClusterNodes(ctx context.Context, clusterID string, opts *godo.ListOptions) ([]*K8sNode, *godo.Response, error)
	GetKubernetesClusterUpgrades(ctx context.Context, clusterID string) ([]*godo.KubernetesVersion, *godo.Response, error)
	GetKubeConfig(ctx context.Context, clusterID string) (*godo.KubernetesClusterConfig, *godo.Response, error)
	UpgradeKubernetesCluster(ctx context.Context, clusterID string, upgradeRequest *godo.KubernetesClusterUpgradeRequest) (*godo.Response, error)
}

// CreateDatabaseUser will create a new database user
func (do *DigitalOceanClient) CreateDatabaseUser(ctx context.Context, name string, databaseCreateUserRequest *godo.DatabaseCreateUserRequest) (*godo.DatabaseUser, *godo.Response, error) {
	return do.Client.Databases.CreateUser(ctx, name, databaseCreateUserRequest)
}

// DeleteDatabaseClusterUser will delete an existing database user
func (do *DigitalOceanClient) DeleteDatabaseClusterUser(ctx context.Context, clusterID string, name string) (*godo.Response, error) {
	return do.Client.Databases.DeleteUser(ctx, clusterID, name)
}

// ListDatabaseClusterUsers returns all database users for the database
func (do *DigitalOceanClient) ListDatabaseClusterUsers(ctx context.Context, clusterID string, listOptions *godo.ListOptions) ([]godo.DatabaseUser, *godo.Response, error) {
	return do.Client.Databases.ListUsers(ctx, clusterID, listOptions)
}

// ListDatabasesInCluster returns a lists of databases in a cluster
func (do *DigitalOceanClient) ListDatabasesInCluster(ctx context.Context, clusterID string, opts *godo.ListOptions) ([]godo.DatabaseDB, *godo.Response, error) {
	return do.Client.Databases.ListDBs(ctx, clusterID, opts)
}

// ListDatabaseClusters returns a list of the Database Clusters visible with the caller's API token
func (do *DigitalOceanClient) ListDatabaseClusters(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Database, *godo.Response, error) {
	return do.Client.Databases.List(ctx, listOptions)
}

// ListDatabaseClusterBackups returns a list of the current backups of a database
func (do *DigitalOceanClient) ListDatabaseClusterBackups(ctx context.Context, clusterID string, listOptions *godo.ListOptions) ([]godo.DatabaseBackup, *godo.Response, error) {
	return do.Client.Databases.ListBackups(ctx, clusterID, listOptions)
}

// ListDomains lists all domains.
func (do *DigitalOceanClient) ListDomains(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return do.Client.Domains.List(ctx, listOptions)
}

// CreateDroplet creates a new droplet.
func (do *DigitalOceanClient) CreateDroplet(ctx context.Context, createRequest *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error) {
	return do.Client.Droplets.Create(ctx, createRequest)
}

// ListDroplets lists all droplets.
func (do *DigitalOceanClient) ListDroplets(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
	return do.Client.Droplets.List(ctx, listOptions)
}

// PowerCycleDroplet powers off a droplet then powers it back on
func (do *DigitalOceanClient) PowerCycleDroplet(ctx context.Context, dropletID int) (*godo.Action, *godo.Response, error) {
	return do.Client.DropletActions.PowerCycle(ctx, dropletID)
}

// RebootDroplet reboots a droplet
func (do *DigitalOceanClient) RebootDroplet(ctx context.Context, dropletID int) (*godo.Action, *godo.Response, error) {
	return do.Client.DropletActions.Reboot(ctx, dropletID)
}

// RenameDroplet renames a droplet
func (do *DigitalOceanClient) RenameDroplet(ctx context.Context, dropletID int, newName string) (*godo.Action, *godo.Response, error) {
	return do.Client.DropletActions.Rename(ctx, dropletID, newName)
}

// ShutdownDroplet shuts down a droplet
func (do *DigitalOceanClient) ShutdownDroplet(ctx context.Context, dropletID int) (*godo.Action, *godo.Response, error) {
	return do.Client.DropletActions.Shutdown(ctx, dropletID)
}

// CreateSSHKey adds an SSH Key
func (do *DigitalOceanClient) CreateSSHKey(ctx context.Context, keyCreateRequest *godo.KeyCreateRequest) (*godo.Key, *godo.Response, error) {
	return do.Client.Keys.Create(ctx, keyCreateRequest)
}

// DeleteSSHKeyByID deletes an SSH key by its ID
func (do *DigitalOceanClient) DeleteSSHKeyByID(ctx context.Context, keyID int) (*godo.Response, error) {
	return do.Client.Keys.DeleteByID(ctx, keyID)
}

// ListSSHKeys lists all SSH keys
func (do *DigitalOceanClient) ListSSHKeys(ctx context.Context, listOptions *godo.ListOptions) ([]godo.Key, *godo.Response, error) {
	return do.Client.Keys.List(ctx, listOptions)
}

// GetSSHKeyByID retrieves an SSH key by ID
func (do *DigitalOceanClient) GetSSHKeyByID(ctx context.Context, keyID int) (*godo.Key, *godo.Response, error) {
	return do.Client.Keys.GetByID(ctx, keyID)
}

// ListKubernetesClusters lists all kubernetes clusters
func (do *DigitalOceanClient) ListKubernetesClusters(ctx context.Context, listOptions *godo.ListOptions) ([]*godo.KubernetesCluster, *godo.Response, error) {
	return do.Client.Kubernetes.List(ctx, listOptions)
}

// ListKubernetesClusterNodePools lists all NodePools in a Kubernetes cluster
func (do *DigitalOceanClient) ListKubernetesClusterNodePools(ctx context.Context, clusterID string, opts *godo.ListOptions) ([]*godo.KubernetesNodePool, *godo.Response, error) {
	return do.Client.Kubernetes.ListNodePools(ctx, clusterID, opts)
}

// ListKubernetesClusterNodes lists all nodes in a Kubernetes cluster
func (do *DigitalOceanClient) ListKubernetesClusterNodes(ctx context.Context, clusterID string, opts *godo.ListOptions) ([]*K8sNode, *godo.Response, error) {
	nodePools, response, err := do.ListKubernetesClusterNodePools(ctx, clusterID, opts)
	if err != nil {
		return nil, &godo.Response{}, err
	}

	var nodes []*K8sNode

	for _, nodePool := range nodePools {
		for _, node := range nodePool.Nodes {
			nodes = append(nodes, &K8sNode{node, nodePool.Name})
		}
	}

	return nodes, response, err
}

// GetKubernetesClusterUpgrades retrieves versions a Kubernetes cluster can be upgraded to
func (do *DigitalOceanClient) GetKubernetesClusterUpgrades(ctx context.Context, clusterID string) ([]*godo.KubernetesVersion, *godo.Response, error) {
	return do.Client.Kubernetes.GetUpgrades(ctx, clusterID)
}

// GetKubeConfig returns a Kubernetes config file for the specified cluster.
func (do *DigitalOceanClient) GetKubeConfig(ctx context.Context, clusterID string) (*godo.KubernetesClusterConfig, *godo.Response, error) {
	return do.Client.Kubernetes.GetKubeConfig(ctx, clusterID)
}

// UpgradeKubernetesCluster upgrades a Kubernetes cluster to a new version. Valid upgrade
// versions for a given cluster can be retrieved with `GetKubernetesClusterUpgrades`.
func (do *DigitalOceanClient) UpgradeKubernetesCluster(ctx context.Context, clusterID string, upgradeRequest *godo.KubernetesClusterUpgradeRequest) (*godo.Response, error) {
	return do.Client.Kubernetes.Upgrade(ctx, clusterID, upgradeRequest)
}
