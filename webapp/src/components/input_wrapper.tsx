
import React from 'react';
import PropTypes from 'prop-types';

type Props = {
    helpText?: string;
    label?: string;
    inputId?: string;
    hideRequiredStar?: boolean;
    required?: boolean;
}

export default class Wrapper extends React.PureComponent<Props, {}> {
    public static propTypes = {
        inputId: PropTypes.string,
        label: PropTypes.string,
        children: PropTypes.node.isRequired,
        helpText: PropTypes.string,
        required: PropTypes.bool,
        hideRequiredStar: PropTypes.bool,
    };

    public render(): JSX.Element {
        const {
            children,
            helpText,
            inputId,
            label,
            required,
            hideRequiredStar,
        } = this.props;

        return (
            <div className='form-group less'>
                {label &&
                <label
                    className='control-label margin-bottom x2'
                    htmlFor={inputId}
                >
                    {label}
                </label>
                }
                {required && !hideRequiredStar &&
                <span
                    className='error-text'
                    style={{marginLeft: '3px'}}
                >
                    {'*'}
                </span>
                }
                <div>
                    {children}
                    <div className='help-text'>
                        {helpText}
                    </div>
                </div>
            </div>
        );
    }
}
