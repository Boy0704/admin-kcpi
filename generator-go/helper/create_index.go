package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func CreateIndex(json ConfigJson) bool {
	pathFolder, err := createFolder(json.FolderName)
	if err != nil {
		fmt.Println("gagal buat folder")
	}
	createFileIndex(pathFolder)
	err = writeFileIndex(pathFolder, json)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func createFolder(nameFolder string) (string, error) {
	path := filepath.Join("../pages/", nameFolder)
	_, err := os.Stat(path)

	if os.IsExist(err) {
		fmt.Println("your directory is already exist but it's ok")
		return path, err
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println("your directory is already exist but it's ok")
		return path, err
	}

	return path, nil
}

func createFileIndex(folderPath string) {
	absPath, err := filepath.Abs(folderPath + "/index.vue")

	if err != nil {
		fmt.Println(err)
	}

	filename, err := os.Create(absPath)
	if err != nil {
		log.Fatal("Cannot create a file please check your directory again")
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())

}

func writeFileIndex(folderPath string, json ConfigJson) error {
	absPath, err := filepath.Abs(folderPath + "/index.vue")
	if err != nil {
		fmt.Println(err)
		return err
	}
	file, err := os.OpenFile(absPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	template := `<template>
	<div>
		<!-- Content Header (Page header) -->
		<section class="content-header">
			<div class="container-fluid">
				<div class="row mb-2">
					<div class="col-sm-6">
						<h1>Data %s</h1>
					</div>
					<div class="col-sm-6">
						<ol class="breadcrumb float-sm-right">
							<li class="breadcrumb-item"><a href="">Home</a></li>
							<li class="breadcrumb-item active">Data %s</li>
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
								<vue-good-table :columns="columns" :rows="%ss" :pagination-options="{
									enabled: true,
								}" :search-options="{
	enabled: true
}">
									<template slot="table-row" slot-scope="props">
										%s
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

							%s

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
				%s
			},
			%ss: [],
			columns: [
				%s
				{
					label: 'Option',
					field: 'option',
					width: '200px',
				},
			],
		}
	},
	methods: {
		%s,
		async onSubmit() {
			let formReq = new FormData()
			%s
			if (this.isEdit == false) {
				await this.$axios
				.$post('%s', formReq, {
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
				.$put('%s/'+this.idEdit, formReq, {
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
			%s
		},
		reset(){
			%s
			this.$refs.file.value = null;
			this.isEdit = false
			this.idEdit = ''
			this.existFile = false
		},
		async getData() {
			await this.$axios
				.$get('%s')
				.then((res) => {
					console.log(res)
					this.%ss = res.data
				})
				.catch((err) => {
					console.log(err)
				})
		},
		async deleteData(id) {
			if (confirm('Are you sure to delete this data ?')) {
				await this.$axios
					.$delete('%s/' + id)
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
</script>`
	spanRTable := createSpanTable(json)
	form := createForm(json)
	varFormData := createVariableFormData(json)
	varColumnData := createColumnVariable(json)
	funcHandleFile := createHandleFile(json)
	request := createFormRequest(json)
	getEditData := createGetEditData(json)
	setResetData := createResetData(json)

	_, err = file.WriteString(fmt.Sprintf(template,
		capitalizeFirstLetter(json.FolderName),
		capitalizeFirstLetter(json.FolderName),
		json.FolderName,
		spanRTable,
		form,
		varFormData,
		json.FolderName,
		varColumnData,
		funcHandleFile,
		request,
		json.Endpoint,
		json.Endpoint,
		getEditData,
		setResetData,
		json.Endpoint,
		json.FolderName,
		json.Endpoint))
	if err != nil {
		return err
	}

	return nil
}

func createSpanTable(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {

		if json.Entity[i].DataType == "false" {
			template := `<span v-if="props.column.field == '%s'">
							<span v-if="props.row.%s == 'y'" class="badge badge-success">yes</span>
							<span v-else class="badge badge-danger">no</span>
						</span>`
			data := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, data)
		}

		if json.Entity[i].DataType == "" {
			template := `<span v-if="props.column.field == '%s'">
							{{ props.row.%s }}
						</span>`
			data := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, data)
		}

		if json.Entity[i].DataType == "image" {
			template := `<span v-if="props.column.field == '%s'">
							<img :src="urlFile + '/img/' + props.row.%s" width="200px" height="150px">
						</span>`
			data := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, data)
		}

	}
	return strings.Join(allString, "\n")
}

func createForm(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {
		//form type text
		if json.Entity[i].DataType == "" {
			template := `<div class="form-group">
							<label>%s</label>
							<input type="text" class="form-control" placeholder="Masukkan %s"
							  v-model="formData.%s">
						  </div>`
			form := fmt.Sprintf(template,
				json.Entity[i].Label,
				json.Entity[i].Label,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		// form type checkbox
		if json.Entity[i].DataType == "false" {
			template := `<div class="form-check">
							<input type="checkbox" :checked="formData.%s" class="form-check-input" v-model="formData.%s">
							<label class="form-check-label" for="exampleCheck1">%s</label>
						</div>`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName,
				json.Entity[i].Label)
			allString = append(allString, form)
		}

		// form type file upload
		if json.Entity[i].DataType == "image" {
			template := `<div class="form-group">
							<label for="exampleInputFile">%s</label>
							<div class="input-group">
								<div class="custom-file">
									<input type="file" ref="file" class="form-control" @change="onFileUpload">
								</div>
							</div>
							<div v-if="titleModal == 'Edit Data'">
								<p>*) Image sebelumnya</p>
								<img :src="urlFile + '/img/' + formData.%sUrl" width="100px" height="100px">
							</div>
						</div>`
			form := fmt.Sprintf(template,
				json.Entity[i].Label,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

	}
	return strings.Join(allString, "\n")
}

func createVariableFormData(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {

		if json.Entity[i].DataType == "" {
			template := `%s: '',`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "image" {
			template := `%s: null,
						%sUrl: '',`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "false" {
			template := `%s: false,`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

	}
	return strings.Join(allString, "\n")
}

func createColumnVariable(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {
		template := `{
					label: '%s',
					field: '%s',
				},`
		variable := fmt.Sprintf(template,
			json.Entity[i].Label,
			json.Entity[i].FieldName)
		allString = append(allString, variable)
	}
	return strings.Join(allString, "\n")
}

func createHandleFile(json ConfigJson) string {
	var variable string
	for i := 0; i < len(json.Entity); i++ {
		if json.Entity[i].DataType == "image" {
			template := `onFileUpload(event) {
					this.existFile = true
					this.formData.image = event.target.files[0]
					console.log(this.formData.%s)
				}`
			variable = fmt.Sprintf(template,
				json.Entity[i].FieldName)

		}
	}
	return variable
}

func createFormRequest(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {

		if json.Entity[i].DataType == "" {
			template := `formReq.append('%s', this.formData.%s)`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "image" {
			template := `if (this.existFile) {
							formReq.append('%s', this.formData.%s)
						}`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "false" {
			template := `formReq.append('%s', this.formData.%s ? 'y' : 't')`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

	}
	return strings.Join(allString, "\n")
}

func createGetEditData(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {

		if json.Entity[i].DataType == "" {
			template := `this.formData.%s = data.%s`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "image" {
			template := `this.formData.%sUrl = data.%s`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "false" {
			template := `this.formData.%s = data.%s == 'y' ? true : false`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

	}
	return strings.Join(allString, "\n")
}

func createResetData(json ConfigJson) string {
	var allString []string
	for i := 0; i < len(json.Entity); i++ {

		if json.Entity[i].DataType == "" {
			template := `this.formData.%s = ''`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "image" {
			template := `this.formData.%s = null
						this.formData.%sUrl = ''`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

		if json.Entity[i].DataType == "false" {
			template := `this.formData.%s = false`
			form := fmt.Sprintf(template,
				json.Entity[i].FieldName)
			allString = append(allString, form)
		}

	}
	return strings.Join(allString, "\n")
}

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}

	// Ubah rune pertama ke huruf besar
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
