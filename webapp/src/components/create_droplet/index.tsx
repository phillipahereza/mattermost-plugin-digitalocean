import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {GlobalState} from 'mattermost-redux/types/store';
import {DispatchFunc} from 'mattermost-redux/types/actions';

import {closeCreateModal, getTeamRegions, getDropletSizes, getImages, createDroplet, sendSizesToGetDetails} from '../../actions';
import {isCreateModalOpen, getRegions, getPreparedRegions, getPreparedSizes, getPreparedImages} from '../../selectors';

import CreateDroplet from './create_droplet';

const mapStateToProps = (state: GlobalState): object => {
    return {
        show: isCreateModalOpen(state),
        regions: getRegions(state),
        regionsSelectData: getPreparedRegions(state),
        sizeSelectData: getPreparedSizes(state),
        imageSelectData: getPreparedImages(state),
    };
};

const mapDispatchToProps = (dispatch: DispatchFunc): object => bindActionCreators({
    closeCreateModal,
    getTeamRegions,
    getDropletSizes,
    createDroplet,
    getImages,
    sendSizesToGetDetails,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(CreateDroplet);
