import axios from 'axios';
import {ClientError} from 'mattermost-redux/client/client4';

class Client {
    constructor() {
        this.baseURL = '/plugins/com.mattermost.digitalocean';
        this.axiosInstance = axios.create({
            baseURL: this.baseURL,
        });
    }

    getDOTeamRegions = async () => {
        return this.doGet('/api/v1/get-do-regions');
    }

    getDOTeamDropletSizes = async () => {
        return this.doGet('/api/v1/get-do-sizes');
    }

    getDOTeamImages = async () => {
        return this.doGet('/api/v1/get-do-images');
    }

    createDroplet = async (droplet) => {
        return this.doPost('/api/v1/create-droplet', droplet);
    }

    doGet = async (url) => {
        try {
            const response = await this.axiosInstance.get(url);
            return response.data;
        } catch (error) {
            throw new ClientError(this.baseURL, {
                message: error.response.data || '',
                status_code: error.response.status,
                url,
            });
        }
    }

    doPost = async (url, data) => {
        try {
            const response = await this.axiosInstance.post(url, data);
            return response.data;
        } catch (error) {
            throw new ClientError(this.baseURL, {
                message: error.response.data || '',
                status_code: error.response.status,
                url,
            });
        }
    }
}

const client = new Client();

export default client;