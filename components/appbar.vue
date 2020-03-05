<template>
  <v-app-bar dark color="pink">
    <v-app-bar-nav-icon></v-app-bar-nav-icon>

    <v-toolbar-title>My Books</v-toolbar-title>

    <v-spacer></v-spacer>

    <nuxt-link :to="{ name: 'books-add' }">
      <v-btn icon>
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </nuxt-link>
  </v-app-bar>
</template>

<style scoped>
a {
  text-decoration: none;
}
</style>