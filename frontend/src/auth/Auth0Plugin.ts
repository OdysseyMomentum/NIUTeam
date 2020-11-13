import { Auth0Options, AuthWrapper, RedirectCallback } from '@/auth/AuthWrapper'
import { VueConstructor } from 'vue'

type Auth0PluginOptions = {
  onRedirectCallback: RedirectCallback;
  redirectUri: string;
  domain: string;
  clientId: string;
  audience?: string;
  [key: string]: string | RedirectCallback | undefined;
}

const DEFAULT_REDIRECT_CALLBACK = (appState: any) => window.history.replaceState(
  {}, document.title, window.location.pathname)

let instance: AuthWrapper

export const getInstance = () => instance

export const auth0 = ({
  onRedirectCallback = DEFAULT_REDIRECT_CALLBACK,
  redirectUri = window.location.origin,
  ...options
}) => {
  if (instance) return instance

  instance = new AuthWrapper()
  instance.init(onRedirectCallback,
    redirectUri,
    options as Auth0Options)

  return instance
}

export const Auth0Plugin = {
  install (Vue: VueConstructor, options: Auth0PluginOptions) {
    Vue.prototype.$auth = auth0(options)
  }
}
