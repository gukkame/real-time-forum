<template>
  <body>
    <div class="center">
      <h1>Sign up</h1>
      <form>
        <div class="txt_field">
          <input name="username" v-model="username" type="text" required />
          <span></span>
          <label>Username</label>
        </div>
        <div class="txt_field1">
          <br />
          <label>Age</label>
        </div>
        <div>
          <input
            style="font-size: 20px; margin-bottom: 20px"
            type="date"
            v-model="age"
            required
          />
          <br />
        </div>
        <div class="txt_field1">
          <br />
          <label>Gender</label>
        </div>
        <div class="radiobtn" style="margin: 0 0">
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
          <label for="no_gender">No gender</label>
        </div>
        <div class="txt_field">
          <input name="firstname" v-model="firstname" type="text" required />
          <span></span>
          <label>Firstname</label>
        </div>
        <div class="txt_field">
          <input name="lastname" v-model="lastname" type="text" required />
          <span></span>
          <label>Lastname</label>
        </div>
        <div class="txt_field">
          <input name="email" v-model="email" type="text" required />
          <span></span>
          <label>Email</label>
        </div>
        <div class="txt_field">
          <input name="password" v-model="password" type="password" required />
          <span></span>
          <label>Password</label>
        </div>
      </form>

      <!-- <div>
        <p style="color: black">Add: {{ add13 }}</p>
      </div>
      <div>
        <p style="color: black">Multiplication: {{ multi }}</p>
      </div> -->

      &nbsp;
      <div class="submitbutton">
        <button class="submit-btn" variant="primary" v-on:click="postreq()">
          Sign up
        </button>
      </div>
      &nbsp;
      <div>
        <p style="color: black">{{ msg }}</p>
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
      gender: "female",
      firstname: "",
      lastname: "",
      age: dateToString(new Date()),
      email: "",
      password: "",
      msg: "",
    };
  },

  methods: {
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
        this.password != ""
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
            console.log("result4 ", result.data);
            if (this.flag == true) {
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
