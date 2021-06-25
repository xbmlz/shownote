import Multipane from './src/multipane.vue';
import MultipaneResizer from './src/multipane-resizer.vue';

export { Multipane, MultipaneResizer };

if (typeof window !== 'undefined' && window.Vue) {
    window.Vue.component('multipane', Multipane);
    window.Vue.component('multipane-resizer', MultipaneResizer);
}