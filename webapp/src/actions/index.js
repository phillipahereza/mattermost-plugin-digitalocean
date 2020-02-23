/* eslint-disable no-unused-vars */
import ActionTypes from '../action_types';

import Client, {doFetch} from '../client';
import {getPluginServerRoute} from '../selectors';

export const openCreateModal = () => (dispatch) => dispatch({
    type: ActionTypes.OPEN_CREATE_DROPLET_MODAL,
});

export const closeCreateModal = () => {
    return {
        type: ActionTypes.CLOSE_CREATE_DROPLET_MODAL,
    };
};

export const getTeamRegions = () => {
    return async (dispatch, getState) => {
        let data;
        const baseUrl = getPluginServerRoute(getState());
        try {
            data = await doFetch(`${baseUrl}/api/v1/get-do-regions`, {
                method: 'get',
            });
        } catch (error) {
            return {error};
        }

        dispatch({
            type: ActionTypes.RECEIVED_DO_REGIONS,
            data,
        });

        return {data};
    };
};

export const getDropletSizes = () => {
    return async (dispatch, getState) => {
        let data;
        const baseUrl = getPluginServerRoute(getState());
        try {
            data = await doFetch(`${baseUrl}/api/v1/get-do-sizes`, {
                method: 'get',
            });
        } catch (error) {
            return {error};
        }

        dispatch({
            type: ActionTypes.RECEIVED_DO_DROPLET_SIZES,
            data,
        });

        return {data};
    };
};

export const getImages = () => {
    return async (dispatch, getState) => {
        let data;
        const baseUrl = getPluginServerRoute(getState());
        try {
            data = await doFetch(`${baseUrl}/api/v1/get-do-images`, {
                method: 'get',
            });
        } catch (error) {
            return {error};
        }

        dispatch({
            type: ActionTypes.RECEIVED_DO_IMAGES,
            data,
        });

        return {data};
    };
};

export const createDroplet = (droplet) => {
    return async (dispatch, getState) => {
        const baseUrl = getPluginServerRoute(getState());
        try {
            const data = await doFetch(`${baseUrl}/api/v1/create-droplet`, {
                method: 'post',
                body: JSON.stringify(droplet),
            });
            return {data};
        } catch (error) {
            return {error};
        }
    };
};