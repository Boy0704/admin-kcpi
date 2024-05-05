<template>
	<div class="wrapper">
		<!-- Navbar -->
		<app-header v-if="isAuthenticated"/>
		<!-- /.navbar -->
		<app-sidebar v-if="isAuthenticated"/>
		<!-- Main Sidebar Container -->

		<!-- Content Wrapper. Contains page content -->
		<div class="content-wrapper" v-if="isAuthenticated">
			<Nuxt />
		</div>
		<!-- /.content-wrapper -->
		<footer class="main-footer">
			<strong>Copyright &copy; 2023 Knowledge Centre Perubahan Iklim</strong> 
		</footer>

		<!-- Control Sidebar -->
		<aside class="control-sidebar control-sidebar-dark">
			<!-- Control sidebar content goes here -->
		</aside>
		<!-- /.control-sidebar -->
	</div>
	<!-- ./wrapper -->
</template>

<script>
import Sidebar from '../components/layout/Sidebar.vue';
import Header from '../components/layout/Header.vue';
import { mapGetters } from 'vuex'
export default {
	components: {
		'app-sidebar': Sidebar,
		'app-header': Header
	},
	computed: {
		...mapGetters(['isAuthenticated', 'loggedInUser']),
	},
}
</script>