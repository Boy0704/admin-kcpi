<template lang="">
  <div>
    <section class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1>Data Survey</h1>
          </div>
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item"><a href="">Home</a></li>
              <li class="breadcrumb-item active">Data Survey</li>
            </ol>
          </div>
        </div>
      </div>
    </section>
    <section class="content">
      <div class="container-fluid">
        <div class="row">
          <div class="col-md-12">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Survey Kepuasan</h3>
              </div>

              <div class="card-body">
                <div class="row">
                  <div class="col-md-4">
                    <input
                      type="date"
                      class="form-control"
                      name="start_date"
                      placeholder="Start Date"
                      v-model="start_date"
                    />
                  </div>
                  <div class="col-md-4">
                    <input
                      type="date"
                      class="form-control"
                      name="end_date"
                      placeholder="End Date"
                      v-model="end_date"
                    />
                  </div>
                  <div class="col-md-4">
                    <button
                      type="button"
                      class="btn btn-primary btn-md"
                      @click="getDataSurvey()"
                    >
                      <i class="fa fa-search"></i> Filter
                    </button>
                  </div>
                </div>

                <br />
                <vue-good-table
                  :columns="columns"
                  :rows="result"
                  :pagination-options="{
                    enabled: true,
                  }"
                  :search-options="{
                    enabled: true,
                  }"
                >
                  <template slot="table-row" slot-scope="props">
                    <span v-if="props.column.field == 'nama'">
                      {{ props.row.nama }}
                    </span>
                    <span v-if="props.column.field == 'email'">
                      {{ props.row.email }}
                    </span>
                    <span v-if="props.column.field == 'jenis_kelamin'">
                      {{ props.row.jenis_kelamin }}
                    </span>
                    <span v-if="props.column.field == 'usia'">
                      {{ props.row.usia }}
                    </span>
                    <span v-if="props.column.field == 'pekerjaan'">
                      {{ props.row.pekerjaan }}
                    </span>
                    <span v-if="props.column.field == 'kesimpulan'">
                      {{ props.row.kesimpulan }}
                    </span>
                    <span v-if="props.column.field == 'saran'">
                      {{ props.row.saran }}
                    </span>
                    <span v-if="props.column.field == 'kunjungan'">
                      {{ props.row.kunjungan }}
                    </span>
                    <span v-if="props.column.field == 'nilai_website'">
                      {{ props.row.nilai_website }}
                    </span>
                    <span
                      v-if="props.column.field == 'nilai_website_info_aktual'"
                    >
                      {{ props.row.nilai_website_info_aktual }}
                    </span>
                    <span
                      v-if="
                        props.column.field == 'nilai_website_kemudahan_navigasi'
                      "
                    >
                      {{ props.row.nilai_website_kemudahan_navigasi }}
                    </span>
                    <span
                      v-if="
                        props.column.field == 'nilai_website_kualitas_konten'
                      "
                    >
                      {{ props.row.nilai_website_kualitas_konten }}
                    </span>
                    <span
                      v-if="
                        props.column.field == 'nilai_website_gangguan_teknis'
                      "
                    >
                      {{ props.row.nilai_website_gangguan_teknis }}
                    </span>
                    <span
                      v-if="props.column.field == 'mengetahui_website_dari'"
                    >
                      <p>
                        <span v-if="props.row.tau_dari_media_sosial"
                          >Media Sosial,</span
                        >
                        <span v-if="props.row.tau_dari_teman"
                          >Rekomendasi teman/kolaborator,</span
                        >
                        <span v-if="props.row.tau_dari_internet"
                          >Pencarian Internet,</span
                        >
                        <span v-if="props.row.tau_dari_lainnya">Lainnya,</span>
                      </p>
                    </span>
                    <span v-if="props.column.field == 'tujuan'">
                      <p>
                        <span v-if="props.row.tujuan_cari_data_publikasi"
                          >Mencari data atau publikasi,</span
                        >
                        <span v-if="props.row.tujuan_melihat_berita"
                          >Melihat berita atau update terbaru,</span
                        >
                        <span v-if="props.row.tujuan_cari_info"
                          >Mencari informasi tentang perubahan iklim,</span
                        >
                        <span v-if="props.row.tujuan_lainnya">Lainnya,</span>
                      </p>
                    </span>
                  </template>
                </vue-good-table>
              </div>

              <div class="card-footer">
                <h3 class="card-title">Laki-Laki : {{ sum_male }}</h3>
                <br />
                <h3 class="card-title">Perempuan : {{ sum_female }}</h3>
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
      start_date: new Date().toISOString().slice(0,10),
      end_date: new Date().toISOString().slice(0,10),
      result: [],
      sum_male: 0,
      sum_female: 0,
      columns: [
        {
          label: "Nama",
          field: "nama",
          width: "200px",
        },
        {
          label: "Email",
          field: "email",
          width: "200px",
        },
        {
          label: "Jenis Kelamin",
          field: "jenis_kelamin",
          width: "200px",
        },
        {
          label: "Usia",
          field: "usia",
          width: "200px",
        },
        {
          label: "Pekerjaan",
          field: "pekerjaan",
          width: "200px",
        },
        {
          label: "Kesimpulan",
          field: "kesimpulan",
          width: "200px",
        },
        {
          label: "Saran",
          field: "saran",
          width: "200px",
        },
        {
          label: "Kunjungan",
          field: "kunjungan",
          width: "200px",
        },
        {
          label: "Penilaian Website",
          field: "nilai_website",
          width: "200px",
        },
        {
          label: "Info Aktual",
          field: "nilai_website_info_aktual",
          width: "200px",
        },
        {
          label: "Kemudahan Navigasi",
          field: "nilai_website_kemudahan_navigasi",
          width: "200px",
        },
        {
          label: "Kualitas Konten",
          field: "nilai_website_kualitas_konten",
          width: "200px",
        },
        {
          label: "Mengetahui Website Dari",
          field: "mengetahui_website_dari",
          width: "200px",
        },
        {
          label: "Gangguan Teknis",
          field: "nilai_website_gangguan_teknis",
          width: "200px",
        },
        {
          label: "Tujuan",
          field: "tujuan",
          width: "200px",
        },
      ],
    };
  },
  mounted(){
    this.getDataSurvey();
  },
  methods: {
    getDataSurvey: async function () {
      await this.$axios
        .$get(
          `/survey_kepuasan?start_date=${this.start_date}&end_date=${this.end_date}`
        )
        .then((res) => {
          this.result = res.data;
          this.sum_female = 0;
          this.sum_male = 0;
          res.data.forEach((survey) => {
            if (survey.jenis_kelamin === "Laki-Laki") {
              this.sum_male++;
            } else if (survey.jenis_kelamin === "Perempuan") {
              this.sum_female++;
            }
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
<style lang=""></style>
