import {id} from '../manifest';
import {prepareRegionsSelectData, prepareSizeSelectData, prepareImageSelectData} from '../utils';

const getPluginState = (state) => state['plugins-' + id] || {};

export const isCreateModalOpen = (state) => getPluginState(state).openModal;

export const getRegions = (state) => getPluginState(state).regions;

export const getPreparedRegions = (state) => prepareRegionsSelectData(getRegions(state));

export const getDropletSizes = (state) => getPluginState(state).sizes;

export const getPreparedSizes = (state) => prepareSizeSelectData(getDropletSizes(state));

export const getImages = (state) => getPluginState(state).images;

export const getPreparedImages = (state) => prepareImageSelectData(getImages(state));
