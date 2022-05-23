<template>
<div class="h-screen">
  <nav class="flex items-center justify-around bg-teal-500 p-6">
  <div class="flex items-center flex-shrink-0 text-white mr-6">
    <svg class="fill-current h-8 w-8 mr-2" width="54" height="54" viewBox="0 0 54 54" xmlns="http://www.w3.org/2000/svg"><path d="M13.5 22.1c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05zM0 38.3c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05z"/></svg>
    <span class="font-semibold text-xl tracking-tight"><NuxtLink to="/">Home</NuxtLink></span>
  </div>
  <div class=" block flex lg:flex lg:items-center">

    <div v-if="!auth">
      <NuxtLink to="/login" class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Login</NuxtLink>
      <NuxtLink to="/register" class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Register</NuxtLink>
    </div>

    <div v-else>
      <a @click="logout" class="cursor-pointer inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Logout</a>  
    </div>
  </div>
</nav>
<nuxt />
</div>
</template>

<script>
export default {
    name: 'layout',
    data(){
        return {
            auth: false
        }
    },
    mounted() {
        this.$nuxt.$on('auth',auth => {
            console.log(auth)
            this.auth = auth;
        })
    },
    methods:{
        async logout(){
            await fetch('http://localhost:8000/api/logout',{
                method:'POST',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            this.auth = false;
            await this.$router.push('/login');
        }
    }
}
</script>

<style>

</style>