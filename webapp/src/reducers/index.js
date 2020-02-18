import {combineReducers} from 'redux';

import {OPEN_CREATE_DROPLET_MODAL, CLOSE_CREATE_DROPLET_MODAL} from '../action_types';

const openModal = (state = false, action) => {
    switch (action.type) {
    case OPEN_CREATE_DROPLET_MODAL:
        return true;
    case CLOSE_CREATE_DROPLET_MODAL:
        return false;
    default:
        return state;
    }
};

export default combineReducers({
    openModal,
});