import {
    GenericAction,
    DispatchFunc,
    ActionFunc,
    GetStateFunc,
    ActionResult,
} from 'mattermost-redux/types/actions';

import ActionTypes from '../action_types';
import {Droplet} from '../ts_types';

import {doFetch} from '../client';
import {getPluginServerRoute} from '../selectors';

export const openCreateModal = (): DispatchFunc => (dispatch: DispatchFunc): DispatchFunc => dispatch({
    type: ActionTypes.OPEN_CREATE_DROPLET_MODAL,
});

export const closeCreateModal = (): GenericAction => {
    return {
        type: ActionTypes.CLOSE_CREATE_DROPLET_MODAL,
    };
};

export const getTeamRegions = (): ActionFunc => {
    return async (dispatch: DispatchFunc, getState: GetStateFunc): ActionResult => {
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

export const getDropletSizes = (): ActionFunc => {
    return async (dispatch: DispatchFunc, getState: GetStateFunc): ActionResult => {
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

export const getImages = (): ActionFunc => {
    return async (dispatch: DispatchFunc, getState: GetStateFunc): ActionResult => {
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

export const createDroplet = (droplet: Droplet): ActionFunc => {
    return async (dispatch: DispatchFunc, getState: GetStateFunc): ActionResult => {
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