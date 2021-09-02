<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
      {{$store.state.counter}}
    <button @click="$store.dispatch('increase')">+</button>
    <button @click="$store.dispatch('decrease')">-</button>
    </div>
</template>

<script>
// @ is an alias to /src
import store from '../store'
export default {
    store,
    name: 'Home',
    created: () => {
        console.log(store.state.counter);
    }
}
</script>
