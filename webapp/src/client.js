/* eslint-disable no-useless-catch */
import axios from 'axios';

class Client {
    constructor() {
        this.axiosInstance = axios.create({
            baseUrl: '/plugins/com.mattermost.digitalocean',
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