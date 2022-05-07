<script setup>
import { getCookie } from "../../cookie";
let cookie = getCookie();
</script>
<template>
  <body class="flex-container">
    <!-- Chat -->
    <div v-for="(user, userid) in privatechats" :key="user._id">
      <section class="msger" v-bind:id="`mySidepanel` + userid">
        <button v-on:click="chat(userid, user)">
          <div class="msger-header">
            &#9776; &nbsp;
            <div class="msger-header-title">&nbsp;{{ user }}</div>
            <div class="msger-header-options"></div>
          </div>
        </button>

        <div
          v-on:scroll="handleScroll(user)"
          v-bind:id="`chat` + userid"
          class="msger-chat"
        >
          <span v-for="msg in orderedMsg()" :key="msg._id">
            <span
              v-if="
                (msg.sender == user && msg.receiver == cookie[0]) ||
                (msg.sender == cookie[0] && msg.receiver == user)
              "
            >
              {{ msg.date }}
            </span>

            <div class="msg left-msg">
              <div
                class="msg-img"
                v-if="msg.sender == user && msg.receiver == cookie[0]"
              ></div>
              <div
                class="msg-bubble"
                v-if="msg.sender == user && msg.receiver == cookie[0]"
              >
                <div class="msg-text">{{ msg.content }}</div>
                <div class="msg-info-time">{{ msg.time }}</div>
              </div>
            </div>

            <div class="msg right-msg">
              <div
                class="msg-img"
                v-if="msg.sender == cookie[0] && msg.receiver == user"
              ></div>
              <div
                class="msg-bubble"
                v-if="msg.sender == cookie[0] && msg.receiver == user"
              >
                <div class="msg-text">{{ msg.content }}</div>
                <div class="msg-info-time">{{ msg.time }}</div>
              </div>
            </div>
          </span>
        </div>

        <form
          :action="sendMessage"
          @click.prevent="onSubmit"
          v-bind:id="`chatform` + userid"
          class="msger-inputarea"
        >
          <input
            type="text"
            class="msger-input"
            placeholder="Enter your message..."
            v-bind:id="`input` + userid"
            v-model="message"
          />

          <button
            type="submit"
            value="Send"
            class="msger-send-btn"
            @click="sendMessage(user, userid)"
            v-bind:id="`input` + userid"
          >
            Send
          </button>
        </form>
      </section>
    </div>
    <!-- All Users panel-->
    <section id="mySidepanelUsers" class="allUsers">
      <button v-on:click="allUsers(), reqAllUsers()">
        <header class="msger-header">
          &#9776;
          <div class="msger-header-title">All Users</div>
          <div class="msger-header-options"></div>
        </header>
      </button>

      <div id="chat1" class="msger-chat">
        <button
          class="users selectUserbtn"
          v-for="user in listOnlineUsers"
          :key="user._id"
          v-on:click="privateChat(user.Username)"
        >
          {{ user.Username }}
        </button>
      </div>
    </section>
  </body>
</template>

<style>
@import "../assets/chat.css";
</style>


<script>
import { getCookie, removeCookie } from "../../cookie";
import _ from "lodash";
import "vue-toast-notification/dist/theme-sugar.css";
let cookie = getCookie();

export default {
  name: "App",
  data() {
    return {
      // message: "",
      socket: null,
      rcvMessage: "",
      showMsg: false,
      allusers: [],
      allMsg: [],
      user1: [],
      user2: [],
      privatechats: "",
      onlineUserList: [],
      message: "",
      lastMsg: "",
      recNotSender: "",
      recNotReciever: "",
    };
  },
  computed: {
    listOnlineUsers() {
      if (this.allusers != null) {
        let filtered = this.allusers.filter(
          (user) => user.Username != cookie[0]
        );
        let allUsersO = _.orderBy(
          filtered,
          ["LastMsg", "Username"],
          ["asc", "asc"]
        );
        return allUsersO;
      }
    },
  },
  created() {
    window.addEventListener("scroll", this.handleScroll);
  },
  destroyed() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  mounted() {
    // window.setInterval(() => {
    // this.checkSession();
    // this.reqAllUsers();
    // }, 10000),
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
      let incomingData = JSON.parse(msg.data.toString());

      if (incomingData == undefined || incomingData == null) {
        console.log("null!!!!");
      } else {
        if (Object.keys(incomingData)[0] == "sender") {
          this.recNotification(incomingData);
        }
        if (Object.keys(incomingData)[0] == "content") {
          this.displayIncomingMsg(incomingData);
        }
        if (Object.keys(incomingData)[0] == "allUsers") {
          this.receivedAllUsers(incomingData);
        }
        if (Object.keys(incomingData)[0] == "msg") {
          this.recCheckSession(incomingData);
        }
        if (Object.keys(incomingData)[0] == "0") {
          this.receivedAllMsg(incomingData);
        }
      }
    };
  },

  methods: {
    displayIncomingMsg(incomingData) {
      if (
        document.getElementById("mySidepanel0") == null ||
        document.getElementById("mySidepanel0").style.height == "0px"
      ) {
      } else {
        this.allMsg.push(incomingData);
      }
    },
    //Orders messages and remove date
    orderedMsg() {
      let lastdate = "";
      let user = [];
      let allMessages = _.orderBy(this.allMsg, "created");
      allMessages.forEach((msg, i) => {
        if (
          msg.date == lastdate &&
          i > 0 &&
          ((msg.receiver == user[0] && msg.sender == user[1]) ||
            (msg.receiver == user[1] && msg.sender == user[0]))
        ) {
          return (msg.date = "");
        }
        user[0] = msg.receiver;
        user[1] = msg.sender;
        lastdate = msg.date;
      });

      let oMsg = _.orderBy(allMessages, "created");
      if (document.getElementById("mySidepanel0")) {
        var c = document.getElementById("chat0");
        c.scrollTo(0, c.scrollHeight + 10);
      }

      return oMsg;
    },
    handleScroll(user) {
      var c = document.getElementById("chat0");
      if (c.scrollTop == 0) {
        let msg = {
          user2: user,
          content: "",
          number: this.allMsg.length,
        };
        this.socket.send(JSON.stringify(msg));
        c.scrollTo(0, c.scrollHeight + 10);
      }
    },
    receivedAllUsers(incomingData) {
      this.allusers = incomingData.allUsers;
      // console.log("AllUsers ", incomingData.allUsers);
    },
    receivedAllMsg(incomingData) {
      this.allMsg = incomingData;
    },
    recNotification(incomingData) {
      this.recNotSender = incomingData.sender;
      this.recNotReciever = incomingData.receiver;
      if (
        document.getElementById("mySidepanel0") == null ||
        document.getElementById("mySidepanel0").style.height == "0px"
      ) {
        alert("You have new message");
      }
      if (document.getElementById("chat0") != null) {
        var c = document.getElementById("chat0");
        c.scrollTo(0, c.scrollHeight);
      }
    },

    checkSession() {
      let cookie = getCookie();
      if (cookie[1] != "") {
        let status = {
          content: "check",
          sessionID: cookie[1],
          user1: cookie[0],
        };
        this.socket.send(JSON.stringify(status));
      }
    },
    //Check session if user is still active
    recCheckSession(incomingData) {
      let cookie = getCookie();
      this.message = incomingData;
      if (this.message.check == false) {
        if (getCookie != []) {
          removeCookie(cookie[1]);
          // location.assign("/logout");
          // console.log("Loging you out");
        }
      }
    },

    //Opens private chat window
    privateChat(username) {
      if (username != cookie[0]) {
        this.privatechats = username;
      }

      //For Multiple chat logs opened

      // let i = 0;
      // this.privatechats.forEach((element) => {
      //   if (username == element) {
      //     i++;
      //   }
      // });
      // if (i == 0 && username != cookie[0]) {
      //   this.privatechats.push(username);
      // }
    },

    sendMessage(user) {
      let cookie = getCookie();
      let msg = {
        user1: cookie[0],
        user2: user,
        content: this.message,
        sessionID: cookie[1],
      };
      this.socket.send(JSON.stringify(msg));
      this.message = "";
      var c = document.getElementById("chat0");
      c.scrollTo(0, c.scrollHeight);
      console.log("Message sended!");
    },
    reqAllUsers() {
      console.log("Req users! ");
      let userlist = {
        content: "getUserList",
      };
      console.log("User List Online ", userlist);
      this.socket.send(JSON.stringify(userlist));
    },

    allUsers() {
      if (
        document.getElementById("mySidepanelUsers").style.height == "40px" ||
        document.getElementById("mySidepanelUsers").style.height == ""
      ) {
        document.getElementById("mySidepanelUsers").style.height = "540px";
        document.getElementById("mySidepanelUsers").style.width = "160px";
        document.getElementById("chat1").style.display = "block";
      } else {
        document.getElementById("mySidepanelUsers").style.height = "40px";
        document.getElementById("mySidepanelUsers").style.width = "100px";
        document.getElementById("chat1").style.display = "none";
      }
    },
    chat(userid, user2) {
      //When pressed Chat window opens
      if (
        document.getElementById("mySidepanel" + userid).style.height ==
          "40px" ||
        document.getElementById("mySidepanel" + userid).style.height == ""
      ) {
        let msg = {
          user2: user2,
          content: "",
        };
        this.socket.send(JSON.stringify(msg));
        document.getElementById("chat" + userid).style.display = "block";
        document.getElementById("chatform" + userid).style.display = "flex";
        document.getElementById("mySidepanel" + userid).style.height = "540px";
      } else {
        //closes
        let msg = {
          user2: user2,
          content: "chatCloses",
        };
        this.socket.send(JSON.stringify(msg));
        document.getElementById("mySidepanel" + userid).style.height = "40px";
        document.getElementById("chat" + userid).style.display = "none";
        document.getElementById("chatform" + userid).style.display = "none";
      }
    },
  },
};
</script>