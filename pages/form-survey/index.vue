<template>
	<div>
		<!-- Content Header (Page header) -->
		<section class="content-header">
			<div class="container-fluid">
				<div class="row mb-2">
					<div class="col-sm-6">
						<h1>Data Form-survey</h1>
					</div>
					<div class="col-sm-6">
						<ol class="breadcrumb float-sm-right">
							<li class="breadcrumb-item"><a href="">Home</a></li>
							<li class="breadcrumb-item active">Data Form-survey</li>
						</ol>
					</div>
				</div>
			</div><!-- /.container-fluid -->
		</section>
		<!-- Main content -->
		<section class="content">
			<div class="container-fluid">
				<div class="row">
					<!-- left column -->
					<div class="col-md-12">
						<div class="card">
							<div class="card-header">
								<h3 class="card-title"></h3>
							</div>
							<!-- /.card-header -->
							<div class="card-body">
								
								<br /><br />
								<vue-good-table :columns="columns" :rows="formSurveys" :pagination-options="{
									enabled: true,
								}" :search-options="{
	enabled: true
}">
									<template slot="table-row" slot-scope="props">
										<span v-if="props.column.field == 'url'">
											{{ props.row.url }}
										</span>
										<span v-if="props.column.field == 'option'">
											<button class="btn btn-sm btn-info" data-toggle="modal" data-target="#modalForm"
												@click="edit(props.row)">Edit</button>
											
										</span>
									</template>

								</vue-good-table>
							</div>
						</div>

					</div>
				</div>
			</div>
		</section>

		<div class="modal fade" id="modalForm">
			<div class="modal-dialog modal-lg">
				<div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title">{{ titleModal }}</h4>
						<button type="button" class="close" data-dismiss="modal" aria-label="Close">
							<span aria-hidden="true">&times;</span>
						</button>
					</div>
					<form @submit.prevent="onSubmit">
						<div class="modal-body">

							<div class="form-group">
								<label>Url</label>
								<input type="text" class="form-control" placeholder="Masukkan Url" v-model="formData.url">
							</div>

						</div>
						<div class="modal-footer justify-content-between">
							<button type="button" class="btn btn-default" ref="Close" data-dismiss="modal">Close</button>
							<button type="submit" class="btn btn-primary">Save changes</button>
						</div>
					</form>
				</div>
				<!-- /.modal-content -->
			</div>
			<!-- /.modal-dialog -->
		</div>
		<!-- /.modal -->

	</div>
</template>

<script>

export default {
	data() {
		return {
			urlFile: this.$config.FileUrl,
			isEdit: false,
			idEdit: '',
			existFile: false,
			titleModal: 'Tambah',
			formData: {
				url: '',
			},
			formSurveys: [],
			columns: [
				{
					label: 'Url',
					field: 'url',
				},
				{
					label: 'Option',
					field: 'option',
					width: '200px',
				},
			],
		}
	},
	methods: {

		async onSubmit() {
			let formReq = new FormData()
			formReq.append('url', this.formData.url)
			if (this.isEdit == false) {
				await this.$axios
					.$post('/form-survey', formReq, {
						headers: {
							'Content-Type': 'multipart/form-data'
						}
					})
					.then((res) => {
						console.log(res)
						this.$refs.Close.click();
						this.getData()
						this.$toast.info('Sukses simpan data !')
					})
					.catch((e) => {
						console.log(e)
						this.$toast.error('Gagal simpan data !')
					})
			} else {
				await this.$axios
					.$put('/form-survey/' + this.idEdit, formReq, {
						headers: {
							'Content-Type': 'multipart/form-data'
						}
					})
					.then((res) => {
						console.log(res)
						this.$refs.Close.click();
						this.getData()
						this.$toast.info('Sukses update data !')
					})
					.catch((e) => {
						console.log(e)
						this.$toast.error('gagal update data !')
					})
			}
		},
		refresh() {
			this.$nuxt.refresh()
		},

		tambah() {
			this.reset()
			this.titleModal = "Tambah Data"
		},
		edit(data) {
			this.reset()
			this.isEdit = true
			this.titleModal = "Edit Data"
			this.idEdit = data.id
			this.formData.url = data.url
		},
		reset() {
			this.formData.url = ''
			
			this.isEdit = false
			this.idEdit = ''
			this.existFile = false
		},
		async getData() {
			await this.$axios
				.$get('/form-surveys')
				.then((res) => {
					console.log(res)
					this.formSurveys = res.data
				})
				.catch((err) => {
					console.log(err)
				})
		},
		async deleteData(id) {
			if (confirm('Are you sure to delete this data ?')) {
				await this.$axios
					.$delete('/form-survey/' + id)
					.then((res) => {
						console.log(res)
						this.getData()
						this.$toast.info('Sukses delete data !')
					})
					.catch((err) => {
						console.log(err)
						this.$toast.error('Gagal delete data !')
					})
			}
		}
	},
	mounted() {
		this.getData()
	}



}
</script>