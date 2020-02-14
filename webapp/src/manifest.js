// This file is automatically generated. Do not modify it manually.

const manifest = JSON.parse(`
{
    "id": "com.mattermost.digitalocean",
    "name": "Digital Ocean Plugin",
    "description": "A digital ocean plugin for Mattermost",
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
        "header": "Mattermost plugin for Digital Ocean Teams.",
        "footer": "",
        "settings": [
            {
                "key": "DOTeamID",
                "display_name": "Unique DO Team Identifier",
                "type": "text",
                "help_text": "",
                "placeholder": "",
                "default": null
            },
            {
                "key": "DOAdmins",
                "display_name": "Users that are not system admins on Mattermost but have advanced plugin privileges",
                "type": "text",
                "help_text": "",
                "placeholder": "",
                "default": null
            }
        ]
    }
}
`);

export default manifest;
export const id = manifest.id;
export const version = manifest.version;
