<script setup>
import { checkCookie } from "../../cookie";
let cookie = checkCookie();
//chat is available only when user is loged in
</script>

<template>
  <body>
    <form :action="sendMessage" @click.prevent="onSubmit">
      <input v-model="message" type="text" />
      <input type="submit" value="Send" @click="sendMessage" />
    </form>
    <p>Two way data binding is fun!</p>
    <p>
      {{ message }}
    </p>

    <div v-for="user in alluse0rs" :key="user._id">
      {{ user.Username }}
    </div>

    <div v-if="showMsg">
      <h3>Message in a WebSocket</h3>
      <p>
        {{ rcvMessage }}
      </p>

      <button @click="showMsg = !showMsg">Dismiss</button>
    </div>
  </body>
</template>

<script>
import axios from "axios";
import { getCookie } from "../../cookie";
let cookie = getCookie();

if (navigator.onLine) {
  console.log("online");
} else {
  console.log("offline");
}
export default {
  name: "App",
  data() {
    return {
      message: "",
      socket: null,
      rcvMessage: "",
      showMsg: false,
      allusers: this.allusers,
    };
  },

  mounted() {
    // window.setInterval(() => {
    //   this.getNotifications();
    // }, 5000),
    this.socket = new WebSocket("ws://localhost:8090/socket", [], {
      headers: {
        Cookie: cookie,
      },
    });
    this.socket.onopen = () => {
      console.log("Websocket connected.");
      this.connectedStatus = "Connected";
    };
    this.socket.onmessage = (msg) => {
      // this.acceptMsg(msg);
      let incomingdata = JSON.parse(msg.data.toString());
      // this.jsonMessage = JSON.parse(msg);
      console.log(incomingdata);
      this.allusers = incomingdata.allUsers;
      // console.log("users ", incomingdata.allUsers[0].Username);
      // console.log("users ", this.allusers[1]);
      console.log("content ", incomingdata.content);
      // this.incomingdata = incomingdata.data;
      // this.acceptMsg(msg);
    };
  },

  methods: {
    getNotifications() {
      let cookie = getCookie();
      if (cookie[1] != "") {
        let status = {
          status: navigator.onLine,
          sessionID: cookie[1],
        };
        console.log(status);
        this.socket.send(JSON.stringify(status));
      }
    },

    sendMessage() {
      let cookie = getCookie();
      console.log(cookie);
      let msg = {
        content: this.message,
        sessionID: cookie[1],
        status: navigator.onLine,
      };
      this.socket.send(JSON.stringify(msg));
    },
    acceptMsg(msg) {
      // this.rcvMessage = msg.data;
      // this.showMsg = true;
    },
  },
};
</script>




<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>