/* eslint-disable react/jsx-filename-extension */
/* eslint-disable react/prop-types */
import React, {PureComponent} from 'react';
import AsyncSelect from 'react-select/async';

export default class MultiSelect extends PureComponent {
    render() {
        const {loadOptionsFunc} = this.props;
        return (
            <AsyncSelect
                isMulti={true}
                cacheOptions={true}
                defaultOptions={true}
                loadOptions={loadOptionsFunc}
            />
        );
    }
}
