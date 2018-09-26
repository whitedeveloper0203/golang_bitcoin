<template>
  <div class="profile-page">
    <br><br><br><br>
    <h3>Your Email: {{this.response.Email}}</h3>
    <h3>Your demo balance: {{this.response.DemoBalance}}</h3>
  </div>
</template>

<script>
  import { HTTP } from '../http-constants'
  import router from '../router'

  export default {
    name: 'trade',
    props: ['auth', 'authenticated'],
    data () {
      return {
        auth0UUID: '',
        response: ''
      }
    },
    methods: {
      // Get Users Profile Info
      getUserInfo: function () {
        console.log('PROFILE getUserInfo')
        HTTP.get(`/api/userinfo/${this.auth0UUID}`)
          .then(response => {
            this.response = response.data
          })
          .catch(e => {
            console.log('err: ' + e)
          })
      }
    },
    mounted () {
      // TODO MAKE SURE THESE MATCH
      this.auth0UUID = localStorage.getItem('auth0UUID')
      if (this.auth0UUID) {
        this.getUserInfo()
      }
    },
    beforeMount () {
      if (!this.authenticated) {
        console.log('PROFILE(beforemount): sorry, log in')
        router.replace('/home')
      } else {
        console.log('PROFILE(beforemount): authenticated:', this.authenticated)
      }
    }
  }
</script>

<style lang="scss" scoped>
	.profile-page {
		background-color: rgb(5, 69, 71);
		height: 100vh;
		padding: 20px;
	}
	h3 {
		color: white;
	}
	.deposit {
		a,
		i {
			color: white !important;
		}
	}
	.deposit:hover {
		a,
		i {
			background-color: white;
			color: #1E84E0 !important;
		}
	}
	a {
		cursor: pointer;
	}
</style>

