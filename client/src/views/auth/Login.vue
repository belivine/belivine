<template>
  <v-app>
    <div class="d-flex align-sm-center mx-auto col-md-6" style="height: 100vh; flex: none;">
      <div class="col-md-12 col-sm-8 col-xm-12 py-10 mx-auto d-flex rounded-lg container-login flex-wrap">
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
              <img src="../../assets/google.png" alt="google-login">
              Sign in with Google
          </div>
          
          <div class="d-flex align-center">
              <v-divider></v-divider>
              <div class="text-caption px-3">or Sign in with Email</div>
              <v-divider></v-divider>

          </div>
          <v-form
            ref="form"
            @submit.prevent="submitData"
            v-model="valid"
            lazy-validation
            class="text-center mt-2"
          >
            <v-text-field
              v-model="form.username"
              label="Username"
              dense
              outline
              required
              outlined
            ></v-text-field>

            <v-text-field
              v-model="form.email"
              :rules="emailRules"
              label="E-mail"
              required
              outlined
              dense
            ></v-text-field>

            <v-text-field
              :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[rules.required, rules.min]"
              :type="showPassword ? 'text' : 'password'"
              name="input-10-2"
              v-model="form.password"
              label="Password"
              dense
              hint="At least 8 characters"
              class="input-group--focused mt-2"
              outlined
              @click:append="showPassword = !showPassword"
            ></v-text-field>

            <v-btn elevation="2" type="submit" class="primary" style="text-transform: none;" v-on:click="login">Sign Up</v-btn>
          </v-form>
        </div>
        <v-img class="col-md-6 hidden-sm-and-down"
            lazy-src="../../assets/work.svg"
            max-height="100%"
            max-width="100%"
            src="../../assets/work.svg"
        ></v-img>
      </div>
    </div>
  </v-app>
</template>

<script>

import gql from 'graphql-tag'

export default {
  data: () => ({
    form: {
        username:'',
        email:'',
        password: '',
    },
    

    valid: true,
    username: "",

    nameRules: [
      (v) => !!v || "Name is required",
      (v) => (v && v.length <= 10) || "Name must be less than 10 characters",
    ],
    email: "",
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
    ],
    showPassword: false,
     rules: {
          required: value => !!value || 'Required.',
          min: v => v && v.length >= 8 || 'Min 8 characters',
          emailMatch: () => (`The email and password you entered don't match`),
        },
  }),

  methods: {
    validate() {
      this.$refs.form.validate();
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
    submitData(){
        console.log(this.form);
    },
    login(){
      this.$apollo.mutate({
         mutation: gql`mutation {
              login(input : {username: "user1", password: "123"})
          }`,
      }).then((data) => {
        this.$cookies.set('belivine',data.data.login)
      }).catch((error) => {
        // Error
        console.error(error)
        console.log("ok");
      })
    }
  },
};
</script>

<style lang="scss">
.container-login{
    border: 1px solid #d2d2d2;
}

.btn-other{
    border: 1px solid #d2d2d2;
    cursor: pointer;
    img{
        width: 23px;
        height: 100%;
        vertical-align: middle;
    }
}

@media (max-width: 600px){
 .container-login{
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
