import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';

type Props = {
    name: string;
    onChangeFunc: (event: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>) => void;
    value: string;
    placeholder?: string;
    largeText?: boolean;
}

export default class FormButton extends PureComponent<Props, {}> {
    public static propTypes = {
        name: PropTypes.string.isRequired,
        onChangeFunc: PropTypes.func.isRequired,
        value: PropTypes.string.isRequired,
        placeholder: PropTypes.string,
        largeText: PropTypes.bool,
    }

    public static defaultProps = {
        placeholder: '',
        largeText: false,
    }

    public render(): JSX.Element {
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
                    rows={5}
                    value={value}
                    onChange={onChangeFunc}
                />
            );
        }

        return textInput;
    }
}
