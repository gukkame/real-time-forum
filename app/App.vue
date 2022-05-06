<script setup>
import { checkCookie } from "./cookie";
import Chat from "./src/components/chat.vue";
let cookie = checkCookie();
</script>

<template>
  <nav class="topnav">
    <div style="float: left">
      <RouterLink to="/home">Home</RouterLink>
      <h2>Forum &nbsp;</h2>
    </div>
    <div v-if="cookie">
      <RouterLink to="/logout">&nbsp;Log out</RouterLink>
      <RouterLink to="/profile">&nbsp; Profile &nbsp;</RouterLink>
    </div>
    <div v-else>
      <RouterLink to="/signup">Sign up</RouterLink>
      <RouterLink to="/login">Log in</RouterLink>
    </div>
  </nav>
  <RouterView />
  <Chat v-if="cookie" />
</template>

<script>
export default {
  components: {
    Chat,
  },
};
</script>

<style>
@import "@/assets/base.css";
@import "@/assets/app.css";
@import "@/assets/forms.css";
@import "@/assets/profile.css";
@import "@/assets/posts.css";
</style>
