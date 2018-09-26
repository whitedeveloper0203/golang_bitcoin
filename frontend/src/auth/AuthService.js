import auth0 from 'auth0-js'
import { AUTH_CONFIG } from './auth0-variables'
import EventEmitter from 'eventemitter3'
import router from './../router'
import { HTTP } from '../http-constants'

export default class AuthService {
  authenticated = this.isAuthenticated()
  authNotifier = new EventEmitter()
  auth0UUID = ''
  userEmail = ''
  userEmailVerified = false
  userFirstTime = false

  constructor () {
    this.login = this.login.bind(this)
    this.setSession = this.setSession.bind(this)
    this.logout = this.logout.bind(this)
    this.isAuthenticated = this.isAuthenticated.bind(this)
  }

  auth0 = new auth0.WebAuth({
    domain: AUTH_CONFIG.domain,
    clientID: AUTH_CONFIG.clientId,
    redirectUri: AUTH_CONFIG.callbackUrl,
    audience: `https://${AUTH_CONFIG.domain}/userinfo`,
    loginAfterSignUp: false,
    responseType: 'token id_token',
    scope: 'openid email'
  })

  // Off to auth0
  login () {
    this.auth0.authorize()
  }

  // Called from callback.vue after auth0.authorize redirect
  handleAuthentication () {
    this.auth0.parseHash((err, authResult) => {
      if (authResult && authResult.accessToken && authResult.idToken) {
        // Get info from session token
        // to send to the server
        this.authResult = authResult
        this.auth0UUID = authResult.idTokenPayload['https://myapp.example.com/uuid']
        this.userEmail = authResult.idTokenPayload.email
        this.userEmailVerified = authResult.idTokenPayload.email_verified
        // Check if it's the users first time logging in after email is verified
        if (authResult.idTokenPayload['https://myapp.example.com/logins_count'] === 2) {
          this.userFirstTime = true
        }
        if (this.userFirstTime === true) {
          this.notifyDB()
        }
        // If successful, set the session and redirect to trade view.
        this.setSession(authResult)
        router.replace('trade')
        // If signing up, let the user know they need to verify their email now.
      } else if (err && err.error === 'unauthorized') {
        router.replace('home')
        console.log(err)
        // alert(`Error: ${err.error}. VERIFY YO EMAIL FIRST`)
        // Any other error just logs user out and redirects to homepage
      } else if (err && err.error !== 'unauthorized') {
        this.logout()
        console.log(err)
        // alert(`Error: ${err.error}. YOU'RE BEING SIGNED OUT`)
      }
    })
  }

  notifyDB () {
    // Register user with custom (cross-application) uuid associated with user's auth0 info and there email / email status.
    HTTP.post('api/userinfo', {
      'auth0uuid': this.auth0UUID,
      'email': this.userEmail,
      'email_verified': this.userEmailVerified,
      'demo_balance': 100
    })
      .then(function (response) {
        console.log(response)
      })
      .catch(function (error) {
        this.errors = error
        console.log(error)
      })
  }

  setSession (authResult) {
    // Set the time that the access token will expire at
    let expiresAt = JSON.stringify(
      authResult.expiresIn * 1000 + new Date().getTime()
    )
    localStorage.setItem('access_token', authResult.accessToken)
    localStorage.setItem('id_token', authResult.idToken)
    localStorage.setItem('expires_at', expiresAt)
    localStorage.setItem('auth0UUID', this.auth0UUID)
    this.authNotifier.emit('authChange', { authenticated: true })
  }

  logout () {
    // Clear access token and ID token from local storage
    localStorage.removeItem('access_token')
    localStorage.removeItem('id_token')
    localStorage.removeItem('expires_at')
    localStorage.removeItem('auth0UUID')
    this.userProfile = null
    this.authNotifier.emit('authChange', false)
    // navigate to the home route
    router.replace('home')
  }

  isAuthenticated () {
    // Check whether the current time is past the
    // access token's expiry time
    let expiresAt = JSON.parse(localStorage.getItem('expires_at'))
    return new Date().getTime() < expiresAt
  }
}
