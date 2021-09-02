<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="12" md="8" sm="12">
          <v-row justify="space-between">
            <v-col cols="6">
              <h3>Timesheet</h3>
            </v-col>
            <v-col cols="3" sm="6" md="3">
              <v-btn :elevation="0" rounded small class="text-capitalize"
                >All Task</v-btn
              >
              <v-btn rounded small class="ml-3 text-capitalize" text
                >My Task</v-btn
              >
            </v-col>
          </v-row>
          <v-list subheader two-line>
            <v-subheader>Fri, 12 Aug</v-subheader>
            <v-divider></v-divider>
            <v-list-item v-for="folder in folders" :key="folder.title">
              <v-list-item-avatar>
                    <v-btn icon class="mb-1 flex-grow-1">
                        <v-icon color="blue lighten-1">mdi-checkbox-multiple-marked-circle-outline</v-icon>
                    </v-btn>
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title v-text="folder.title"></v-list-item-title>

                <v-list-item-subtitle
                  v-text="folder.subtitle"
                ></v-list-item-subtitle>
              </v-list-item-content>

              <v-list-item-action>
                <div class="d-flex">
                    <p class="mb-1 flex-grow-1 time mr-2">{{project.time}}</p>
                    <v-btn icon class="mb-1 flex-grow-1">
                        <v-icon color="blue lighten-1">mdi-play</v-icon>
                    </v-btn>
                </div>
              </v-list-item-action>
            </v-list-item>

            <!-- <v-divider inset></v-divider> -->
          </v-list>
        </v-col>
        <v-col cols="12" md="4" sm="12" class="bd-left">
          <h3>Time Tracking</h3>
          <div class="mt-8 text-center">
            <p class="gray--text text-center pd-0 mt-0">{{ project.name }}</p>
            <h2 class="text-center">{{ project.time }}</h2>
            <v-btn class="mx-2" fab dark small color="primary">
              <v-icon dark> mdi-play </v-icon>
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import gql from 'graphql-tag'

export default {
  data: () => {
    return {
      project: {
        name: "Design System Model CAP 1",
        time: "00:08:19",
      },
      files: [
        {
          color: "blue",
          icon: "mdi-clipboard-text",
          subtitle: "Jan 20, 2014",
          title: "Design Task On Model",
        },
        {
          color: "amber",
          icon: "mdi-gesture-tap-button",
          subtitle: "Jan 10, 2014",
          title: "Kitchen remodel",
        },
      ],
      folders: [
        {
          subtitle: "11:23 AM - 11:23 PM",
          title: "Design Task Model",
        },
        {
          subtitle: "11:23 AM - 11:23 PM",
          title: "Design RTO Cap TWO",
        },
        {
          subtitle: "11:23 AM - 11:23 PM",
          title: "Make Layers for Construct",
        },
      ],
    };
  },
  apollo: {
    links: gql`query{
      links{
        title
      }
    }`
  }
};
</script>

<style lang="scss">
.bd-left {
  height: 100vh;
  border-left: 1px solid #e2e2e2;
}
.time{
    line-height: 2.3rem;
}
</style>