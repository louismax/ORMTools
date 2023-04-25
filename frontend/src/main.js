import {createApp} from 'vue'
import App from './App.vue'
import { store } from './store/index.js'
import hljs from 'highlight.js/lib/core';
import javascript from 'highlight.js/lib/languages/javascript';
import golang from 'highlight.js/lib/languages/go';
import hljsVuePlugin from "@highlightjs/vue-plugin";
hljs.registerLanguage('javascript', javascript);
hljs.registerLanguage('golang', golang);

const app = createApp(App)
app.use(store)
app.use(hljsVuePlugin)
app.mount('#app')
