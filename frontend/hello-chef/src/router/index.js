import { createRouter, createWebHistory } from 'vue-router';

// Import your components
import UserDashboard from '../views/UserDashboard.vue';
import MyList from '../views/MyList.vue';
import GenerateRecipe from '../views/GenerateRecipe.vue';

// Define your routes
const routes = [
  {
    path: '/',
    name: 'UserDashboard',
    component: UserDashboard
  },
  {
    path: '/list',
    name: 'MyList',
    component: MyList
  },
  {
    path: '/generate',
    name: 'GenerateRecipe',
    component: GenerateRecipe
  }
];

// Create the router instance and pass the `routes` option
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;