<script setup>
import { getCookie, removeCookie } from "../../cookie";
var cookie = getCookie();
removeCookie(cookie[0]);
location.assign("/");
</script>
