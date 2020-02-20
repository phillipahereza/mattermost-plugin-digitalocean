import ActionTypes from '../action_types';
import Client from '../client';

export const openCreateModal = () => (dispatch) => dispatch({
    type: ActionTypes.OPEN_CREATE_DROPLET_MODAL,
});

export const closeCreateModal = () => {
    return {
        type: ActionTypes.CLOSE_CREATE_DROPLET_MODAL,
    };
};

export const getTeamRegions = () => {
    return async (dispatch) => {
        let data;
        try {
            data = await Client.getDOTeamRegions();
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
    return async (dispatch) => {
        let data;
        try {
            data = await Client.getDOTeamDropletSizes();
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
    return async (dispatch) => {
        let data;
        try {
            data = await Client.getDOTeamImages();
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
    return async (dispatch) => {
        let data;
        try {
            data = await Client.createDroplet(droplet);
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