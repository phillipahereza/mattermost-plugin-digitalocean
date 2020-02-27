/* eslint-disable @typescript-eslint/camelcase */
/* eslint-disable no-magic-numbers */
import React from 'react';
import {Modal} from 'react-bootstrap';
import PropTypes from 'prop-types';

import {GenericAction, ActionFunc} from 'mattermost-redux/types/actions';

import FormButton from '../form_button';

import {prepareSizeSelectData} from '../../utils';

import InputWrapper from '../input_wrapper';
import TextInput from '../text_input';
import MultiSelect from '../multi_select';

import {GenericSelectData, Droplet} from '../../ts_types';

const Checkbox = (props: {checked: boolean; name: string; onChange: (event: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>) => void}): JSX.Element => (
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

const Note = ({Tag = 'div', ...props}): JSX.Element => (
    <Tag
        {...props}
    />
);

Note.propTypes = {
    Tag: PropTypes.string.isRequired,
};

type Props = {
    closeCreateModal: () => GenericAction;
    regions: any[];
    getDropletSizes: () => ActionFunc;
    getImages: () => ActionFunc;
    getTeamRegions: () => ActionFunc;
    theme: object;
    show: boolean;
    createDroplet: (droplet: Droplet) => ActionFunc;
    regionsSelectData: GenericSelectData[];
    imageSelectData: GenericSelectData[];
}

type State = {
    rawSizes: any[];
    saving: boolean;
    name: string;
    user_data: string;
    backups: boolean;
    ipV6: boolean;
    private_networking: boolean;
    monitoring: boolean;
    size: GenericSelectData | null;
    region?: GenericSelectData;
    image?: any;
    ssh_keys?: any;
    volumes?: any;
    tags?: any;
}

const initialState: State = {

    // Sizes based on region to select from
    rawSizes: [],
    saving: false,
    name: '',
    user_data: '',
    backups: false,
    ipV6: false,
    private_networking: false,
    monitoring: false,
    size: null,
};

export default class CreateDropletModal extends React.PureComponent<Props, State> {
    public static propTypes = {
        closeCreateModal: PropTypes.func.isRequired,
        regions: PropTypes.array.isRequired,
        getDropletSizes: PropTypes.func.isRequired,
        getImages: PropTypes.func.isRequired,
        getTeamRegions: PropTypes.func.isRequired,
        theme: PropTypes.object.isRequired,
        show: PropTypes.bool.isRequired,
        createDroplet: PropTypes.func.isRequired,
        regionsSelectData: PropTypes.array.isRequired,
        imageSelectData: PropTypes.array.isRequired,
    }

    public constructor(props: Props) {
        super(props);
        this.state = initialState;
    }

    public onTextInputChange = (event: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>): void => {
        this.setState({
            [event.target.name]: event.target.value,
        }as any);
    }

    public onMultiSelectChange = (inputValue: GenericSelectData, name: string): void => {
        this.setState({
            [name]: inputValue,
        } as any);

        if (name === 'region') {
            this.loadSizesBasedOnRegion(inputValue.value);
        }
    }

    public loadSizesBasedOnRegion = (regionSlug: string | undefined): void => {
        const {regions} = this.props;
        const selectedRegion = regions.filter((region): boolean => region.slug === regionSlug);
        const sizes = prepareSizeSelectData(selectedRegion[0].sizes);
        this.setState({rawSizes: sizes, size: null});
    }

    public onToggleChange = (event: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>): void => {
        this.setState({
            [event.target.name]: !this.state[event.target.name],
        } as any);
    }

    public handleCloseModal = (e: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>): void => {
        if (e && e.preventDefault) {
            e.preventDefault();
        }

        const {closeCreateModal} = this.props;
        this.setState(initialState);
        closeCreateModal();
    }

    public componentDidMount(): void {
        const {getTeamRegions, getDropletSizes, getImages} = this.props;
        getTeamRegions();
        getDropletSizes();
        getImages();
    }

    public prepareFormMultiKeys = (keys: any[], name: string): any[] => {
        if (typeof keys === 'undefined') {
            return [];
        }

        if (keys.length === 0) {
            return [];
        }

        const prepKeys = [];

        if (name === 'ssh') {
            keys.forEach((key): void => {
                prepKeys.push({Fingerprint: key.value});
            });
            return prepKeys;
        }

        if (name === 'tags') {
            keys.forEach((key): void => {
                prepKeys.push(key.value);
            });
            return prepKeys;
        }

        keys.forEach((key): void => {
            prepKeys.push({ID: key.value});
        });
        return prepKeys;
    }

    public createDropletDataFromState = (): Droplet => {
        const {
            name, region, size,
            image, ssh_keys, backups,
            ipV6, private_networking, user_data,
            monitoring, volumes, tags} = this.state;

        const droplet: Droplet = {name: '', region: '', size: '', image: {ID: ''}};
        droplet.name = name;
        droplet.region = region.value;
        droplet.size = size.value;
        droplet.image = {ID: image.value};
        droplet.ssh_keys = this.prepareFormMultiKeys(ssh_keys, 'ssh');
        droplet.backups = backups;
        droplet.ipV6 = ipV6;
        droplet.private_networking = private_networking;
        droplet.user_data = user_data;
        droplet.monitoring = monitoring;
        droplet.volumes = this.prepareFormMultiKeys(volumes, '');
        droplet.tags = this.prepareFormMultiKeys(tags, 'tags');

        return droplet;
    }

    public handleCreate = (e: React.ChangeEvent<HTMLTextAreaElement> | React.ChangeEvent<HTMLInputElement>): void => {
        if (e && e.preventDefault) {
            e.preventDefault();
        }

        const {createDroplet} = this.props;
        const droplet = this.createDropletDataFromState();

        this.setState({saving: true});

        createDroplet(droplet).then((data: any): unknown => {
            if (data.error) {
                this.setState({saving: false});
                return;
            }

            this.handleCloseModal(e);
        });
    }

    public render(): JSX.Element {
        const {
            show,
            regionsSelectData,
            imageSelectData,
        } = this.props;

        const {saving, name, user_data, backups, monitoring, private_networking, ipV6, size} = this.state;
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
                        options={this.state.rawSizes}
                        handleSelectChange={this.onMultiSelectChange}
                        theme={this.props.theme}
                        selectedValue={size}
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