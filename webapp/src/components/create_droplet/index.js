/* eslint-disable import/no-unresolved */
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {closeCreateModal, getTeamRegions, getDropletSizes, getImages, createDroplet} from '../../actions';
import {isCreateModalOpen, getPreparedRegions, getPreparedSizes, getPreparedImages} from '../../selectors';

import CreateDroplet from './create_droplet';

const mapStateToProps = (state) => {
    return {
        show: isCreateModalOpen(state),
        regionsSelectData: getPreparedRegions(state),
        sizeSelectData: getPreparedSizes(state),
        imageSelectData: getPreparedImages(state),
    };
};

const mapDispatchToProps = (dispatch) => bindActionCreators({
    closeCreateModal,
    getTeamRegions,
    getDropletSizes,
    createDroplet,
    getImages,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(CreateDroplet);
