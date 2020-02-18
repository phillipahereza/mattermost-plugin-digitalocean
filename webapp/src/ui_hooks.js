import {openCreateModal} from './actions';

export default class UIHook {
    constructor(store) {
        this.store = store;
    }

    slashCommandWillBePostedHook = (message, contextArgs) => {
        let messageTrimmed;
        if (message) {
            messageTrimmed = message.trim();
        }

        if (messageTrimmed && messageTrimmed.startsWith('/do create')) {
            // Make checks like if user is connected

            this.store.dispatch(openCreateModal());
            return Promise.resolve({});
        }

        return Promise.resolve({message, args: contextArgs});
    }
}