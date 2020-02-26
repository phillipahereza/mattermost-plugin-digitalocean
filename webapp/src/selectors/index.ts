import {getConfig} from 'mattermost-redux/selectors/entities/general';
import {GlobalState} from 'mattermost-redux/types/store';

import {id} from '../manifest';
import {prepareRegionsSelectData, prepareSizeSelectData, prepareImageSelectData} from '../utils';
import {PluginState, GenericSelectData} from '../ts_types';

const getPluginState = (state: GlobalState): PluginState => state['plugins-' + id] || {};

export const getPluginServerRoute = (state: GlobalState): string => {
    const config = getConfig(state);

    let basePath = '/';
    if (config && config.SiteURL) {
        basePath = new URL(config.SiteURL).pathname;

        if (basePath && basePath[basePath.length - 1] === '/') {
            basePath = basePath.substr(0, basePath.length - 1);
        }
    }

    return basePath + '/plugins/' + id;
};

export const isCreateModalOpen = (state: GlobalState): boolean => getPluginState(state).openModal;

export const getRegions = (state: GlobalState): any[] => getPluginState(state).regions;

export const getPreparedRegions = (state: GlobalState): GenericSelectData[] => prepareRegionsSelectData(getRegions(state));

export const getDropletSizes = (state: GlobalState): any[] => getPluginState(state).sizes;

export const getPreparedSizes = (state: GlobalState): GenericSelectData[] => prepareSizeSelectData(getDropletSizes(state));

export const getImages = (state: GlobalState): any[] => getPluginState(state).images;

export const getPreparedImages = (state: GlobalState): GenericSelectData[] => prepareImageSelectData(getImages(state));
