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
            <v-subheader class="d-flex justify-space-between align-center">
                <div>
                  Fri, 12 Aug
                </div>
                <div>
                  <v-tooltip top>
                    <template v-slot:activator="{ on, attrs }">
                        <v-btn
                          outlined
                          x-small
                          fab
                          color="indigo"
                          v-bind="attrs"
                          v-on="on"
                          @click="dialog = true"
                        >
                          <v-icon>mdi-plus</v-icon>
                        </v-btn>
                      </template>
                    <span>Add task</span>
                  </v-tooltip>
                </div>
            </v-subheader>
            <v-divider></v-divider>
            <v-list-item v-for="folder in folders" :key="folder.title">
              <v-list-item-avatar>
                <v-tooltip top>
                    <template v-slot:activator="{ on, attrs }">
                        <v-btn
                          icon
                          medium
                          v-bind="attrs"
                          v-on="on"
                        >
                          <v-icon color="green">mdi-progress-check</v-icon>
                        </v-btn>
                      </template>
                    <span>Complete task</span>
                  </v-tooltip>
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

        <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
    >
     <v-card>
         <v-card-title>
            New Task
          </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                cols="12"
                sm="12"
                md="12"
              >
                <v-text-field
                  label="Task Name*"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            @click="dialog = false"
          >
            Close
          </v-btn>
          <v-btn
            color="blue darken-1"
            text
            @click="dialog = false"
          >
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    </v-container>
  </div>
</template>

<script>
import gql from 'graphql-tag'

export default {
  data: () => {
    return {
      dialog: false,
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
.v-dialog>.v-card>.v-card__subtitle, .v-dialog>.v-card>.v-card__text {
    padding: 0 24px 0px!important;
}
</style>