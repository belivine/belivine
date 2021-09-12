<template>
  <v-app>
    <div
      class="d-flex align-sm-center mx-auto col-md-6"
      style="height: 100vh; flex: none"
    >
      <div
        class="
          col-md-12 col-sm-8 col-xm-12
          py-10
          mx-auto
          d-flex
          rounded-lg
          container-login
          flex-wrap
        "
      >
        <div class="col-md-6 col-xm-12 col-sm-12 px-md-10 px-sm-1">
          <v-img
            lazy-src="https://cdn.freelogovectors.net/wp-content/uploads/2020/11/cloudways-logo.png"
            width="130"
            src="https://cdn.freelogovectors.net/wp-content/uploads/2020/11/cloudways-logo.png"
          ></v-img>
          <div class="mb-4 mt-4">
            <div class="text-h6 font-weight-bold">Sign Up</div>
            <div class="text-caption">Make your work easy with us.</div>
          </div>
          <div class="rounded-xl btn-other text-center pa-2 mb-3">
            <img src="../../assets/google.png" alt="google-login" />
            Sign in with Google
          </div>

          <div class="d-flex align-center">
            <v-divider></v-divider>
            <div class="text-caption px-3">or Sign in with Email</div>
            <v-divider></v-divider>
          </div>
          <v-form
            ref="form"
            @submit.prevent="loginUser"
            v-model="valid"
            lazy-validation
            class="text-center mt-2"
          >
            <v-text-field
              v-model="form.username"
              label="Username"
              dense
              outline
              rounded
              :rules="[v => !!v || 'Username is required']"
              outlined
            ></v-text-field>

            <v-text-field
              :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :type="showPassword ? 'text' : 'password'"
              name="input-10-2"
              v-model="form.password"
              label="Password"
              :rules="[v => !!v || 'Password is required']"
              dense
              rounded
              class="input-group--focused"
              outlined
              @click:append="showPassword = !showPassword"
            ></v-text-field>

            <v-btn
              elevation="2"
              type="submit"
              class="primary"
              style="text-transform: none"
              :loading="loadingBtn"
              :disabled="loadingBtn"
              rounded
              >Login</v-btn
            >
          </v-form>
          <p class="caption red--text text-center mt-2">
            {{ message }}
          </p>
        </div>
        <v-img
          class="col-md-6 hidden-sm-and-down"
          :lazy-src="banner"
          max-height="100%"
          max-width="100%"
          :src="banner"
        ></v-img>
      </div>
    </div>
  </v-app>
</template>

<script>
import { mapActions } from "vuex";
export default {
  data: () => ({
    form: {
      username: "",
      password: "",
    },
    message: "",
    valid: true,
    showPassword: false,
    loadingBtn: false
  }),
  computed: {
    banner() {
      return require("@/assets/work.svg");
    },
  },
  methods: {
    ...mapActions({
        login : 'user/login'
    }),
    loginUser: function() {
      let validate = this.$refs.form.validate();
      if(validate){
        this.loadingBtn = true; 
        this.login(this.form).then((res) => {
          this.loadingBtn = false; 
          this.message = res;
          this.$router.push('/time/tracking');
        });
      }
    },
  },
};
</script>

<style lang="scss">
.container-login {
  border: 1px solid #d2d2d2;
}

.btn-other {
  border: 1px solid #d2d2d2;
  cursor: pointer;
  img {
    width: 23px;
    height: 100%;
    vertical-align: middle;
  }
}

@media (max-width: 600px) {
  .container-login {
    border: none;
  }
}
.v-application--is-ltr .v-text-field .v-label {
  font-size: 14px;
}
.v-input input {
  font-size: 14px;
}
</style>
