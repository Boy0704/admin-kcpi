<template>
    <div>
        <editor :api-key="apiTiny" :init="{
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
        }" v-model="internalValue" required />
    </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'

export default {
    components: {
        'editor': Editor
    },
    data() {
        return {
            apiTiny: this.$config.ApiKeyTiny,
            urlFile: this.$config.FileUrl,
            internalValue: this.model
        }
    },
    props: {
        model: String,
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
                    success(this.urlFile + "/img/" + response.data);
                })
                .catch((error) => {
                    console.error('Error uploading image', error);
                    failure('Error uploading image');
                });
        },

    },
    watch: {
        model(newValue) {
            this.internalValue = newValue
        },
        internalValue(newValue){
            this.$emit('input', newValue)
        }
    }

}
</script>