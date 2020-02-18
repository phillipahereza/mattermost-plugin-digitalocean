import {id} from '../manifest';

const getPluginState = (state) => state['plugins-' + id] || {};

export const isCreateModalOpen = (state) => getPluginState(state).openModal;
