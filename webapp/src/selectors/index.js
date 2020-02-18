import {id} from '../manifest';

const getPluginState = (state) => state['plugins-' + id] || {};

export const isCreateModalOpen = (state) => getPluginState(state).openModal;

export const getRegions = (state) => getPluginState(state).regions;

export const getDropletSizes = (state) => getPluginState(state).sizes;

export const getImages = (state) => getPluginState(state).images;
