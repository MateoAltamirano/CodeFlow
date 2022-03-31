import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import { createPinia } from 'pinia';
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from 'axios';
import VueAxios from 'vue-axios';

const app = createApp(App);

app.use(ElementPlus);
app.use(VueAxios, axios);
app.provide('axios', app.config.globalProperties.axios);
app.use(createPinia());
app.use(router);

app.mount('#app');
