<script setup>
import { checkCookie } from "../../cookie";
let c = checkCookie();
if (c == false) {
 
}
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
  <RouterView />
</template> 