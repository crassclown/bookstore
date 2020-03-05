<template>
  <v-card
    max-width="400"
    class="mx-auto"
  >
    <Appbar/>
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
import Appbar from '~/components/appbar'
import axios from 'axios'
export default {
  components: {
    Appbar
  },
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