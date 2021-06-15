<template>
  <div class=""></div>
</template>
<script>
import { createMemWords } from "@/server/account/account";
import { loginUser, layoutUser } from "@/server/user/login";
export default {
  methods: {
    registerUser() {
      console.log("=============");
      loginUser();
    },
    layout() {
      layoutUser();
    },
  },
};
</script>
