// This file is automatically generated. Do not modify it manually.

const manifest = JSON.parse(`
{
    "id": "com.mattermost.digitalocean",
    "name": "DigitalOcean Plugin",
    "description": "A DigitalOcean plugin for Mattermost",
    "version": "0.1.0",
    "min_server_version": "5.12.0",
    "server": {
        "executables": {
            "linux-amd64": "server/dist/plugin-linux-amd64",
            "darwin-amd64": "server/dist/plugin-darwin-amd64",
            "windows-amd64": "server/dist/plugin-windows-amd64.exe"
        },
        "executable": ""
    },
    "webapp": {
        "bundle_path": "webapp/dist/main.js"
    },
    "settings_schema": {
        "header": "Mattermost plugin for DigitalOcean Teams.",
        "footer": "",
        "settings": [
            {
                "key": "DOTeamID",
                "display_name": "Unique DigitalOcean Team Identifier",
                "type": "text",
                "help_text": "",
                "placeholder": "",
                "default": null
            },
            {
                "key": "DOAdmins",
                "display_name": "DO Admins",
                "type": "text",
                "help_text": "Users that are not system admins on Mattermost but have advanced plugin privileges",
                "placeholder": "",
                "default": null
            },
            {
                "key": "IMAPServer",
                "display_name": "IMAP Server Address",
                "type": "text",
                "help_text": "Enable imap on your email provider. Example of Google's imap server is 'imap.gmail.com:993'. 933 is usually the default port",
                "placeholder": "",
                "default": ""
            },
            {
                "key": "IMAPUsername",
                "display_name": "IMAP Username",
                "type": "text",
                "help_text": "This is the same as your Digital Ocean alerts email. Usually the email address. Something like example@gmail.com.",
                "placeholder": "",
                "default": ""
            },
            {
                "key": "IMAPPassword",
                "display_name": "IMAP Password",
                "type": "text",
                "help_text": "",
                "placeholder": "",
                "default": ""
            },
            {
                "key": "CronConfig",
                "display_name": "Updates schedule",
                "type": "dropdown",
                "help_text": "Define how often to check for Monitoring updates and alert channels. Default is 5 minutes.",
                "placeholder": "",
                "default": "*/5 * * * *",
                "options": [
                    {
                        "display_name": "Demo, every minute",
                        "value": "*/1 * * * *"
                    },
                    {
                        "display_name": "Every 5 minutes",
                        "value": "*/5 * * * *"
                    },
                    {
                        "display_name": "Every 10 minutes",
                        "value": "*/10 * * * *"
                    },
                    {
                        "display_name": "Every 30 minutes",
                        "value": "*/30 * * * *"
                    },
                    {
                        "display_name": "Every 1 hour",
                        "value": "*/60 * * * *"
                    }
                ]
            }
        ]
    }
}
`);

export default manifest;
export const id = manifest.id;
export const version = manifest.version;
