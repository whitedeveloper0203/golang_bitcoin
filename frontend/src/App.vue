<template>
  <div class="app-vue">

    <header id="header">
      <nav class="navbar navbar-expand-lg navbar-light nav-bg">
        <a href="/home" class="navbar-brand"><img src="./assets/images/logo.png" alt="Logo"></a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
        <i class="fas fa-bars"></i>
        </button>
        <div class="collapse navbar-collapse" id="navbarText">
          <ul class="navbar-nav mr-auto">
            <li class="nav-item default-item">
              <a v-if="authenticated" v-b-tooltip.hover title="Trade!" class="nav-link" href="/trade" ><i class="fas fa-th-large"></i></a>
            </li>
            <li class="nav-item default-item">
              <a v-if="authenticated" v-b-tooltip.hover title="Trade!" class="nav-link" href="/trade"><i class="fas fa-layer-group"></i></a>
            </li>
            <li class="nav-item default-item">
              <a v-if="authenticated" v-b-tooltip.hover title="Trade!" class="nav-link" href="/trade"><i class="fas fa-hand-holding-usd"></i></a>
            </li>

            <li v-if="!simpleNav" class="nav-item switch-item">
              <a v-b-tooltip.hover title="currently unavailable" class="switch_link">
                <div id="switches">
                  <div class="switch switch-yellow">

                    <input type="radio" class="switch-input" name="view2" value="week2" id="week2" checked>
                    <label for="week2" class="switch-label switch-label-off">BTC</label>
                    <input type="radio" class="switch-input" name="view2" value="month2" id="month2">
                    <label for="month2" class="switch-label switch-label-on" style="pointer-events: none;">ETH</label>
                    <span class="switch-selection"></span>

                    <i class="fas fa-exchange-alt"></i>
                  </div>
                </div>
              </a>
            </li>

            <li v-if="!simpleNav" class="nav-item switch-item">
              <a v-b-tooltip.hover title="currently unavailable"  class="switch_link">
                <div id="switches">
                  <div class="switch color-primary">
                    <input class="switch-input" name="view" value="week" id="week" type="radio">
                    <label for="week" class="switch-label switch-label-off" style="pointer-events: none;">REAL</label>
                    <input class="switch-input" name="view" value="month" id="month" checked  type="radio">
                    <label for="month" class="switch-label switch-label-on" style="pointer-events: none;"> DEMO </label>
                    <span class="switch-selection"></span>
                    <i class="fas fa-exchange-alt"></i>
                  </div>
                </div>
              </a>
            </li>

            <li v-if="!simpleNav" class="nav-item switch-item custom-switch-item">
              <a class="switch_link">
                <div id="switches">
                  <div class="switch switch-blue">
                    <label for="demo3" class="switch-label switch-label-off" style="pointer-events: none;">Balance:</label>
                    <label for="balance3" class="switch-label switch-label-on" style="pointer-events: none;">{{appDemoBalance}}</label>
                  </div>
                </div>
              </a>
            </li>

            <li v-if="!simpleNav" class="nav-item deposit">
              <a class="nav-link" href="#"><i class="fas fa-wallet"></i> Deposit</a>
            </li>
          </ul>

          <div class="right-nav">
            <ul class="navbar-nav mr-auto">

              <li class="nav-item">
                <span v-if="!authenticated" @click="login()" id="login-button" class="nav-link"><i class="far fa-user-circle"></i></span>
                <a v-else v-b-tooltip.hover title="Your Profile" class="nav-link" href="/profile"><i class="far fa-user-circle"></i></a>
              </li>

              <li class="nav-item">
                <a v-if="authenticated" v-b-tooltip.hover title="Notifications (none yet)" class="nav-link" href="/profile"><i class="far fa-bell"></i></a>
              </li>

              <li class="nav-item">
                <span v-if="authenticated" @click="auth.logout(), simplifyNav = true" v-b-tooltip.hover title="Logout" class="nav-link"><i class="fas fa-sign-out-alt"></i></span>
              </li>

            </ul>
            <b-tooltip v-if="!authenticated" show target="login-button" placement="left" offset="0, 20px">
              Log in / Sign up
            </b-tooltip>
          </div>
        </div>
      </nav>
    </header>

    <div>
      <router-view
        :auth="auth"
        :authenticated="authenticated"
        @updateBalanceFromTrade="updateBalanceFromTrade">
      </router-view>
    </div>
  </div>
</template>

<script>
import AuthService from './auth/AuthService'
const auth = new AuthService()
const { authenticated, authNotifier, login, logout } = auth

export default {
  name: 'app',
  data () {
    authNotifier.on('authChange', authState => {
      this.authenticated = authState.authenticated
    })
    return {
      auth0UUID: null,
      auth,
      authenticated,
      simplifyNav: true,
      appDemoBalance: 0
    }
  },
  computed: {
    simpleNav: function () {
      if (this.simplifyNav != null) {
        return this.simplifyNav
      }
    }
  },
  methods: {
    login,
    logout,
    updateBalanceFromTrade: function (newDemoBalance) { // UPDATE DEMOBALANCE WITH CUSTOM EVENT FROM TRADE PAGE.
      this.appDemoBalance = newDemoBalance
    }
  },
  mounted () {
    if (this.$route.path === '/home' || this.$route.path === '/profile' || this.$route.path === '/') {
      this.simplifyNav = true
    } else if (this.$route.path === '/trade' || this.$route.path === '/callback') {
      this.simplifyNav = false
    }
    // console.log('APP(mounted):', this.$route.path, 'simplenav:', this.simplifyNav, 'authenticated:', this.authenticated, this.appDemoBalance)
  },
  updated () {
    if (this.$route.path === '/home' || this.$route.path === '/profile' || this.$route.path === '/') {
      this.simplifyNav = true
    } else if (this.$route.path === '/trade' || this.$route.path === '/callback') {
      this.simplifyNav = false
    }
    // console.log('APP(updated):', this.$route.path, 'simplenav:', this.simplifyNav, 'authenticated:', this.authenticated, this.appDemoBalance)
  }
}
</script>
<style>
@import './assets/css/bootstrap.min.css';
@import './assets/css/style.css';
@import 'https://use.fontawesome.com/releases/v5.2.0/css/all.css';
</style>

<style lang="scss" scoped>
.navbar {
	#main-wrapper {
		background-image: '../assets/images/bj.jpg';
	}

	.app-vue {
		&__header {
			align-items: center;
			background-color: rgb(135, 162, 235);
			display: flex;
			justify-content: space-between;
			.logo,
			.demo-balance {
				margin: 0;
			}
		}
	}

	.nav-item {
		p {
			color: white;
			font-size: 12px;
			margin: 0;
		}
		.nav-link {
			&:hover {
				i {
					color: #01D014 !important;
				}
			}
		}
	}

	.custom-switch-item {
		padding: 3px 0;
	}

	.deposit {
		padding: 3px;
		a,
		i {
			color: white !important;
			margin-top: 0;
		}
		&:hover {
			a,
			i {
				background-color: white;
				color: #1E84E0 !important;
			}
		}
	}

	.section-head-span {
		text-transform: lowercase !important;
	}
	.trades-section,
	.order-pending,
	.order-buy {
		min-height: 320px;
	}
}
</style>