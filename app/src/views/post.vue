<script setup>
import { checkCookie } from "../../cookie";
let cookie = checkCookie();
</script>
<template>
  <body>
    <h1>Single Post page</h1>
    <div>This is Bar {{ $route.params.postid }}</div>

    <div v-for="post in allposts" :key="post._id">
      <div class="post">
        <div class="user">
          <img src="../assets/original.jpg" alt="profile_img" />
          <h3>{{ post.username }}</h3>
        </div>
        <div class="content">
          <h2>{{ post.title }}</h2>
          <p>{{ post.content }}</p>
        </div>
        <div class="created">{{ post.postcreated }}</div>
      </div>
    </div>
    <!-- <Comments /> -->
    <div class="main" style="margin-top: 0px">
      <div class="post">
        <h2>Comments</h2>
      </div>

      <div v-for="post in com_result" :key="post._id">
        <div class="post">
          <div class="user">
            <img src="../assets/original.jpg" alt="profile_img" />
            <h3>{{ post.username }}</h3>
          </div>
          <div class="content">
            <h2>{{ post.title }}</h2>
            <p>{{ post.content }}</p>
          </div>
          <div class="created">{{ post.created }}</div>
        </div>
      </div>

      <div v-if="cookie">
        <form class="newpostform">
          <div class="txt_field">
            <input
              name="comContent"
              v-model="comContent"
              type="text"
              required
            />
            <span></span>
            <label>Write your comment here... </label>
            
          </div>
        </form>

        <div class="newPost">
          <button class="submit-btn" variant="primary" v-on:click="postreq()">
            Add comment
          </button>
        </div>
        <h3 style="color: red;   text-align: center;">{{ msg }}</h3>
      </div>
    </div>
  </body>
</template>
<script>
import axios from "axios";
import { getCookie } from "../../cookie";
let cookie = getCookie();
export default {
  name: "Posts",
  data() {
    return {
      id: this.id,
      title: this.title,
      content: this.content,
      postcreated: this.postcreated,
      username: this.username,
      category: this.category,
      postid: this.postid,
      result: [],
      com_id: this.com_id,
      com_content: this.com_content,
      com_username: this.com_username,
      com_result: [],
      comContent: "",
      msg: "",
    };
  },
  created() {
    let postid = window.location.pathname.split("/");
    this.postid = postid[3];
  },
  computed: {
    allposts() {
      console.log("postid", this.postid);
      return this.result.filter((post) => post.id == this.postid);
    },
  },
  async mounted() {
    var data = {
      username: cookie[0],
      cookie: cookie[1],
    };
    await axios({
      method: "POST",
      url: "http://localhost:8090/home/posts",
      data: data,
      headers: { "content-type": "text/plain" },
    }).then((res) => {
      this.id = res.data["id"];
      this.title = res.data["title"];
      this.content = res.data["content"];
      this.postcreated = res.data["postcreated"];
      this.category = res.data["category"];
      this.username = res.data["username"];
      this.result = res.data;
      console.log(this.result);
    });
    await axios({
      method: "POST",
      url: "http://localhost:8090/home/singlepost",
      data: this.postid,
      headers: { "content-type": "text/plain" },
    }).then((response) => {
      this.com_id = response.data["id"];
      this.com_content = response.data["content"];
      //   this.title = res.data["postid"];
      this.com_username = response.data["username"];
      this.com_result = response.data;
      console.log("ssxxxs ", this.com_result);
    });
  },
  methods: {
    postreq: function () {
      //data we are sending
      var data = {
        comContent: this.comContent,
        postid: parseInt(this.postid),
        username: cookie[0],
        cookie: cookie[1],
      };
      console.log("all data from form", data);
      console.log("all data from form", typeof this.postid);
      console.log("all data from form", cookie[0]);
      if (this.comContent != "" && cookie[0] != "") {
        axios({
          method: "POST",
          url: "http://localhost:8090/home/newcomment",
          data: data,
          headers: { "content-type": "text/plain" },
        })
          .then((result) => {
            //data we receive
            this.msg = result.data["msg"];
            this.flag = result.data["check"];
            console.log("result4 ", result.data);

            if (this.flag == true) {
              console.log("restart", window.location.pathname);
              window.location.reload();
            }
          })
          .catch((error) => {
            console.error(error);
          });
      } else {
        this.msg = "Please write comment to add it!";
      }
    },
  },
};
</script>