<template>
  <div class="home">
    <h1 class="title is-1">Home</h1>
    <LoginWithSpotify></LoginWithSpotify>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import LoginWithSpotify from "@/components/LoginWithSpotify.vue"; // @ is an alias to /src

@Component({
  components: {
    LoginWithSpotify,
  },
})
export default class Home extends Vue {}
</script>
