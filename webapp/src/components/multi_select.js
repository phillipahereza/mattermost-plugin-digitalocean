/* eslint-disable react/jsx-filename-extension */
/* eslint-disable react/prop-types */
import React, {PureComponent} from 'react';
import Select from 'react-select';
import CreatableSelect from 'react-select/creatable';

export default class MultiSelect extends PureComponent {
    render() {
        const {creatable, options, name} = this.props;
        let selectComponent;
        selectComponent = (
            <Select
                name={name}
                isSearchable={true}
                options={options}
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
                />
            );
        }
        return selectComponent;
    }
}
