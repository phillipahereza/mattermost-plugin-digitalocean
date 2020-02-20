/* eslint-disable camelcase */
/* eslint-disable no-magic-numbers */
/* eslint-disable react/jsx-filename-extension */
import React from 'react';
import {Modal} from 'react-bootstrap';
import PropTypes from 'prop-types';

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

Note.propTypes = {
    Tag: PropTypes.string.isRequired,
};

const initialState = {
    saving: false,
    name: '',
    user_data: '',
    backups: false,
    ipV6: false,
    private_networking: false,
    monitoring: false,
};

export default class CreateDropletModal extends React.PureComponent {
    static propTypes = {
        closeCreateModal: PropTypes.func.isRequired,
        getDropletSizes: PropTypes.func.isRequired,
        getImages: PropTypes.func.isRequired,
        getTeamRegions: PropTypes.func.isRequired,
        theme: PropTypes.object.isRequired,
        show: PropTypes.bool.isRequired,
        createDroplet: PropTypes.func.isRequired,
        regionsSelectData: PropTypes.array.isRequired,
        sizeSelectData: PropTypes.array.isRequired,
        imageSelectData: PropTypes.array.isRequired,
    }

    constructor(props) {
        super(props);
        this.state = initialState;
    }

    onTextInputChange = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        });
    }

    onMultiSelectChange = (inputValue, name) => {
        this.setState({
            [name]: inputValue,
        });
    }

    onToggleChange = (event) => {
        this.setState({
            [event.target.name]: !this.state[event.target.name],
        });
    }

    handleCloseModal = () => {
        const {closeCreateModal} = this.props;
        this.setState(initialState, closeCreateModal);
    }

    componentDidMount() {
        const {getTeamRegions, getDropletSizes, getImages} = this.props;
        getTeamRegions();
        getDropletSizes();
        getImages();
    }

    prepareFormMultiKeys = (keys) => {
        if (keys.length === 0) {
            return [];
        }

        const prepKeys = [];
        keys.forEach((key) => {
            prepKeys.push(key.value);
        });
        return prepKeys;
    }

    createDropletDataFromState = () => {
        const {
            name, region, size,
            image, ssh_keys, backups,
            ipV6, private_networking, user_data,
            monitoring, volumes, tags} = this.state;

        const droplet = {};
        droplet.name = name;
        droplet.region = region.value;
        droplet.size = size.value;
        droplet.image = image.value;
        droplet.ssh_keys = this.prepareFormMultiKeys(ssh_keys);
        droplet.backups = backups;
        droplet.ipV6 = ipV6;
        droplet.private_networking = private_networking;
        droplet.user_data = user_data;
        droplet.monitoring = monitoring;
        droplet.volumes = this.prepareFormMultiKeys(volumes);
        droplet.tags = this.prepareFormMultiKeys(tags);

        return droplet;
    }

    handleCreate = (e) => {
        if (e && e.preventDefault) {
            e.preventDefault();
        }

        const {createDroplet, closeCreateModal} = this.props;
        const droplet = this.createDropletDataFromState();

        this.setState({saving: true});

        createDroplet(droplet).then((data) => {
            if (data.error) {
                this.setState({saving: false});
                return;
            }

            closeCreateModal();
        });
    }

    render() {
        const {
            show,
            regionsSelectData,
            sizeSelectData,
            imageSelectData,
        } = this.props;

        const {saving, name, user_data, backups, monitoring, private_networking, ipV6} = this.state;
        const footer = (
            <React.Fragment>
                <FormButton
                    type='button'
                    btnClass='btn-link'
                    defaultMessage='Cancel'
                    onClick={this.handleCloseModal}
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
                    >
                        <Checkbox
                            checked={backups}
                            name='backups'
                            onChange={this.onToggleChange}
                        />
                        {'backups'}
                    </Note>
                    <Note
                        Tag='label'
                        style={{marginLeft: '1em', ...noteStyle}}
                    >
                        <Checkbox
                            checked={ipV6}
                            name='ipV6'
                            onChange={this.onToggleChange}
                        />
                        {'ipV6'}
                    </Note>
                    <Note
                        Tag='label'
                        style={{marginLeft: '1em', ...noteStyle}}
                    >
                        <Checkbox
                            checked={private_networking}
                            name='private_networking'
                            onChange={this.onToggleChange}
                        />
                        {'private_networking'}
                    </Note>
                    <Note
                        Tag='label'
                        style={{marginLeft: '1em', ...noteStyle}}
                    >
                        <Checkbox
                            checked={monitoring}
                            name='monitoring'
                            onChange={this.onToggleChange}
                        />
                        {'Monitoring'}
                    </Note>
                </InputWrapper>
                <InputWrapper
                    label='Select region'
                    required={true}
                >
                    <MultiSelect
                        name='region'
                        options={regionsSelectData}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Select droplet size'
                    required={true}
                >
                    <MultiSelect
                        name='size'
                        options={sizeSelectData}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Select droplet image'
                    required={true}
                >
                    <MultiSelect
                        name='image'
                        options={imageSelectData}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Add SSH keys'
                    required={false}
                >
                    <MultiSelect
                        name='ssh_keys'
                        creatable={true}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Add tags'
                    required={false}
                >
                    <MultiSelect
                        name='tags'
                        creatable={true}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
                    />
                </InputWrapper>
                <InputWrapper
                    label='Add volumes'
                    required={false}
                >
                    <MultiSelect
                        name='volumes'
                        creatable={true}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
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
                onHide={this.handleCloseModal}
                onExited={this.handleCloseModal}
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
                    onSubmit={this.handleCreate}
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