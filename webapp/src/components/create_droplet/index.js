import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import CreateDroplet from './create_droplet.js';

const mapStateToProps = (state) => {
    console.log(state);
    return {};
};

const mapDispatchToProps = (dispatch) => bindActionCreators({}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(CreateDroplet);
