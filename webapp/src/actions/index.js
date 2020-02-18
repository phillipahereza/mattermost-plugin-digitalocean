import {OPEN_CREATE_DROPLET_MODAL, CLOSE_CREATE_DROPLET_MODAL} from '../action_types';

export const openCreateModal = () => {
    return {
        type: OPEN_CREATE_DROPLET_MODAL,
    };
};

export const closeCreateModal = () => {
    return {
        type: CLOSE_CREATE_DROPLET_MODAL,
    };
};