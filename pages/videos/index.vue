<template>
  <div>
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1>Data Videos</h1>
          </div>
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item"><a href="">Home</a></li>
              <li class="breadcrumb-item active">Data Videos</li>
            </ol>
          </div>
        </div>
      </div>
      <!-- /.container-fluid -->
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
                <button
                  type="button"
                  class="btn btn-primary"
                  data-toggle="modal"
                  data-target="#modalForm"
                  @click="tambah"
                >
                  <i class="fa fa-plus"></i> Tambah
                </button>
                <br /><br />
                <vue-good-table
                  :columns="columns"
                  :rows="videoss"
                  :pagination-options="{
                    enabled: true,
                  }"
                  :search-options="{
                    enabled: true,
                  }"
                >
                  <template slot="table-row" slot-scope="props">
                    <span v-if="props.column.field == 'judul'">
                      {{ props.row.judul }}
                    </span>
                    <span v-if="props.column.field == 'deskripsi'">
                      {{ props.row.deskripsi }}
                    </span>
                    <span v-if="props.column.field == 'thumbnail'">
                      <img
                        :src="urlFile + '/img/' + props.row.thumbnail"
                        width="200px"
                        height="150px"
                      />
                    </span>
                   
                    <span v-if="props.column.field == 'option'">
                      <button
                        class="btn btn-sm btn-info"
                        data-toggle="modal"
                        data-target="#modalForm"
                        @click="edit(props.row)"
                      >
                        Edit
                      </button>
                      <button
                        class="btn btn-sm btn-danger"
                        @click="deleteData(props.row.id)"
                      >
                        Delete
                      </button>
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
            <button
              type="button"
              class="close"
              data-dismiss="modal"
              aria-label="Close"
            >
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <form @submit.prevent="onSubmit">
            <div class="modal-body">
              <div class="form-group">
                <label>Judul</label>
                <input
                  type="text"
                  class="form-control"
                  placeholder="Masukkan Judul"
                  v-model="formData.judul"
                />
              </div>
              <div class="form-group">
                <label>Deskripsi</label>
                <input
                  type="text"
                  class="form-control"
                  placeholder="Masukkan Deskripsi"
                  v-model="formData.deskripsi"
                />
              </div>
              <div class="form-group">
                <label for="exampleInputFile">Thumbnail</label>
                <div class="input-group">
                  <div class="custom-file">
                    <input
                      type="file"
                      ref="file"
                      class="form-control"
                      @change="onFileUpload"
                    />
                  </div>
                </div>
                <div v-if="titleModal == 'Edit Data'">
                  <p>*) Image sebelumnya</p>
                  <img
                    :src="urlFile + '/img/' + formData.thumbnailUrl"
                    width="100px"
                    height="100px"
                  />
                </div>
              </div>
              <div class="form-group">
                <label>link</label>
                <input
                  type="text"
                  class="form-control"
                  placeholder="Masukkan link"
                  v-model="formData.link"
                />
              </div>
			  <div v-if="titleModal == 'Edit Data' && formData.link != ''">
                  <p>*) Preview Video</p>
                  <iframe width="560" height="315" :src="'https://www.youtube-nocookie.com/embed/'+formData.link+'?si=THBFSuqnmzii2y_h&amp;controls=0'" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
                </div>
            </div>
            <div class="modal-footer justify-content-between">
              <button
                type="button"
                class="btn btn-default"
                ref="Close"
                data-dismiss="modal"
              >
                Close
              </button>
              <button type="submit" class="btn btn-primary">
                Save changes
              </button>
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
      idEdit: "",
      existFile: false,
      titleModal: "Tambah",
      formData: {
        judul: "",
        deskripsi: "",
        thumbnail: null,
        thumbnailUrl: "",
        link: "",
      },
      videoss: [],
      columns: [
        {
          label: "Judul",
          field: "judul",
        },
        {
          label: "Deskripsi",
          field: "deskripsi",
        },
        {
          label: "Thumbnail",
          field: "thumbnail",
        },
        {
          label: "Option",
          field: "option",
          width: "200px",
        },
      ],
    };
  },
  methods: {
    onFileUpload(event) {
      this.existFile = true;
      this.formData.thumbnail = event.target.files[0];
      console.log(this.formData.thumbnail);
    },
    async onSubmit() {
      let formReq = new FormData();
      formReq.append("judul", this.formData.judul);
      formReq.append("deskripsi", this.formData.deskripsi);
      if (this.existFile) {
        formReq.append("thumbnail", this.formData.thumbnail);
      }
      formReq.append("link", this.formData.link);
      if (this.isEdit == false) {
        await this.$axios
          .$post("/video", formReq, {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          })
          .then((res) => {
            console.log(res);
            this.$refs.Close.click();
            this.getData();
            this.$toast.info("Sukses simpan data !");
          })
          .catch((e) => {
            console.log(e);
            this.$toast.error("Gagal simpan data !");
          });
      } else {
        await this.$axios
          .$put("/video/" + this.idEdit, formReq, {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          })
          .then((res) => {
            console.log(res);
            this.$refs.Close.click();
            this.getData();
            this.$toast.info("Sukses update data !");
          })
          .catch((e) => {
            console.log(e);
            this.$toast.error("gagal update data !");
          });
      }
    },
    refresh() {
      this.$nuxt.refresh();
    },

    tambah() {
      this.reset();
      this.titleModal = "Tambah Data";
    },
    edit(data) {
      this.reset();
      this.isEdit = true;
      this.titleModal = "Edit Data";
      this.idEdit = data.id;
      this.formData.judul = data.judul;
      this.formData.deskripsi = data.deskripsi;
      this.formData.thumbnailUrl = data.thumbnail;
      this.formData.link = data.link;
    },
    reset() {
      this.formData.judul = "";
      this.formData.deskripsi = "";
      this.formData.thumbnail = null;
      this.formData.thumbnailUrl = "";
      this.formData.link = "";
      this.$refs.file.value = null;
      this.isEdit = false;
      this.idEdit = "";
      this.existFile = false;
    },
    async getData() {
      await this.$axios
        .$get("/video")
        .then((res) => {
          console.log(res);
          this.videoss = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    async deleteData(id) {
      if (confirm("Are you sure to delete this data ?")) {
        await this.$axios
          .$delete("/video/" + id)
          .then((res) => {
            console.log(res);
            this.getData();
            this.$toast.info("Sukses delete data !");
          })
          .catch((err) => {
            console.log(err);
            this.$toast.error("Gagal delete data !");
          });
      }
    },
  },
  mounted() {
    this.getData();
  },
};
</script>
