<script setup>
import { checkCookie } from "../../cookie";
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
  <div class="main mainpage">
    <h1>Welcome to the home page!</h1>
    <h2>Happy to see you visiting our forum page! <br /></h2>
    <h3>
      <li>
        Choose categories, Log in or Sign Up, add posts, share opinion and help
        others!
      </li>
      <li>
        Hope you will learn more about tech and meet new friends, mentors or
        future employes! Enjoy! 
      </li>
    </h3>
    <br />
  </div>
  <RouterView />
</template>

