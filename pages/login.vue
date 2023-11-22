<template>
	<div class="card card-outline card-primary">
		<div class="card-header text-center">
			<a href="../../index2.html" class="h1"><b>Admin</b> Panel</a>
		</div>
		<div class="card-body">
			<p class="login-box-msg">Sign in to start your session</p>
			<div class="input-group mb-3">
				<input type="text" class="form-control" placeholder="Username" v-on:keyup.enter="onEnter" v-model="login.username" required>
				<div class="input-group-append">
					<div class="input-group-text">
						<span class="fas fa-user"></span>
					</div>
				</div>
			</div>
			<div class="input-group mb-3">
				<input type="password" class="form-control" placeholder="Password" v-on:keyup.enter="onEnter" v-model="login.password" required>
				<div class="input-group-append">
					<div class="input-group-text">
						<span class="fas fa-lock"></span>
					</div>
				</div>
			</div>
			<div class="row">

				<div class="col-12">
					<button @click="userLogin" class="btn btn-primary btn-block">Sign In</button>
				</div>
				<!-- /.col -->
			</div>

		</div>
		<!-- /.card-body -->
	</div>
</template>

<script>
export default {
	layout: 'auth',
	data() {
		return {
			loading: false,
			login: {
				username: '',
				password: '',
			},
		}
	},
	methods: {
		onEnter(){
			this.userLogin()
		},
		async userLogin() {
			this.loading = true
			try {
				let response = await this.$auth.loginWith('local', { data: this.login })
				this.$auth.setUser(response.data.data)
				console.log(response)
				this.loading = false
				this.$toast.info("Sukses login !")
				this.$router.push({
						name: 'index',
					})
			} catch (err) {
				console.log(err.response)
				this.loading = false
				this.$toast.error("Gagal login !")
			}
		},
	},
}
</script>