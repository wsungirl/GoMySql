import Vue from 'vue'
import Router from 'vue-router'
import Auth from '@/components/Auth'
import DBList from '@/components/DBList'
import TablesList from '@/components/TablesList'
import Table from '@/components/Table'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/auth',
      name: 'auth',
      component: Auth
    },
    {
      path: '/db/:db_id/tables',
      name: 'tableslist',
      meta: { auth: true },
      component: TablesList
    },
    {
      path: '/db/:db_id/tables/:table_id/rows',
      name: 'table',
      meta: { auth: true },
      component: Table
    },
    {
      path: '/',
      name: 'dblist',
      meta: { auth: true },
      component: DBList
    }
  ]
})

router.beforeEach(function (to, from, next) {
  Vue.nextTick(function () {
    if (to.name === 'auth' && Auth.user.authenticated) next('/')
    if (to.meta.auth && !Auth.user.authenticated) next('/auth')
    else next()
  })
})

export default router
