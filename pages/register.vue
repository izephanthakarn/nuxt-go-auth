<template>
  <div class="mt-7 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-96 space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">Register</h2>
      </div>
      <form @submit.prevent="submit" class="mt-8 space-y-6" action="#" method="POST">
        <div class="rounded-md -space-y-px">
          <div>
            <label for="username" class="sr-only">Username</label>
            <input v-model="name" id="username" name="username" type="text" autocomplete="username" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-teal-500 focus:border-teal-500 focus:z-10 sm:text-sm" placeholder="Username" />
          </div>
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input v-model="email" id="email-address" name="email" type="email" autocomplete="email" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-teal-500 focus:border-teal-500 focus:z-10 sm:text-sm" placeholder="Email address" />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input v-model="password" id="password" name="password" type="password" autocomplete="current-password" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-teal-500 focus:border-teal-500 focus:z-10 sm:text-sm" placeholder="Password" />
          </div>
          <div v-if="alert">
            <h2 class="text-red-400 border-transparent">{{alert}}</h2>
          </div>
        </div>

        <div>
          <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-teal-500 hover:bg-teal-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-teal-500">
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
            </span>
            Register
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
    name:"register",
    data(){
        return{
            name: '',
            email: '',
            password: '',
            alert: '',
        }
    },
    methods: {
        async submit() {
            const response = await fetch('http://localhost:8000/api/register',{
                method:'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({
                    name: this.name,
                    email: this.email,
                    password: this.password
                })
            });
            const content = await response.json();
            console.log(content);
            this.alert = content.message;
            if (content.message){
              return ""
            }
            await this.$router.push('/login');
        }
    }
}
</script>

<style>

</style>