/* eslint-disable react/jsx-filename-extension */
/* eslint-disable react/prop-types */
import React, {PureComponent} from 'react';

export default class FormButton extends PureComponent {
    render() {
        const {name, onChangeFunc, value, placeholder, largeText} = this.props;
        let textInput;
        textInput = (
            <input
                name={name}
                className='form-control'
                type='text'
                placeholder={placeholder}
                value={value}
                onChange={onChangeFunc}
            />
        );

        if (largeText) {
            textInput = (
                <textarea
                    name={name}
                    className='form-control'
                    rows='5'
                    value={value}
                    onChange={onChangeFunc}
                />
            );
        }

        return textInput;
    }
}
