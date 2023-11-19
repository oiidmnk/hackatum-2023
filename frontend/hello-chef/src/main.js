import { createApp } from 'vue';
import App from './App.vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faClock } from '@fortawesome/free-regular-svg-icons';
import { far } from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faPepperHot, faCarrot, faSeedling, faWheatAwn } from '@fortawesome/free-solid-svg-icons';
import router from './router';

library.add(faClock);
library.add(far);
library.add(faPepperHot, faCarrot, faSeedling, faWheatAwn);

const app = createApp(App);

app.component('font-awesome-icon', FontAwesomeIcon);

app.use(router).mount('#app');