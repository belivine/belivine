<template>
  <router-view> </router-view>
</template>

<script>
export default {
  created:function(){
    this.$store.dispatch("user/setUser")
  }
}
</script>