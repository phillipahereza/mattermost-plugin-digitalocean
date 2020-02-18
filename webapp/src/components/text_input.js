/* eslint-disable react/jsx-filename-extension */
/* eslint-disable react/prop-types */
import React, {PureComponent} from 'react';

export default class FormButton extends PureComponent {
    render() {
        const {name, onChangeFunc, value, placeholder} = this.props;
        return (
            <input
                name={name}
                className='form-control'
                type='text'
                placeholder={placeholder}
                value={value}
                onChange={onChangeFunc}
            />
        );
    }
}
