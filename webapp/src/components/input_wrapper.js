/* eslint-disable react/jsx-filename-extension */
import React from 'react';
import PropTypes from 'prop-types';

export default class Wrapper extends React.PureComponent {
    static propTypes = {
        inputId: PropTypes.string,
        label: PropTypes.node,
        children: PropTypes.node.isRequired,
        helpText: PropTypes.node,
        required: PropTypes.bool,
        hideRequiredStar: PropTypes.bool,
    };

    render() {
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
