import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';
import AlunoList from './components/AlunoList.vue';
import AlunoEdit from './components/AlunoEdit.vue';

const routes = [
    { path: '/', component: AlunoList },
    { path: '/aluno/:id/edit', component: AlunoEdit, props: true }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

const app = createApp(App);
app.use(router);
app.mount('#app');
