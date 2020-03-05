<template>
  <v-card
    max-width="400"
    class="mx-auto"
  >
    <v-system-bar
      color="pink"
      dark
    >
      <v-spacer></v-spacer>

      <v-icon>mdi-window-minimize</v-icon>

      <v-icon>mdi-window-maximize</v-icon>

      <v-icon>mdi-close</v-icon>
    </v-system-bar>

    <v-app-bar
      dark
      color="pink"
    >
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>My Books</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </v-app-bar>

    <v-container>
      <v-row dense>
        <v-col
          v-for="(item, i) in items"
          :key="i"
          cols="12"
        >
          <v-card>
            <div class="d-flex flex-no-wrap justify-space-between">
              <div>
                <v-card-title
                  class="headline"
                  v-text="item.title"
                ></v-card-title>

                <v-card-subtitle v-text="item.author"></v-card-subtitle>
              </div>

              <v-card-text>
                {{ item.description }}
              </v-card-text>
            </div>
          </v-card>
        </v-col>
      </v-row>
      <div class="text-center">
        <v-btn block color="primary" dark @click="loadMore">Load More</v-btn>
      </div>
    </v-container>
  </v-card>
</template>

<script>
import env from '~/env.json'
import axios from 'axios'
export default {
  data: () => ({
    allItems: [],
    items: [],
    current: 9
  }),
  mounted() {
    this.getBooks()
  },
  methods: {
    getBooks() {
      axios(''+env.BASE_URL_API+'/books', {
        crossDomain: true
      }).then(({data}) => {
        this.allItems = data.Data
        data.Data.map((item, key) => {
          if(this.items.length < 9) {
            this.items.push(item)
          }
        })
      }).catch((e) => {
        error({ statusCode: 404, message: 'Books not found' })
      })
    },
    loadMore() {
      this.items = []
      this.current += 9
      this.allItems.map((item, key) => {
        this.items.length < this.current ? this.items.push(item) : ''
      })
    }
  }
}
</script>