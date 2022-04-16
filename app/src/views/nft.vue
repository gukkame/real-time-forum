<script setup>
import { checkCookie } from "../../cookie";
import Posts from "../components/posts.vue";
checkCookie();
</script>
<template>
  <div class="wrapper">
    <nav>
      Categories: &nbsp;
      <RouterLink to="/home/nft">NFT</RouterLink>
      <RouterLink to="/home/nftgames">NFT Games</RouterLink>
      <RouterLink to="/home/crypto">Crypto</RouterLink>
      <RouterLink to="/home/blockchain">Blockchain</RouterLink>
      <RouterLink to="/home/web3">Web 3.0</RouterLink>
    </nav>
  </div>
  <Posts :pageCategory="pageCategory" />
  <RouterView />
</template> 
<script>
export default {
  components: {
    Posts,
  },
  data() {
    return {
      pageCategory: this.pageCategory,
    };
  },
  created() {
    let pageCategory = window.location.pathname.split("/");
    this.pageCategory = pageCategory[2];
  },
};
</script>