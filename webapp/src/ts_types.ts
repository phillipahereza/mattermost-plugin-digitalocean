export type Droplet = {
    name: string;
    region: string;
    size: string;
    image: Image;
    ssh_keys?: SSHKey[];
    backups?: boolean;
    ipV6?: boolean;
    private_networking?: boolean;
    user_data?: string;
    monitoring?: boolean;
    volumes?: Volume[];
    tags?: string[];
}

export type Image = {
    ID?: string;
    Slug?: string;
}

export type SSHKey = {
    ID?: string;
    Fingerprint?: string;
}

export type Volume = {
    ID?: string;
    Name?: string;
}

// Work on being more explicit with regions, sizes and image types
export type PluginState = {
    openModal: boolean;
    regions: any[];
    sizes: any[];
    images: any[];
}

export type GenericSelectData = {
    value?: string;
    label?: string;
}