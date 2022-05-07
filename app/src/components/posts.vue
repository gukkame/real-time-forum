<script setup>
import { checkCookie } from "../../cookie";
let cookie = checkCookie();
</script>
<template>
  <body>
    <div class="main">
      <div class="title">
        <h2 class="t1">User</h2>
        <h2 class="t2">Post</h2>
        <h2 class="t3">Created</h2>
      </div>

      <div v-for="post in allposts" :key="post._id">
        <router-link
          :to="{
            name: 'post',
            params: { category: post.category, postid: post.id },
          }"
        >
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
        </router-link>
      </div>

      <div class="newPost" v-if="cookie">
        <RouterLink to="/home/newpost"><h2>Create new post</h2></RouterLink>
      </div>
      <div class="signup_link" v-else>
        <h4>
          Please <RouterLink to="/login">Log In </RouterLink>or
          <RouterLink to="/signup">Sign Up</RouterLink> to add posts!
        </h4>
      </div>
    </div>
  </body>
</template>
 <script>
import axios from "axios";
import { getCookie } from "../../cookie";
export default {
  name: "Posts",
  props: {
    pageCategory: String,
  },

  data() {
    return {
      id: this.id,
      title: this.title,
      content: this.content,
      created: this.created,
      username: this.username,
      category: this.category,
      pageCategory: this.pageCategory,
      result: [],
    };
  },

  computed: {
    allposts: function () {
      return this.result.filter((post) => post.category == this.pageCategory);
    },
  },
  mounted() {
    let cookie = getCookie();
    var data = {
      username: cookie[0],
      cookie: cookie[1],
    };
    axios({
      method: "POST",
      url: "http://localhost:8090/home/posts",
      data: data,
      headers: { "content-type": "text/plain" },
    }).then((res) => {
      this.id = res.data["id"];
      this.title = res.data["title"];
      this.content = res.data["content"];
      this.created = res.data["created"];
      this.category = res.data["category"];
      this.username = res.data["username"];
      this.result = res.data;
    });
  },
};
</script>