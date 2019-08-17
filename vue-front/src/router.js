import Vue from 'vue'
import Router from 'vue-router'
import ArticleCreate from './views/ArticleCreate.vue'
import ArticleList from './views/ArticleList.vue'
import JiroList from './views/JiroList.vue'
import About from './views/About.vue'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'about',
            component: About,
            props: true
        },
        {
            path: '/articles/list',
            name: 'articleList',
            component: ArticleList,
            props: true
        },
        {
            path: '/articles/create',
            name: 'articleCreate',
            component: ArticleCreate,
            props: true
        },
        {
            path: '/jiros/list',
            name: 'jiroList',
            component: JiroList,
            props: true
        },
    ]
})