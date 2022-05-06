<script setup>
import { checkCookie } from "../../cookie";
let c = checkCookie();
</script>

<template>
  <body>
    <div class="center" v-if="c">
      <h1>You have already Loged in!</h1>
      <div class="signup_link" style="margin: 0px">
        Go back to <RouterLink to="/">Home page</RouterLink>
      </div>
    </div>

    <div class="center" v-else>
      <h1>Login</h1>
      <form>
        <div class="txt_field">
          <input name="username" v-model="username" type="text" required />
          <span></span>
          <label>Username</label>
        </div>
        <div class="txt_field">
          <input name="password" v-model="password" type="password" required />
          <span></span>
          <label>Password</label>
        </div>
      </form>
      &nbsp;
      <div class="submitbutton">
        <button class="submit-btn" variant="primary" v-on:click="postreq()">
          Log In
        </button>
      </div>
      &nbsp;

      <div>
        <p style="color: black">&nbsp; {{ msg }}</p>
      </div>
      <div class="signup_link">
        Not a member? <RouterLink to="/signup">SignUp</RouterLink>
      </div>

      <br />
    </div>
  </body>
</template>


<script>
import axios from "axios";
import { dateToString, checkCookie } from "../../cookie";

export default {
  data() {
    //data that will be displayed in page
    return {
      username: "",
      password: "",
      msg: "",
      date: dateToString(new Date()),
    };
  },

  methods: {
    postreq: function () {
      //data we are sending
      var data = {
        username: this.username,
        password: this.password,
      };

      axios({
        method: "POST",
        url: "http://localhost:8090/login",
        data: data,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
          this.username = result.data["username"]; // 1st add13 is name for data{} and 2nd what we recieve json from go
          this.msg = result.data["msg"];
          this.password = result.data["password"];
          this.check = result.data["check"];
          this.value = result.data["value"];
          this.name = result.data["name"];
          this.maxage = result.data["maxage"];

          if (this.check == true) {
            document.cookie =
              "session=" +
              this.value +
              ";" +
              "max-age=" +
              this.maxage +
              ";path=/";
            setTimeout(() => location.assign("/"), 1200);
          }

          // var cookie = getCookie();

          if (result.data.msg != "") {
            console.log("response", result.data.msg);
          }
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>