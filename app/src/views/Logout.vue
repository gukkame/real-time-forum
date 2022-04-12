<script setup>
import { getCookie, removeCookie } from "../../cookie";
// var data = {
//   username: this.username,
//   password: this.password,
// };
// axios({
//   method: "POST",
//   url: "http://localhost:8090/logout",
//   data: data,
//   headers: { "content-type": "text/plain" },
// });

var cookie = getCookie();
removeCookie(cookie[0]);
location.assign("/");
</script>
