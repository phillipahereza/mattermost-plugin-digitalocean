/* eslint-disable react/jsx-filename-extension */
import React, {PureComponent} from 'react';
import Select from 'react-select';
import CreatableSelect from 'react-select/creatable';
import PropTypes from 'prop-types';

import {getStyleForReactSelect} from '../utils';

export default class MultiSelect extends PureComponent {
    static propTypes = {
        creatable: PropTypes.bool,
        options: PropTypes.array,
        name: PropTypes.string.isRequired,
        handleSelectChange: PropTypes.func.isRequired,
        theme: PropTypes.object.isRequired,
    }

    render() {
        const {creatable, options, name, handleSelectChange} = this.props;
        let selectComponent;
        selectComponent = (
            <Select
                name={name}
                isSearchable={true}
                options={options}
                onChange={(value) => handleSelectChange(value, name)}
                styles={getStyleForReactSelect(this.props.theme)}
            />
        );

        if (creatable) {
            selectComponent = (
                <CreatableSelect
                    noOptionsMessage={() => 'Start typing...'}
                    formatCreateLabel={(value) => `Add "${value}"`}
                    placeholder=''
                    menuPortalTarget={document.body}
                    menuPlacement='auto'
                    isClearable={true}
                    isMulti={true}
                    onChange={(value) => handleSelectChange(value, name)}
                    styles={getStyleForReactSelect(this.props.theme)}
                />
            );
        }
        return selectComponent;
    }
}
