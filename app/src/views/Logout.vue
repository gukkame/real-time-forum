<script setup>
import { getCookie, removeCookie } from "../../cookie";
var cookie = getCookie();
console.log("cookie!!! ", cookie);
removeCookie(cookie[1]);
location.assign("/");
</script>
<template>
  <body></body>
</template>
<script>
export default {};
</script>