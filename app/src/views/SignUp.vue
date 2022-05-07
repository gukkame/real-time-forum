<template>
  <body>
    <div class="center">
      <h1>Sign up</h1>
      <form>
        <div class="txt_field">
          <input
            name="username"
            v-model="username"
            type="text"
            placeholder="User Name"
            required
          />
          <span></span>
        </div>
        <div class="txt_field">
          <label>Age</label>
          <br />
        </div>
        <div class="formdate">
          <input type="date" v-model="age" required />
          <br />
        </div>
        <div class="txt_field">
          <br />
          <label>Gender</label>
        </div>
        <div class="radiobtn">
          <input type="radio" id="female" value="female" v-model="gender" />
          <label for="female">Female</label> &nbsp; &nbsp;
          <input type="radio" id="male" value="male" v-model="gender" />
          <label for="male">Male</label>&nbsp; &nbsp;
          <input
            type="radio"
            id="no_gender"
            value="no_gender"
            v-model="gender"
          />
          <label for="no_gender">No gender</label>&nbsp; &nbsp;
        </div>
        <div class="txt_field">
          <input
            name="firstname"
            v-model="firstname"
            type="text"
            placeholder="First Name"
            required
          />
          <span></span>
        </div>
        <div class="txt_field">
          <input
            name="lastname"
            v-model="lastname"
            type="text"
            placeholder="Lastname"
            required
          />
          <span></span>
        </div>
        <div :class="['input-group', isEmailValid()]" class="txt_field">
          <!-- <input name="email" v-model="email" type="text" required /> -->
          <input type="email" placeholder="Email Address" v-model="email" />
          <span></span>
        </div>
        <div class="txt_field">
          <input
            name="password"
            v-model="password"
            type="password"
            placeholder="Password"
            required
          />
          <span></span>
        </div>
      </form>
      &nbsp;
      <div class="submitbutton">
        <button class="submit-btn" variant="primary" v-on:click="postreq()">
          Sign up
        </button>
      </div>
      &nbsp;
      <div>
        <p style="color: black">&nbsp;{{ msg }}</p>
      </div>
      <div class="signup_link">
        Do you have account already? Then
        <RouterLink to="/login">LogIn</RouterLink>
      </div>
      <br />
    </div>
  </body>
</template>

<script>
import axios from "axios";
import { dateToString } from "../../cookie";

export default {
  name: "SignUp",

  data: function () {
    //data that will be displayed in page
    return {
      username: "",
      gender: "",
      firstname: "",
      lastname: "",
      age: "",
      email: "",
      isValid: "",
      reg: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,24}))$/,
      password: "",
      msg: "",
    };
  },

  methods: {
    isEmailValid() {
      return this.email == ""
        ? ""
        : this.reg.test(this.email)
        ? ((this.isValid = true), (this.msg = ""))
        : (this.msg = "email is incorect");
    },

    postreq: function () {
      //data we are sending
      var data = {
        username: this.username,
        age: this.age,
        gender: this.gender,
        firstname: this.firstname,
        lastname: this.lastname,
        email: this.email,
        password: this.password,
      };
      console.log("all data from form", data);
      if (
        this.username != "" &&
        this.firstname != "" &&
        this.lastname != "" &&
        this.email != "" &&
        this.age != "" &&
        this.gender != "" &&
        this.password != "" &&
        this.isValid == true
      ) {
        axios({
          method: "POST",
          url: "http://localhost:8090/signup",
          data: data,
          headers: { "content-type": "text/plain" },
        })
          .then((result) => {
            //data we receive
            this.msg = result.data["msg"];
            this.flag = result.data["check"];
            if (this.flag == true) {
              console.log("Data Sended ", result.data);
              setTimeout(() => this.$router.push({ path: "/" }), 2000);
            }
          })
          .catch((error) => {
            console.error(error);
          });
      } else {
        this.msg = "Please fill all the form";
      }
    },
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Noto+Sans:wght@700&family=Poppins:wght@400;500;600&display=swap");
</style>
