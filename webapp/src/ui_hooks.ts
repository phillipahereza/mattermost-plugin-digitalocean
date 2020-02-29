import {Store, Action} from 'redux';

import {openCreateModal} from './actions';

export default class UIHook {
    public store: Store<object, Action<object>>;

    public constructor(store: Store<object, Action<object>>) {
        this.store = store;
    }

    public slashCommandWillBePostedHook = (message: string, contextArgs: any): Promise<{message?: string; args?: any}> => {
        let messageTrimmed;
        if (message) {
            messageTrimmed = message.trim();
        }

        if (messageTrimmed && messageTrimmed === '/do create-droplet') {
            this.store.dispatch(openCreateModal());
            return Promise.resolve({});
        }

        return Promise.resolve({message, args: contextArgs});
    }
}