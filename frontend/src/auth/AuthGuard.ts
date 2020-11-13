import { NavigationGuard } from 'vue-router'
import { getInstance } from '@/auth/Auth0Plugin'

export const AuthGuard: NavigationGuard = (to, from, next) => {
  const authService = getInstance()

  const fn = () => {
    unwatch && unwatch()

    if(to.fullPath == "/" && authService.isAuthenticated){
      return next("/dashboard")
    }
    if(to.fullPath == "/" && !authService.isAuthenticated){
      return next()
    }
    if (authService.isAuthenticated) {
      return next()
    }

    authService.loginWithRedirect({ appState: { targetUrl: to.fullPath } })
  }

  if (!authService.loading) {
    return fn()
  }

  const unwatch = authService.$watch('loading', (loading: boolean) => {
    if (loading === false) {
      return fn()
    }
  })
}
