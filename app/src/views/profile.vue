<script setup>
import { checkCookie } from "../../cookie";
let c = checkCookie();
</script>

<template>
  <div>
    <div class="profile" v-if="c">
      <h1>Profile page!</h1>

      <table>
        <tr>
          <th>Username</th>
          <th>{{ username }}</th>
        </tr>
        <tr>
          <th>Age</th>
          <th>{{ age }}</th>
        </tr>
        <tr>
          <th>Gender</th>
          <th>{{ gender }}</th>
        </tr>
        <tr>
          <th>First Name</th>
          <th>{{ firstname }}</th>
        </tr>
        <tr>
          <th>Last name</th>
          <th>{{ lastname }}</th>
        </tr>
        <tr>
          <th>E-mail</th>
          <th>{{ email }}</th>
        </tr>
      </table>
    </div>

    <div v-else>
      <h1>You have not logged in or your session has expired!</h1>
    </div>
  </div>
  <RouterView />
</template>

<script>
import axios from "axios";
import { getCookie } from "../../cookie";

export default {
  data() {
    return {
      info: null,
      username: this.username,
      age: this.age,
      gender: this.gender,
      firstname: this.firstname,
      lastname: this.lastname,
      email: this.email,
    };
  },
  mounted() {
    let cookie = getCookie();
    var data = {
      username: cookie[0],
      cookie: cookie[1],
    };
    axios({
      method: "POST",
      url: "http://localhost:8090/profile",
      data: data,
      headers: { "content-type": "text/plain" },
    }).then((res) => {
      this.info = res;
      this.username = res.data["username"];
      this.age = res.data["age"];
      this.age = getAge(this.age);
      this.gender = res.data["gender"];
      this.firstname = res.data["firstname"];
      this.lastname = res.data["lastname"];
      this.email = res.data["email"];
    });
  },
};

function getAge(dateString) {
  dateString = dateString.toString();
  var today = new Date();
  var birthDate = new Date(dateString);
  var age = today.getFullYear() - birthDate.getFullYear();
  var m = today.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}
</script>