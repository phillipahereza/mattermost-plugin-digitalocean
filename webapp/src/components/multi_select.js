/* eslint-disable react/jsx-filename-extension */
/* eslint-disable react/prop-types */
import React, {PureComponent} from 'react';
import Select from 'react-select';
import CreatableSelect from 'react-select/creatable';

import {getStyleForReactSelect} from '../utils';

export default class MultiSelect extends PureComponent {
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
