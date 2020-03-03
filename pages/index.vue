<template>
  <v-card
    max-width="400"
    class="mx-auto"
  >
    <v-system-bar
      color="pink darken-2"
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
          <v-card
            :color="item.color"
            dark
          >
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
    </v-container>
  </v-card>
</template>

<script>
import axios from 'axios'
export default {
  data: () => ({
    allItems: [],
    items: [],
  }),
  mounted() {
    this.getBooks()
  },
  methods: {
    getBooks() {
      axios('http://localhost:1234/books', {
        crossDomain: true
      }).then(({data}) => {
        this.allItems = data.Data
        data.Data.map((item, key) => {
          this.items.push(item)
        })
      })
    }
  }
}
</script>