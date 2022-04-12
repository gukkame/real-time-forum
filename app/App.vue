<script setup>
// import { getCookie } from "./cookie";
import { checkCookie } from "./cookie";
let c = checkCookie();
</script>

<template>
  <nav class="topnav">
    <div style="float: left">
      <RouterLink to="/home">Home</RouterLink>
      <h2>Forum &nbsp;</h2>
    </div>
    <div class="button1" v-if="c">
      <RouterLink to="/logout">&nbsp;Log out</RouterLink>
      <RouterLink to="/profile">&nbsp; Profile &nbsp;</RouterLink>
    </div>
    <div v-else>
      <RouterLink to="/signup">Sign up</RouterLink>
      <RouterLink to="/login">Log in</RouterLink>
    </div>
  </nav>
  <RouterView />
</template>

<script>
</script>

<style>
@import "@/assets/base.css";
@import "@/assets/app.css";
@import "@/assets/forms.css";
@import "@/assets/profile.css";
</style>
