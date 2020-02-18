/* eslint-disable react/prop-types */
/* eslint-disable react/jsx-filename-extension */
import React from 'react';
import {Modal} from 'react-bootstrap';

export default class CreateDropletModal extends React.PureComponent {
    render() {
        const {show, closeCreateModal} = this.props;
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
                        {'Create Droplet'}
                    </Modal.Title>
                </Modal.Header>
                <form
                    role='form'
                >
                    <Modal.Body >
                        <div>{'Create'}</div>
                    </Modal.Body>
                    <Modal.Footer>
                        <div>{'Footer'}</div>
                    </Modal.Footer>
                </form>
            </Modal>
        );
    }
}