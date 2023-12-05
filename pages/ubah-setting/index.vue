<template>
    <div>
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <div class="container-fluid">
                <div class="row mb-2">
                    <div class="col-sm-6">

                    </div>
                    <div class="col-sm-6">
                        <ol class="breadcrumb float-sm-right">
                            <li class="breadcrumb-item"><a href="">Home</a></li>
                            <li class="breadcrumb-item active">Ubah Setting</li>
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
                    <div class="col-md-6">

                        <div class="card card-success">
                            <div class="card-header">
                                <h3 class="card-title">Ubah Setting</h3>
                            </div>
                            <div class="card-body">
                                <form @submit.prevent="onSubmit">
                                    <div class="modal-body">

                                        <div class="form-group">
                                            <label>Alamat</label>
                                            <textarea class="form-control" cols="30" rows="3" v-model="formData.alamat"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Facebook</label>
                                            <textarea class="form-control" cols="30" rows="3" v-model="formData.fb"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Instagram</label>
                                            <textarea class="form-control" cols="30" rows="3" v-model="formData.ig"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Twitter</label>
                                            <textarea class="form-control" cols="30" rows="3" v-model="formData.twitter"></textarea>
                                        </div>
                                        <div class="form-group">
                                            <label>Youtube</label>
                                            <textarea class="form-control" cols="30" rows="3" v-model="formData.youtube"></textarea>
                                        </div>

                                    </div>
                                    <div class="modal-footer justify-content-between">
                                        <button type="submit" class="btn btn-primary">Save changes</button>
                                    </div>
                                </form>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</template>

<script>

export default {
    data() {
        return {
            formData: {
                alamat: '',
                fb: '',
                ig: '',
                twitter: '',
                youtube: ''
            },

        }
    },
    methods: {

        async onSubmit() {
            await this.$axios
                .$put('/setting', this.formData)
                .then((res) => {
                    console.log(res)
                    this.$toast.info('Sukses ubah setting !')
                })
                .catch((e) => {
                    console.log(e.response)
                    this.$toast.error(e.response.data.meta.message)
                })
        },
        async getData() {
			await this.$axios
				.$get('/setting')
				.then((res) => {
					console.log(res)
					this.formData.alamat = res.data.alamat
					this.formData.ig = res.data.ig
					this.formData.twitter = res.data.twitter
					this.formData.fb = res.data.fb
					this.formData.youtube = res.data.youtube
				})
				.catch((err) => {
					console.log(err)
				})
		},

    },
    created(){
        this.getData()
    }



}
</script>