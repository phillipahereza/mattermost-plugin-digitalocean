/* eslint-disable import/no-unresolved */
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {closeCreateModal} from '../../actions';
import {isCreateModalOpen} from '../../selectors';

import CreateDroplet from './create_droplet';

const mapStateToProps = (state) => {
    return {
        show: isCreateModalOpen(state),
    };
};

const mapDispatchToProps = (dispatch) => bindActionCreators({
    closeCreateModal,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(CreateDroplet);
