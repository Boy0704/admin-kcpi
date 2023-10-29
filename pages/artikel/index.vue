<template>
	<div>
		<!-- Content Header (Page header) -->
		<section class="content-header">
			<div class="container-fluid">
				<div class="row mb-2">
					<div class="col-sm-6">
						<h1>Data Artikel</h1>
					</div>
					<div class="col-sm-6">
						<ol class="breadcrumb float-sm-right">
							<li class="breadcrumb-item"><a href="">Home</a></li>
							<li class="breadcrumb-item active">Data Artikel</li>
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
								<vue-good-table :columns="columns" :rows="artikels" :pagination-options="{
									enabled: true,
								}" :search-options="{
	enabled: true
}">
									<template slot="table-row" slot-scope="props">
										<span v-if="props.column.field == 'image'">
											<img :src="urlFile + '/img/' + props.row.image" width="200px" height="150px">
										</span>
										<span v-if="props.column.field == 'title'">
											{{ props.row.title }}
										</span>
										<span v-if="props.column.field == 'content'">
											{{ props.row.content }}
										</span>
										<span v-if="props.column.field == 'id_category'">
											{{ props.row.id_category }}
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
								<label for="exampleInputFile">Image</label>
								<div class="input-group">
									<div class="custom-file">
										<input type="file" ref="file" class="form-control" @change="onFileUpload">
									</div>
								</div>
								<div v-if="titleModal == 'Edit Data'">
									<p>*) Image sebelumnya</p>
									<img :src="urlFile + '/img/' + formData.imageUrl" width="100px" height="100px">
								</div>
							</div>
							<div class="form-group">
								<label>Judul</label>
								<input type="text" class="form-control" placeholder="Masukkan Judul"
									v-model="formData.title">
							</div>
							<div class="form-group">
								<label>Isi Konten</label>
								<client-only>
									<quill-editor ref="editor" v-model="formData.content" :options="editorOption"
										@blur="onEditorBlur($event)" @focus="onEditorFocus($event)"
										@ready="onEditorReady($event)" />
								</client-only>
							</div>
							<div class="form-group">
								<label>Katory</label>
								<select class="form-control" v-model="formData.id_category">
									<option value="">Pilih Kategori</option>
									<option v-for="item in category" :value="item.id">{{ item.category }}</option>
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
				image: null,
				imageUrl: '',
				title: '',
				content: '',
				id_category: '',
				publish: false,
			},
			category: [],
			artikels: [],
			columns: [
				{
					label: 'Image',
					field: 'image',
				},
				{
					label: 'Judul',
					field: 'title',
				},
				{
					label: 'Isi Konten',
					field: 'content',
				},
				{
					label: 'Kategori',
					field: 'id_category',
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
			editorOption: {
				// Some Quill options...
				theme: 'snow',
				modules: {
					toolbar: [
						['bold', 'italic', 'underline', 'strike'],
						['blockquote', 'code-block'],
						[{ 'header': 1 }, { 'header': 2 }],
						[{ 'list': 'ordered' }, { 'list': 'bullet' }],
						[{ 'script': 'sub' }, { 'script': 'super' }],
						[{ 'indent': '-1' }, { 'indent': '+1' }],
						[{ 'direction': 'rtl' }],
						[{ 'size': ['small', false, 'large', 'huge'] }],
						[{ 'header': [1, 2, 3, 4, 5, 6, false] }],
						[{ 'font': [] }],
						[{ 'color': [] }, { 'background': [] }],
						[{ 'align': [] }],
						['clean'],
						['link', 'image', 'video']
					]
				}
			}
		}
	},
	methods: {
		onEditorBlur(editor) {
			console.log('editor blur!', editor)
		},
		onEditorFocus(editor) {
			console.log('editor focus!', editor)
		},
		onEditorReady(editor) {
			console.log('editor ready!', editor)
		},
		onFileUpload(event) {
			this.existFile = true
			this.formData.image = event.target.files[0]
			console.log(this.formData.image)
		},
		async onSubmit() {
			let formReq = new FormData()
			if (this.existFile) {
				formReq.append('file_img', this.formData.image)
			}
			formReq.append('title', this.formData.title)
			formReq.append('content', this.formData.content)
			formReq.append('id_category', this.formData.id_category)
			formReq.append('publish', this.formData.publish ? 'y' : 't')
			if (this.isEdit == false) {
				await this.$axios
					.$post('/artikel', formReq, {
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
					.$put('/artikel/' + this.idEdit, formReq, {
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
			this.formData.imageUrl = data.image
			this.formData.title = data.title
			this.formData.content = data.content
			this.formData.id_category = data.id_category
			this.formData.publish = data.publish == 'y' ? true : false
		},
		reset() {
			this.formData.image = null
			this.formData.imageUrl = ''
			this.formData.title = ''
			this.formData.content = ''
			this.formData.id_category = ''
			this.formData.publish = false
			this.$refs.file.value = null;
			this.isEdit = false
			this.idEdit = ''
			this.existFile = false
		},
		async getData() {
			await this.$axios
				.$get('/artikel')
				.then((res) => {
					console.log(res)
					this.artikels = res.data
				})
				.catch((err) => {
					console.log(err)
				})
		},
		async getCategory(){
			await this.$axios
					.$get('/category')
					.then((res) => {
						console.log(res)
						this.category = res.data
					})
					.catch((err) => {
						console.log(err)
						this.$toast.error('Gagal get category data !')
					})
		},
		async deleteData(id) {
			if (confirm('Are you sure to delete this data ?')) {
				await this.$axios
					.$delete('/artikel/' + id)
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
		this.getCategory()
	}



}
</script>