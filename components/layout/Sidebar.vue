<template>
	<!-- Main Sidebar Container -->
	<aside class="main-sidebar sidebar-dark-primary elevation-4">
		<!-- Brand Logo -->
		<NuxtLink to="/" class="brand-link">
			<img src="/disttemplate/img/AdminLTELogo.png" alt="AdminLTE Logo" class="brand-image img-circle elevation-3"
				style="opacity: .8">
			<span class="brand-text font-weight-light">Admin Panel</span>
		</NuxtLink>

		<!-- Sidebar -->
		<div class="sidebar">
			<!-- Sidebar user (optional) -->
			<div class="user-panel mt-3 pb-3 mb-3 d-flex">
				<div class="image">
					<img src="/disttemplate/img/user2-160x160.jpg" class="img-circle elevation-2" alt="User Image">
				</div>
				<div class="info">
					<a href="#" class="d-block">{{ this.$auth.user.name }}</a>
				</div>
			</div>

			<!-- SidebarSearch Form -->
			<div class="form-inline">
				<div class="input-group" data-widget="sidebar-search">
					<input class="form-control form-control-sidebar" type="search" placeholder="Search" aria-label="Search">
					<div class="input-group-append">
						<button class="btn btn-sidebar">
							<i class="fas fa-search fa-fw"></i>
						</button>
					</div>
				</div>
			</div>

			<!-- Sidebar Menu -->
			<nav class="mt-2">
				<ul class="nav nav-pills nav-sidebar flex-column" data-widget="treeview" role="menu" data-accordion="false">
					<li class="nav-item">
						<NuxtLink to="slider" class="nav-link">
							<i class="nav-icon far fa-image"></i>
							<p>
								Data Slider
							</p>
						</NuxtLink>
					</li>

					<li class="nav-item">
						<NuxtLink to="category" class="nav-link">
							<i class="nav-icon fas fa-table"></i>
							<p>
								Data Kategori
							</p>
						</NuxtLink>
					</li>

					<li class="nav-item">
						<NuxtLink to="artikel" class="nav-link">
							<i class="nav-icon fas fa-newspaper"></i>
							<p>
								Data Artikel
							</p>
						</NuxtLink>
					</li>

					<li class="nav-item">
						<NuxtLink to="content-statis" class="nav-link">
							<i class="nav-icon fas fa-columns"></i>
							<p>
								Data Konten Statis
							</p>
						</NuxtLink>
					</li>
					
					<li class="nav-item">
						<NuxtLink to="istilah" class="nav-link">
							<i class="nav-icon fas fa-database"></i>
							<p>
								Data Istilah
							</p>
						</NuxtLink>
					</li>
					<li class="nav-item">
						<NuxtLink to="library" class="nav-link">
							<i class="nav-icon fas fa-database"></i>
							<p>
								Data Library
							</p>
						</NuxtLink>
					</li>
					<li class="nav-item">
						<NuxtLink to="stackholder" class="nav-link">
							<i class="nav-icon fas fa-database"></i>
							<p>
								Data Info Penggiat
							</p>
						</NuxtLink>
					</li>

					<li class="nav-item">
						<NuxtLink to="ubah-setting" class="nav-link">
							<i class="nav-icon fas fa-setting"></i>
							<p>
								Setting
							</p>
						</NuxtLink>
					</li>
					<li class="nav-item">
						<NuxtLink to="ubah-password" class="nav-link">
							<i class="nav-icon fas fa-share-square"></i>
							<p>
								Ubah Password
							</p>
						</NuxtLink>
					</li>

					<li class="nav-item">
						<a href="javascript:;" @click="logout" class="nav-link">
							<i class="nav-icon fas fa-sign-out-alt"></i>
							<p>
								Logout
							</p>
						</a>
					</li>

				</ul>
			</nav>
			<!-- /.sidebar-menu -->
		</div>
		<!-- /.sidebar -->
	</aside>
</template>

<script>
export default {
	methods: {
		async logout() {
			if (confirm('Are you sure to logout ?')) {
				await this.$auth.logout()
				console.log("logout")
				this.$router.push({
						name: 'login',
					})
				this.$toast.info("Sukses logout !")
			}
            
        },
	}
}
</script>