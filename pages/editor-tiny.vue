<template>
  <div id="app">
    <editor
      api-key="4mo39ri6dgnfnyfqwhr6nhicdjgg3nckwd3ruoyr8sa3d5z7"
      :init="{
        height: 500,
        menubar: true,
        images_upload_url: urlBase+'/upload-img',
        automatic_uploads: true,
        images_reuse_filename: true,
        images_upload_handler: handleImageUpload,
        plugins: [
          'advlist autolink lists link image charmap print preview anchor',
          'searchreplace visualblocks code fullscreen',
          'insertdatetime media table paste code help wordcount'
        ],
        toolbar:
          'undo redo | formatselect | bold italic backcolor | \
          alignleft aligncenter alignright alignjustify | \
          bullist numlist outdent indent | removeformat | help'
      }"
    />
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'

const handleImageUpload = async (blobInfo, progress, failure) => {
  const formData = new FormData();
  formData.append("file_img", blobInfo.blob(), blobInfo.filename());

  try {
    const response = await this.$axios
					.$post('/upload-img',
      formData,
      {
        headers: { "Content-Type": "multipart/form-data", "Authorization": `Bearer ${this.$auth.user.token}` },
        onUploadProgress: (progressEvent) => {
          const percentCompleted = Math.round(
            (progressEvent.loaded * 100) / progressEvent.total
          );
          if (progress && typeof progress === "function") {
            progress(percentCompleted);
          }
        },
      }
    );

    if (response.status === 403) {
      throw new Error("HTTP Error: " + response.status);
    }

    if (response.status < 200 || response.status >= 300) {
      throw new Error("HTTP Error: " + response.status);
    }

    const json = response.data;

    if (!json || typeof json.location !== "string") {
      throw new Error("Invalid JSON: " + JSON.stringify(json));
    } else response;
    {
      return json.location;
    }
  } catch (error) {
    if (failure && typeof failure === "function") {
      failure(error.message);
    }
  }
};

export default {
  name: 'app',
  components: {
    'editor': Editor
  },
  data() {
		return {
			urlFile: this.$config.FileUrl,
			urlBase: this.$config.ApiUrl,
      handleImageUpload: null
		}
	},
}
</script>