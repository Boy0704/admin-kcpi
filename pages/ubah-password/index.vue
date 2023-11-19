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
                            <li class="breadcrumb-item active">Ubah Password</li>
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
                                <h3 class="card-title">Ubah Password</h3>
                            </div>
                            <div class="card-body">
                                <form @submit.prevent="onSubmit">
                                    <div class="modal-body">

                                        <div class="form-group">
                                            <label>Password Lama</label>
                                            <input-password class-form-control="form-control"
                                                placeholder="Masukkan Password Lama" v-model="formData.passwordOld" />
                                        </div>
                                        <div class="form-group">
                                            <label>Password Baru</label>
                                            <input-password class-form-control="form-control"
                                                placeholder="Masukkan Password Baru" v-model="formData.newPassword" />
                                        </div>
                                        <div class="form-group">
                                            <label>Konfirmasi Password Baru</label>
                                            <input-password :class-form-control="form_control2"
                                                placeholder="Masukkan Password Baru" v-model="formData.newPasswordAgain" />
                                            <span class="text-red">{{ error }}</span>
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

import Password from '../../components/form/Password.vue';

export default {
    components: {
        'input-password': Password,
    },
    data() {
        return {
            form_control2: "form-control",
            error: '',
            formData: {
                passwordOld: '',
                newPassword: '',
                newPasswordAgain: ''
            },

        }
    },
    methods: {

        async onSubmit() {
            await this.$axios
                .$post('/reset-password', this.formData)
                .then((res) => {
                    console.log(res)
                    this.$toast.info('Sukses ubah password !')
                    this.logout()
                })
                .catch((e) => {
                    console.log(e)
                    this.$toast.error('Gagal ubah password !')
                })
        },
        async logout() {
            await this.$auth.logout()
				console.log("logout")
				this.$router.push({
						name: 'login',
					})
				this.$toast.info("Sukses logout !")
        }

    },
    mounted() {
    },
    watch: {
        'formData.newPasswordAgain': function (newValue) {
            if (newValue != this.formData.newPassword) {
                this.form_control2 = "form-control is-invalid"
                this.error = "password tidak sama"
            } else {
                this.form_control2 = "form-control"
                this.error = ""
            }
        }
    }



}
</script>