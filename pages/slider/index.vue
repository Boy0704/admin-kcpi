<template>
	<div>
		<!-- Content Header (Page header) -->
		<section class="content-header">
			<div class="container-fluid">
				<div class="row mb-2">
					<div class="col-sm-6">
						<h1>Data Slider</h1>
					</div>
					<div class="col-sm-6">
						<ol class="breadcrumb float-sm-right">
							<li class="breadcrumb-item"><a href="">Home</a></li>
							<li class="breadcrumb-item active">Data Slider</li>
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
								<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#modalForm"
									@click="tambah"><i class="fa fa-plus"></i> Tambah</button>
								<br /><br />
								<vue-good-table :columns="columns" :rows="sliders" :pagination-options="{
									enabled: true,
								}" :search-options="{
	enabled: true
}">
									<template slot="table-row" slot-scope="props">
										<span v-if="props.column.field == 'text'">
											{{ props.row.text }}
										</span>
										<span v-if="props.column.field == 'publish'">
											<span v-if="props.row.publish == 'y'" class="badge badge-success">yes</span>
											<span v-else class="badge badge-danger">no</span>
										</span>
										<span v-if="props.column.field == 'url'">
											<img :src="urlFile + '/img/' + props.row.url" width="200px" height="150px">
										</span>
										<span v-if="props.column.field == 'option'">
											<button class="btn btn-sm btn-info" data-toggle="modal" data-target="#modalForm"
												@click="edit(props.row)">Edit</button>
											<button class="btn btn-sm btn-danger"
												@click="deleteData(props.row.id)">Delete</button>
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
								<label>Judul</label>
								<input type="text" class="form-control" placeholder="Masukkan Judul"
									v-model="formData.text">
							</div>
							<div class="form-group">
								<label for="exampleInputFile">Gambar</label>
								<div class="input-group">
									<div class="custom-file">
										<input type="file" ref="file" class="form-control" @change="onFileUpload">
									</div>
								</div>
								<div v-if="titleModal == 'Edit Data'">
									<p>*) File sebelumnya</p>
									<img :src="urlFile + '/img/' + formData.imgUrl" width="100px" height="100px">
								</div>
							</div>
							<div class="form-check">
								<input type="checkbox" :checked="formData.publish" class="form-check-input" v-model="formData.publish">
								<label class="form-check-label" for="exampleCheck1">Publish</label>
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
				image: null,
				imgUrl: '',
				text: '',
				publish: false
			},
			sliders: [],
			columns: [
				{
					label: 'Judul',
					field: 'text',
				},
				{
					label: 'Image',
					field: 'url',
				},
				{
					label: 'Publish',
					field: 'publish',
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
		onFileUpload(event) {
			this.existFile = true
			this.formData.image = event.target.files[0]
			console.log(this.formData.image)
		},
		async onSubmit() {
			let formReq = new FormData()
			if (this.existFile) {
				formReq.append('file_slider', this.formData.image)
			}
			formReq.append('text', this.formData.text)
			formReq.append('publish', this.formData.publish ? 'y' : 't')
			if (this.isEdit == false) {
				await this.$axios
				.$post('/slider', formReq, {
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
				.$put('/slider/'+this.idEdit, formReq, {
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
			this.idEdit = data.id
			this.formData.text = data.text
			this.formData.imgUrl = data.url
			this.formData.publish = data.publish == 'y' ? true : false
			this.titleModal = "Edit Data"
		},
		reset(){
			this.formData.image = null
			this.formData.text = ''
			this.formData.imgUrl = ''
			this.formData.publish = false
			this.$refs.file.value = null;
			this.isEdit = false
			this.idEdit = ''
			this.existFile = false
		},
		async getData() {
			await this.$axios
				.$get('/slider')
				.then((res) => {
					console.log(res)
					this.sliders = res.data
				})
				.catch((err) => {
					console.log(err)
				})
		},
		async deleteData(id) {
			if (confirm('Are you sure to delete this data ?')) {
				await this.$axios
					.$delete('/slider/' + id)
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