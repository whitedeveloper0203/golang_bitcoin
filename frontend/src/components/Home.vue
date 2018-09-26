<template>
  <div>
    <h4 v-if="authenticated">
        You are logged in!
    </h4>
    <h4 v-if="!authenticated">
      You are NOT logged in!
    </h4>
  </div>
</template>

<script>
  export default {
    name: 'home',
    props: ['auth', 'authenticated']
  }
</script>

<style scoped>
  h4 {
  	margin-top: 300px;
  }
  h4,
  a {
    text-align: center;
  	color: black !important;
  	cursor: pointer;
  }
</style>

