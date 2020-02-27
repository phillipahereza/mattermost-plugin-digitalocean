import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';

type Props = {
    id?: string;
    executing?: boolean;
    disabled?: boolean;
    executingMessage?: React.ReactNode;
    defaultMessage: React.ReactNode;
    btnClass?: string;
    extraClasses?: string;
    saving?: boolean;
    savingMessage?: string;
    type?: 'button' | 'submit' | 'reset' | undefined;
    onClick?: (e: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>) => void;
}

export default class FormButton extends PureComponent<Props, {}> {
    public static propTypes = {
        executing: PropTypes.bool,
        disabled: PropTypes.bool,
        executingMessage: PropTypes.node,
        defaultMessage: PropTypes.node,
        btnClass: PropTypes.string,
        extraClasses: PropTypes.string,
        saving: PropTypes.bool,
        savingMessage: PropTypes.string,
        type: PropTypes.string,
    };

    public static defaultProps = {
        disabled: false,
        savingMessage: 'Creating',
        defaultMessage: 'Create',
        btnClass: 'btn-primary',
        extraClasses: '',
    };

    public render(): JSX.Element {
        const {saving, disabled, savingMessage, defaultMessage, btnClass, extraClasses, ...props} = this.props;

        let contents;
        if (saving) {
            contents = (
                <span>
                    <span
                        className='fa fa-spinner icon--rotate'
                        title={'Loading Icon'}
                    />
                    {savingMessage}
                </span>
            );
        } else {
            contents = defaultMessage;
        }

        let className = 'save-button btn ' + btnClass;

        if (extraClasses) {
            className += ' ' + extraClasses;
        }

        return (
            <button
                id='saveSetting'
                className={className}
                disabled={disabled}
                {...props}
            >
                {contents}
            </button>
        );
    }
}
