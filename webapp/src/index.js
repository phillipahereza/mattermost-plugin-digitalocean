import manifest from './manifest';
import reducers from './reducers';
import UIHook from './ui_hooks';
import CreateDropletModal from './components/create_droplet/index';

export default class Plugin {
    initialize(registry, store) {
        registry.registerReducer(reducers);
        registry.registerRootComponent(CreateDropletModal);

        const hooks = new UIHook(store);
        registry.registerSlashCommandWillBePostedHook(hooks.slashCommandWillBePostedHook);
    }
}

window.registerPlugin(manifest.id, new Plugin());
