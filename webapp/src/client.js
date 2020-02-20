/* eslint-disable no-useless-catch */
import axios from 'axios';

class Client {
    constructor() {
        this.axiosInstance = axios.create({
            baseURL: '/plugins/com.mattermost.digitalocean',
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
            throw error;
        }
    }

    doPost = async (url, data) => {
        try {
            const response = await this.axiosInstance.post(url, data);
            return response.data;
        } catch (error) {
            throw error;
        }
    }
}

const client = new Client();

export default client;