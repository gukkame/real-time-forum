<script setup>
import { checkCookie } from "../../cookie";
let c = checkCookie();
</script>
<template>
  <body>
    <div class="main" v-if="c">
      <h1>New post!</h1>
      <form class="newpostform">
          
        <h2 style="color: red;   text-align: center;">{{ msg }}</h2>

        <div class="txt_field">
          <input
            placeholder="Title"
            name="title"
            v-model="title"
            type="text"
            required
          />
          <span></span>
          <!-- <label>Title</label> -->
        </div>
    
        <div class="txt_field">
          <textarea
            placeholder="Content"
            v-model="content"
            type="text"
            required
          ></textarea>
        </div>
        <br />
        <div class="txt_field1">
          <br />
          <label>Category</label>
        </div>
        <div class="radiobtn">
          <input type="radio" id="nft" value="nft" v-model="category" />
          <label for="nft">NFT</label> &nbsp; &nbsp;
          <input
            type="radio"
            id="nftgames"
            value="nftgames"
            v-model="category"
          />
          <label for="nftgames">Nftgames</label>&nbsp; &nbsp;
          <input type="radio" id="crypto" value="crypto" v-model="category" />
          <label for="crypto">Crypto</label>&nbsp; &nbsp;
          <input
            type="radio"
            id="blockchain"
            value="blockchain"
            v-model="category"
          />
          <label for="blockchain">Blockchain</label>&nbsp; &nbsp;
          <input type="radio" id="web3" value="web3" v-model="category" />
          <label for="web3">Crypto</label>
        </div>
      </form>
      <div class="newPost">
        <button class="submit-btn" variant="primary" v-on:click="postreq()">
          Add post
        </button>
      </div>
    </div>
    <div class="main" v-else>
      <h1>You have not logged in or your session has expired!</h1>
      <h1 class="signup_link">
        Please <RouterLink to="/login">Log In </RouterLink>or
        <RouterLink to="/signup">Sign Up</RouterLink> to add posts!
      </h1>
    </div>
  </body>
</template>
<script>
import axios from "axios";
import { getCookie } from "../../cookie";

export default {
  name: "NewPost",

  data() {
    //data that will be displayed in page
    return {
      title: "",
      content: "",
      category: "",
      msg: "",
    };
  },
  methods: {
    postreq: function () {
      //data we are sending
      let cookie = getCookie();
      var data = {
        title: this.title,
        content: this.content,
        category: this.category,
        username: cookie[0],
        cookie: cookie[1],
      };
      console.log("all data from form", data);
      if (this.title != "" && this.content != "" && this.category != "") {
        axios({
          method: "POST",
          url: "http://localhost:8090/newpost",
          data: data,
          headers: { "content-type": "text/plain" },
        })
          .then((result) => {
            //data we receive
            this.msg = result.data["msg"];
            this.flag = result.data["check"];
            console.log("result4 ", result.data);
            if (this.flag == true) {
              setTimeout(() => this.$router.push({ path: "/" }), 1000);
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