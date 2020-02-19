/* eslint-disable camelcase */
/* eslint-disable no-magic-numbers */
/* eslint-disable react/prop-types */
/* eslint-disable react/jsx-filename-extension */
import React from 'react';
import {Modal} from 'react-bootstrap';

import FormButton from '../form_button';
import InputWrapper from '../input_wrapper';
import TextInput from '../text_input';
import MultiSelect from '../multi_select';

const Checkbox = (props) => (
    <input
        type='checkbox'
        {...props}
    />);

const noteStyle = {
    color: 'hsl(0, 0%, 40%)',
    display: 'inline-block',
    fontSize: 12,
    fontStyle: 'italic',
    marginTop: '1em',
};

const Note = ({Tag = 'div', ...props}) => (
    <Tag
        {...props}
    />
);

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

    componentDidMount() {
        const {getTeamRegions, getDropletSizes, getImages} = this.props;
        getTeamRegions();
        getDropletSizes();
        getImages();
    }

    render() {
        const {
            show,
            closeCreateModal,
            regionsSelectData,
            sizeSelectData,
            imageSelectData,
        } = this.props;
        const {saving, name, user_data} = this.state;
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
                    <Note
                        Tag='label'
                        style={noteStyle}
                    ><Checkbox checked={true}/>{'backups'}</Note>
                    <Note
                        Tag='label'
                        style={{marginLeft: '1em', ...noteStyle}}
                    ><Checkbox checked={true}/>{'ipV6'}</Note>
                    <Note
                        Tag='label'
                        style={{marginLeft: '1em', ...noteStyle}}
                    ><Checkbox checked={true}/>{'private_networking'}</Note>
                    <Note
                        Tag='label'
                        style={{marginLeft: '1em', ...noteStyle}}
                    ><Checkbox checked={true}/>{'Monitoring'}</Note>
                </InputWrapper>
                <InputWrapper
                    label='Select region'
                    required={true}
                >
                    <MultiSelect
                        name='region'
                        options={regionsSelectData}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Select droplet size'
                    required={true}
                >
                    <MultiSelect
                        name='size'
                        options={sizeSelectData}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Select droplet image'
                    required={true}
                >
                    <MultiSelect
                        name='image'
                        options={imageSelectData}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Add SSH keys'
                    required={false}
                >
                    <MultiSelect
                        name='keys'
                        creatable={true}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Add tags'
                    required={false}
                >
                    <MultiSelect
                        name='tags'
                        creatable={true}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Add volumnes'
                    required={false}
                >
                    <MultiSelect
                        name='volumes'
                        creatable={true}
                    />
                </InputWrapper>
                <InputWrapper
                    label='User data'
                >
                    <TextInput
                        onChangeFunc={this.onTextInputChange}
                        name='user_data'
                        placeholder=''
                        value={user_data}
                        largeText={true}
                    />
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