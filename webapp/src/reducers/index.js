import {combineReducers} from 'redux';

import ActionTypes from '../action_types';

const openModal = (state = false, action) => {
    switch (action.type) {
    case ActionTypes.OPEN_CREATE_DROPLET_MODAL:
        return true;
    case ActionTypes.CLOSE_CREATE_DROPLET_MODAL:
        return false;
    default:
        return state;
    }
};

const regions = (state = [], action) => {
    switch (action.type) {
    case ActionTypes.RECEIVED_DO_REGIONS:
        return action.data;
    case ActionTypes.REQUEST_FAILED:
        return state;
    default:
        return state;
    }
};

const sizes = (state = [], action) => {
    switch (action.type) {
    case ActionTypes.RECEIVED_DO_DROPLET_SIZES:
        return action.data;
    case ActionTypes.REQUEST_FAILED:
        return state;
    default:
        return state;
    }
};

const images = (state = [], action) => {
    switch (action.type) {
    case ActionTypes.RECEIVED_DO_IMAGES:
        return action.data;
    case ActionTypes.REQUEST_FAILED:
        return state;
    default:
        return state;
    }
};

export default combineReducers({
    openModal,
    regions,
    sizes,
    images,
});