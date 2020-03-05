<template>
	<div id="bookAdd" class="container" data-app="true">
		<v-container>
			<v-form
			ref="form"
			v-model="valid"
			lazy-validation
			>
				<v-text-field
					v-model="title"
					:rules="titleRules"
					label="Title"
					required
				></v-text-field>

				<v-text-field
					v-model="desc"
					:rules="descRules"
					label="Description"
					required
				></v-text-field>

				<v-select
					v-model="author"
					:items="authors"
					:rules="[v => !!v || 'Author is required']"
					label="Author"
					item-text="name"
					item-value="id"
					required
				></v-select>

				<div class="text-center">
					<v-btn block color="primary" dark @click="addBook">Submit</v-btn>
				</div>
			</v-form>
		</v-container>
	</div>
</template>

<script>
import env from '~/env.json'
import axios from 'axios'
export default {
	name: 'BookAdd',
	data: () => ({
		valid: true,
		title: '',
		titleRules: [
			v => !!v || 'Title is required',
			v => (v && v.length >= 10) || 'Title must be more than 10 characters',
		],
		desc: '',
		descRules: [
			v => !!v || 'Description is required',
			v => (v && v.length >= 25) || 'Description must be more than 25 characters'
		],
		author: '',
		authors: [],
	}),
	mounted() {
		this.getAuthor()
	},
	methods: {
		getAuthor() {
			axios(''+env.BASE_URL_API+'/authors', {
				crossDomain: true
			}).then(({data}) => {
				let authors = data.Data
				authors.map((item, key) => {
					this.authors.push(item)
				})
			})
		},
		addBook() {
			
		},
		validate () {
			if (this.$refs.form.validate()) {
				this.snackbar = true
			}
		},
		reset () {
			this.$refs.form.reset()
		},
		resetValidation () {
			this.$refs.form.resetValidation()
		},
	},
}
</script>
