import React, {PureComponent} from 'react';
import Select from 'react-select';
import CreatableSelect from 'react-select/creatable';
import PropTypes from 'prop-types';

import {GenericSelectData} from '../ts_types';

import {getStyleForReactSelect} from '../utils';

type Props = {
    creatable?: boolean;
    options?: GenericSelectData[];
    name: string;
    handleSelectChange: (value: GenericSelectData, name: string) => void;
    theme: object;
    selectedValue?: GenericSelectData[];
}

export default class MultiSelect extends PureComponent<Props, {}> {
    public static propTypes = {
        creatable: PropTypes.bool,
        options: PropTypes.array,
        name: PropTypes.string.isRequired,
        handleSelectChange: PropTypes.func.isRequired,
        theme: PropTypes.object.isRequired,
        selectedValue: PropTypes.array,
    }

    public render(): JSX.Element {
        const {creatable, options, name, handleSelectChange, selectedValue} = this.props;
        let selectComponent;
        selectComponent = (
            <Select
                name={name}
                isSearchable={true}
                options={options}
                isClearable={true}
                onChange={(value: GenericSelectData | any): void => handleSelectChange(value, name)}
                styles={getStyleForReactSelect(this.props.theme)}
                value={selectedValue}
            />
        );

        if (creatable) {
            selectComponent = (
                <CreatableSelect
                    placeholder=''
                    menuPortalTarget={document.body}
                    menuPlacement='auto'
                    isClearable={true}
                    isMulti={true}
                    onChange={(value: GenericSelectData | any): void => handleSelectChange(value, name)}
                    styles={getStyleForReactSelect(this.props.theme)}
                />
            );
        }
        return selectComponent;
    }
}
