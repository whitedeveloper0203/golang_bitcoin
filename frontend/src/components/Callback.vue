<template>
  <div class="spinner">
    <img src="../assets/loading.svg" alt="loading"/>
  </div>
</template>

<script>
  import router from '../router'
  export default {
    name: 'callback',
    props: ['auth', 'authenticated'],
    data () {
      this.auth.handleAuthentication()
      return {}
    },
    created () {
      if (this.authenticated) {
        console.log('CALLBACK(created): I dont think so...')
        router.replace('/home')
      }
    }
  }
</script>

<style>
  .spinner {
  	background-color: white;
  	bottom: 0;
  	display: flex;
  	height: 100vh;
  	justify-content: center;
  	left: 0;
  	position: absolute;
  	right: 0;
  	top: 0;
  	width: 100vw;
  }
</style>
