import ActionTypes from '../action_types';

export const openCreateModal = () => (dispatch) => dispatch({
    type: ActionTypes.OPEN_CREATE_DROPLET_MODAL,
});

export const closeCreateModal = () => {
    return {
        type: ActionTypes.CLOSE_CREATE_DROPLET_MODAL,
    };
};