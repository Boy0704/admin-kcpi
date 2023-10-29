<template>
	<div>
		<!-- Content Header (Page header) -->
		<section class="content-header">
			<div class="container-fluid">
				<div class="row mb-2">
					<div class="col-sm-6">
						<h1>Data Gallery</h1>
					</div>
					<div class="col-sm-6">
						<ol class="breadcrumb float-sm-right">
							<li class="breadcrumb-item"><a href="">Home</a></li>
							<li class="breadcrumb-item active">Data Gallery</li>
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
								<vue-good-table :columns="columns" :rows="gallerys" :pagination-options="{
									enabled: true,
								}" :search-options="{
	enabled: true
}">
									<template slot="table-row" slot-scope="props">
										<span v-if="props.column.field == 'media'">
											<video width="200" controls v-if="props.row.type == 'video'">
												<source :src="urlFile + '/video/' + props.row.media" type="video/mp4">
												Your browser does not support HTML video.
											</video>
											<img :src="urlFile + '/img/' + props.row.media" width="200px" height="150px"
												v-else>
										</span>
										<span v-if="props.column.field == 'type'">
											{{ props.row.type }}
										</span>
										<span v-if="props.column.field == 'publish'">
											<span v-if="props.row.publish == 'y'" class="badge badge-success">yes</span>
											<span v-else class="badge badge-danger">no</span>
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
								<label for="exampleInputFile">Media</label>
								<div class="input-group">
									<div class="custom-file">
										<input type="file" ref="file" class="form-control" @change="onFileUpload"
											accept="image/*, video/*">
									</div>
								</div>
								<div v-if="titleModal == 'Edit Data'">
									<p>*) File sebelumnya</p>
									<video width="200" controls v-if="formData.type == 'video'">
										<source :src="urlFile + '/video/' + formData.mediaUrl" type="video/mp4">
										Your browser does not support HTML video.
									</video>
									<img :src="urlFile + '/img/' + formData.mediaUrl" width="200px" height="150px" v-else>
								</div>
							</div>
							<div class="form-group">
								<label>Type</label>
								<select class="form-control" v-model="formData.type">
									<option value="">Type</option>
									<option value="image">Image</option>
									<option value="video">Video</option>
								</select>
							</div>
							<div class="form-check">
								<input type="checkbox" :checked="formData.publish" class="form-check-input"
									v-model="formData.publish">
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
				media: null,
				mediaUrl: '',
				type: '',
				publish: false,
			},
			gallerys: [],
			columns: [
				{
					label: 'Media',
					field: 'media',
				},
				{
					label: 'Type',
					field: 'type',
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
			this.formData.media = event.target.files[0]
			console.log(this.formData.media)
		},
		async onSubmit() {
			let formReq = new FormData()
			if (this.existFile) {
				formReq.append('file_img', this.formData.media)
			}
			formReq.append('type', this.formData.type)
			formReq.append('publish', this.formData.publish ? 'y' : 't')
			if (this.isEdit == false) {
				await this.$axios
					.$post('/gallery', formReq, {
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
					.$put('/gallery/' + this.idEdit, formReq, {
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
			this.formData.mediaUrl = data.media
			this.formData.type = data.type
			this.formData.publish = data.publish == 'y' ? true : false
		},
		reset() {
			this.formData.media = null
			this.formData.mediaUrl = ''
			this.formData.type = ''
			this.formData.publish = false
			this.$refs.file.value = null;
			this.isEdit = false
			this.idEdit = ''
			this.existFile = false
		},
		async getData() {
			await this.$axios
				.$get('/gallery')
				.then((res) => {
					console.log(res)
					this.gallerys = res.data
				})
				.catch((err) => {
					console.log(err)
				})
		},
		async deleteData(id) {
			if (confirm('Are you sure to delete this data ?')) {
				await this.$axios
					.$delete('/gallery/' + id)
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