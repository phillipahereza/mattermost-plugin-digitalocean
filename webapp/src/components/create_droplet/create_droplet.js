/* eslint-disable react/prop-types */
/* eslint-disable react/jsx-filename-extension */
import React from 'react';
import {Modal} from 'react-bootstrap';

import FormButton from '../form_button';
import InputWrapper from '../input_wrapper';
import TextInput from '../text_input';
import MultiSelect from '../multi_select';

export default class CreateDropletModal extends React.PureComponent {
    state = {
        saving: false,

        // Droplet name
        name: '',
    }

    onTextInputChange = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        });
    }

    render() {
        const {show, closeCreateModal} = this.props;
        const {saving, name} = this.state;
        const footer = (
            <React.Fragment>
                <FormButton
                    type='button'
                    btnClass='btn-link'
                    defaultMessage='Cancel'
                    onClick={() => closeCreateModal()}
                />
                <FormButton
                    id='submit-button'
                    type='submit'
                    btnClass='btn btn-primary'
                    saving={saving}
                >
                    {'Create'}
                </FormButton>
            </React.Fragment>
        );

        const formFields = (
            <>
                <InputWrapper
                    label='Droplet name'
                    required={true}
                >
                    <TextInput
                        onChangeFunc={this.onTextInputChange}
                        name='name'
                        placeholder='Droplet name'
                        value={name}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Select region'
                    required={true}
                >
                    <MultiSelect/>
                </InputWrapper>
            </>
        );
        return (
            <Modal
                dialogClassName='modal--scroll'
                show={show}
                onHide={() => closeCreateModal()}
                bsSize='large'
                backdrop='static'
            >
                <Modal.Header closeButton={true}>
                    <Modal.Title>
                        {'Create Digital Ocean Droplet'}
                    </Modal.Title>
                </Modal.Header>
                <form
                    role='form'
                >
                    <Modal.Body >
                        {formFields}
                    </Modal.Body>
                    <Modal.Footer>
                        {footer}
                    </Modal.Footer>
                </form>
            </Modal>
        );
    }
}