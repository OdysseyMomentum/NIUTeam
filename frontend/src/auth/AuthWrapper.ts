import { Component, Vue } from 'vue-property-decorator'
import { User } from '@/auth/User'
import createAuth0Client, {
  Auth0Client, CacheLocation,
  GetIdTokenClaimsOptions,
  GetTokenSilentlyOptions, GetTokenWithPopupOptions, IdToken, LogoutOptions,
  PopupLoginOptions,
  RedirectLoginOptions
} from '@auth0/auth0-spa-js'

export type Auth0Options = {
  domain: string;
  clientId: string;
  audience: string;
  cacheLocation: CacheLocation;
  [key: string]: string | undefined;
}

// @ts-ignore
export type RedirectCallback = (appState) => void

@Component({})
export class AuthWrapper extends Vue {
  loading = true
  isAuthenticated? = false
  user?: User
  auth0Client?: Auth0Client
  popupOpen = false
  error?: Error

  async getUser () {
    return new User(await this.auth0Client?.getUser())
  }

  async loginWithPopup (popupLoginOptions: PopupLoginOptions) {
    this.popupOpen = true

    try {
      await this.auth0Client?.loginWithPopup(popupLoginOptions)
    } catch (e) {
      console.error(e)
      this.error = e
    } finally {
      this.popupOpen = false
    }

    this.user = await this.getUser()
    this.isAuthenticated = true
  }

  loginWithRedirect (redirectLoginOptions: RedirectLoginOptions): Promise<void> | undefined {
    return this.auth0Client?.loginWithRedirect(redirectLoginOptions)
  }

  getIdTokenClaims (getIdTokenClaimsOptions: GetIdTokenClaimsOptions): Promise<IdToken> | undefined {
    return this.auth0Client?.getIdTokenClaims(getIdTokenClaimsOptions)
  }

  getTokenSilently (getTokenSilentlyOptions: GetTokenSilentlyOptions): Promise<any> | undefined {
    return this.auth0Client?.getTokenSilently(getTokenSilentlyOptions)
  }

  getTokenWithPopup (getTokenWithPopupOptions: GetTokenWithPopupOptions): Promise<string> | undefined {
    return this.auth0Client?.getTokenWithPopup(getTokenWithPopupOptions)
  }

  logout (logoutOptions: LogoutOptions): void {
    return this.auth0Client?.logout(logoutOptions)
  }

  // We initialize all that shit with our own method (we cannot pass any parameters to the created() lifecycle hook).
  async init (onRedirectCallback: RedirectCallback,
    redirectUri: string,
    auth0Options: Auth0Options) {
    this.auth0Client = await createAuth0Client({
      domain: auth0Options.domain,
      client_id: auth0Options.clientId,
      audience: auth0Options.audience,
      cacheLocation: auth0Options.cacheLocation,
      redirect_uri: redirectUri
    })

    try {
      if (window.location.search.includes('error') ||
        (window.location.search.includes('code=')) &&
        window.location.search.includes('state=')) {
        const { appState } = await this.auth0Client?.handleRedirectCallback() ?? { appState: undefined }

        onRedirectCallback(appState)
      }
    } catch (e) {
      console.error(e)
      this.error = e
    } finally {
      this.isAuthenticated = await this.auth0Client?.isAuthenticated()
      this.user = await this.getUser()
      this.loading = false
    }
  }
}
