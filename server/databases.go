package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"text/tabwriter"
	"time"
)

func (p *Plugin) listClustersFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	opts := &godo.ListOptions{}

	databases, response, err := client.Databases.List(context.TODO(), opts)

	if err != nil {
		p.API.LogError("failed to fetch databases", "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching databases list"),
			&model.AppError{Message: err.Error()}
	}

	if len(databases) == 0 {
		return p.responsef(args, "You don't have any databases configured"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|%s|", "ID", "Name", "Engine", "Status", "Size", "Region")
	fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|%s|", "------", "----", "----", "----", "----", "----")

	for _, db := range databases {

		fmt.Fprintf(w, "\n |%s|%s|%s|%s|%s|%s|", db.ID, db.Name, db.EngineSlug, db.Status, db.SizeSlug, db.RegionSlug)
	}

	w.Flush()

	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) listClusterBackupsFunc(args *model.CommandArgs, id string) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	backups, response, err := client.Databases.ListBackups(context.TODO(), id, nil)

	if err != nil {
		p.API.LogError("failed to get backups for database", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching backups list for database cluster %s", id),
			&model.AppError{Message: err.Error()}
	}

	if len(backups) == 0 {
		return p.responsef(args, "You don't have any cluster backups"), nil
	}

	buffer := new(bytes.Buffer)

	w := new(tabwriter.Writer)

	w.Init(buffer, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "\n |%s|%s|", "Created At", "Size (GB)")
	fmt.Fprintf(w, "\n |%s|%s|", "------", "----")

	for _, backup := range backups {

		fmt.Fprintf(w, "\n |%s|%f|", backup.CreatedAt.Format(time.RFC822), backup.SizeGigabytes)
	}

	w.Flush()

	return p.responsef(args, buffer.String()), nil
}

func (p *Plugin) addUserToClusterFunc(args *model.CommandArgs, id, name string) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}
	dbUserCreateReq := &godo.DatabaseCreateUserRequest{
		Name:          name,
		MySQLSettings: nil,
	}
	user, response, err := client.Databases.CreateUser(context.TODO(), id, dbUserCreateReq)
	if err != nil {
		p.API.LogError("failed to create user for database", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while creating a user on database %s because %s", id, err.Error()),
			&model.AppError{Message: err.Error()}
	}
	return p.responsef(args, "Name: `%s`\t Password: `%s`\t Role: `%s`", user.Name, user.Password, user.Role), nil
}

func (p *Plugin) listClusterUsersFunc(args *model.CommandArgs, id string) (*model.CommandResponse, *model.AppError) {
	client, err := p.GetClient(args.UserId)
	if err != nil {
		p.API.LogError("Failed to get digitalOcean client", "Err", err.Error())
		return p.responsef(args, "Failed to get DigitalOcean client"),
			&model.AppError{Message: err.Error()}
	}

	users, response, err := client.Databases.ListUsers(context.TODO(), id, nil)

	if err != nil {
		p.API.LogError("failed to get users for database", "id", id, "response", response, "Err", err.Error())
		return p.responsef(args, "Error while fetching users for database cluster %s", id),
			&model.AppError{Message: err.Error()}
	}

	if len(users) == 0 {
		return p.responsef(args, "You don't have any cluster backups"), nil
	}

	usersList := ""
	for _, user := range users {
		usersList += fmt.Sprintf("- Name: `%s`\t Role: `%s`\n", user.Name, user.Role)
	}
	return p.responsef(args, usersList), nil
}
