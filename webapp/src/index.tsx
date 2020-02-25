import {PluginRegistry} from 'mattermost-webapp/plugins/registry';
import {Store, Action} from 'redux';

import manifest from './manifest';

import reducers from './reducers';

import UIHook from './ui_hooks';
import CreateDropletModal from './components/create_droplet/index';

const w = window as any;

export default class Plugin {
    public initialize(registry: PluginRegistry, store: Store<object, Action<object>>): void {
        registry.registerReducer(reducers);
        registry.registerRootComponent(CreateDropletModal);

        const hooks = new UIHook(store);
        registry.registerSlashCommandWillBePostedHook(hooks.slashCommandWillBePostedHook);
    }
}

w.registerPlugin(manifest.id, new Plugin());
