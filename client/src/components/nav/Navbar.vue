<template>
  <div>
    <!-- App.vue -->
    <v-app>
      <v-navigation-drawer
        app
        v-model="drawer"
        :mini-variant.sync="miniParent"
        mini-variant-width="56"
      >
        <div class="mx-auto leftBar">
          <v-navigation-drawer
            permanent
            width="100%"
            style="overflow: hidden !important"
          >
            <v-row class="fill-height" no-gutters>
              <v-navigation-drawer
                class="primary customBar"
                mini-variant
                mini-variant-width="56"
                permanent
              >
                <v-list-item class="px-2">
                  <v-list-item-avatar>
                    <v-img
                      src="https://randomuser.me/api/portraits/women/75.jpg"
                      @click="check()"
                    ></v-img>
                  </v-list-item-avatar>
                </v-list-item>

                <v-list dense nav>
                  <v-list-item-group v-model="selectedMenu" color="white">
                    <v-list-item v-for="item in menu" :key="item.title">
                      <v-list-item-action style="margin-left: 0px!important;">
                        <v-icon dense color="white">{{ item.icon }}</v-icon>
                      </v-list-item-action>
                    </v-list-item>
                  </v-list-item-group>
                </v-list>
                <template v-slot:append>
                  <div class="px-2 py-2">
                    <v-btn
                      color="white"
                      icon
                      medium
                    >
                      <v-icon>mdi-cog-outline</v-icon>
                    </v-btn>
                    </div>
                </template>
              </v-navigation-drawer>
              <v-list class="grow" v-if="!miniParent" key="list">
                <v-subheader> 
                  <v-row>
                    <v-col
                      md="9"
                    >
                      <p class="title-page">{{subMenuTitle}}</p> 
                    </v-col>
                    <v-col
                      md="3"
                    >
                    <v-btn
                      icon
                      color="secondary"
                      @click.stop="miniParent = !miniParent"
                    >
                      <v-icon>mdi-menu</v-icon>
                    </v-btn>
                    </v-col>
                  </v-row>
                </v-subheader>
                <!-- <v-divider></v-divider> -->
                <v-list shaped>
                <v-list-item-group
                  v-model="selectedSubMenu"
                  color="primary"
                >
                  <v-list-item
                    style="min-height: 41px;"
                    v-for="(item, i) in subMenu"
                    :key="i"
                  >
                    <v-list-item-icon style="margin: 9px 0px;min-width: 35px;">
                      <v-icon v-text="item.icon" dense></v-icon>
                    </v-list-item-icon>
                    <v-list-item-content style="padding: 0px;">
                      <v-list-item-title v-text="item.label"></v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                </v-list-item-group>
                </v-list>
              </v-list>
              <!-- </v-slide-y-transition> -->
            </v-row>
          </v-navigation-drawer>
        </div>
      </v-navigation-drawer>
      <v-app-bar
        absolute
        app
        color="white"
        elevate-on-scroll
        class="customAppBar"
        scroll-target="#scrolling-techniques-7"
      >
        <v-toolbar-title>{{$route.name}}</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>

        <v-btn icon>
          <v-icon>mdi-heart</v-icon>
        </v-btn>

        <v-btn icon>
          <v-icon>mdi-dots-vertical</v-icon>
        </v-btn>
      </v-app-bar>

      <!-- Sizes your content based upon application components -->
      <v-main>
         <v-sheet
          id="scrolling-techniques-7"
          class="overflow-y-auto"
          max-height="calc(100vh - 100px)"
        >
          <v-container style="height: 1500px;">
            <router-view></router-view>
          </v-container>
        </v-sheet>
      </v-main>
    </v-app>
  </div>
</template>
<script>
import menu from '../../json/menu.json';

export default {
  name:"Navbar",
  data() {
    return {
      selectedMenu: 0,
      selectedSubMenu: 0,
      drawer: true,
      mini: true,
      miniParent: false,
    };
  },
  computed: {
    menu: function(){
      return menu;
    },
    subMenu: function(){
      return menu[this.selectedMenu].children;
    },
    subMenuTitle: function(){
      return menu[this.selectedMenu].label;
    }
  },
  methods: {
    check: function(){
      console.log(this.selectedMenu);
    }
  }
};
</script>
<style lang="scss">
.leftBar {
  height: 100%;
  .title-page{
    padding: 0px;
    line-height: 36px;
    margin: 0px;
  }
}
html{
  overflow: hidden!important;
}
body{
  overflow: hidden!important;
}
.customAppBar{
  height: 57px!important;
}
.v-navigation-drawer__content {
  overflow-y: hidden !important;
}
.customBar {
  .v-navigation-drawer__border {
    background-color: transparent !important;
  }
  .v-list-item {
    min-height: 30px !important;
  }
}
.v-list-item__title {
  font-size: 0.9rem !important;
}
</style>
