/* eslint-disable react/jsx-filename-extension */
import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';

export default class FormButton extends PureComponent {
    static propTypes = {
        name: PropTypes.string.isRequired,
        onChangeFunc: PropTypes.func.isRequired,
        value: PropTypes.string.isRequired,
        placeholder: PropTypes.string,
        largeText: PropTypes.bool,
    }

    static defaultProps = {
        placeholder: '',
        largeText: false,
    }

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
