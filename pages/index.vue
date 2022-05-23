<template>
  <div class="mt-7 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-96 space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          {{ message }}
        </h2>
      </div>
    </div>
  </div>
</template>
<script>
import nuxtConfig from '../nuxt.config';

export default {
  data() {
    return {
      message: "",
    };
  },
  async mounted() {
    const response = await fetch("http://localhost:8000/api/user", {
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    const content = await response.json();
    console.log(content);
    if (content.name != null) {
      this.message = "Hello " + content.name+"\nYou are "+content.Role.name;
      this.$nuxt.$emit("auth", true);
    }
    // console.log(content.name);
    else {
      // console.log("ไม่ได้ login");
      // this.message = "You are not logged in.";
      this.$nuxt.$emit("auth", false);
      this.$router.push("/login");
    }
  },
};
</script>

<style>
</style>