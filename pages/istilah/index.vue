<template>
	<div>
		<!-- Content Header (Page header) -->
		<section class="content-header">
			<div class="container-fluid">
				<div class="row mb-2">
					<div class="col-sm-6">
						<h1>Data Istilah</h1>
					</div>
					<div class="col-sm-6">
						<ol class="breadcrumb float-sm-right">
							<li class="breadcrumb-item"><a href="">Home</a></li>
							<li class="breadcrumb-item active">Data Istilah</li>
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
								<vue-good-table :columns="columns" :rows="istilahs" :pagination-options="{
									enabled: true,
								}" :search-options="{
	enabled: true
}">
									<template slot="table-row" slot-scope="props">
										<span v-if="props.column.field == 'istilah'">
											{{ props.row.istilah }}
										</span>
										<span v-if="props.column.field == 'definisi'">
											{{ props.row.definisi }}
										</span>
										<span v-if="props.column.field == 'referensi'">
											{{ props.row.referensi }}
										</span>
										<span v-if="props.column.field == 'tahun'">
											{{ props.row.tahun }}
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
								<label>Istilah</label>
								<input type="text" class="form-control" placeholder="Masukkan Istilah"
									v-model="formData.istilah">
							</div>
							<div class="form-group">
								<label>Definisi</label>
								<input type="text" class="form-control" placeholder="Masukkan Definisi"
									v-model="formData.definisi">
							</div>
							<div class="form-group">
								<label>Referensi</label>
								<input type="text" class="form-control" placeholder="Masukkan Referensi"
									v-model="formData.referensi">
							</div>
							<div class="form-group">
								<label>Tahun</label>
								<input type="text" class="form-control" placeholder="Masukkan Tahun"
									v-model="formData.tahun">
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
				istilah: '',
				definisi: '',
				referensi: '',
				tahun: '',
			},
			istilahs: [],
			columns: [
				{
					label: 'Istilah',
					field: 'istilah',
				},
				{
					label: 'Definisi',
					field: 'definisi',
				},
				{
					label: 'Referensi',
					field: 'referensi',
				},
				{
					label: 'Tahun',
					field: 'tahun',
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
			if (this.isEdit == false) {
				await this.$axios
					.$post('/istilah', this.formData)
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
					.$put('/istilah/' + this.idEdit, this.formData)
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
			this.formData.istilah = data.istilah
			this.formData.definisi = data.definisi
			this.formData.referensi = data.referensi
			this.formData.tahun = data.tahun
		},
		reset() {
			this.formData.istilah = ''
			this.formData.definisi = ''
			this.formData.referensi = ''
			this.formData.tahun = ''
			this.isEdit = false
			this.idEdit = ''
			this.existFile = false
		},
		async getData() {
			await this.$axios
				.$get('/istilah')
				.then((res) => {
					console.log(res)
					this.istilahs = res.data
				})
				.catch((err) => {
					console.log(err)
				})
		},
		async deleteData(id) {
			if (confirm('Are you sure to delete this data ?')) {
				await this.$axios
					.$delete('/istilah/' + id)
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