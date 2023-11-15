<template>
  <div id="app">
    <editor api-key="4mo39ri6dgnfnyfqwhr6nhicdjgg3nckwd3ruoyr8sa3d5z7" :init="{
      height: 500,
      menubar: true,
      images_upload_handler: this.uploadImageHandler,
      plugins: [
        'advlist autolink lists link image charmap print preview anchor',
        'searchreplace visualblocks code fullscreen',
        'insertdatetime media table paste code help wordcount'
      ],
      toolbar:
        'undo redo | formatselect | bold italic backcolor | \
              alignleft aligncenter alignright alignjustify | \
              bullist numlist outdent indent | removeformat | help'
    }" />
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'

export default {
  name: 'app',
  components: {
    'editor': Editor
  },
  data() {
    return {
      urlFile: this.$config.FileUrl,
    }
  },
  methods: {
    async uploadImageHandler(blobInfo, success, failure) {
        const formData = new FormData();
        formData.append('file_img', blobInfo.blob(), blobInfo.filename());
        this.$axios
          .post('/upload-img', formData, {
            headers: {
              'Content-Type': 'multipart/form-data',
            },
          })
          .then((response) => {
            success(this.urlFile+"/img/"+response.data);
          })
          .catch((error) => {
            console.error('Error uploading image', error);
            failure('Error uploading image');
          });
      },
  }
}
</script>