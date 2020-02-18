import React from 'react';
import {Modal} from 'react-bootstrap';

export default class CreateDropletModal extends React.PureComponent {
    render() {
        return (
            <Modal
                dialogClassName='modal--scroll'
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