<template>
    <div class="input-group mb-3">
        <input :type="typeInput" :class="classFormControl" :placeholder="placeholder" v-model="internalValue" required>
        <div class="input-group-append">
            <span class="input-group-text" @click="handler"><i :class="classHandler"></i></span>
        </div>
    </div>
</template>

<script>

export default {
    data() {
        return {
            view: false,
            typeInput: "password",
            classHandler: "fas fa-eye-slash",
            internalValue: this.model
        }
    },
    props: {
        placeholder: {type: String},
        model: String,
        classFormControl: String
    },
    methods: {
        handler() {
            if( this.view == true ) {
                this.typeInput = "text"
                this.classHandler = "fas fa-eye"
                this.view = false
            } else {
                this.typeInput = "password"
                this.classHandler = "fas fa-eye-slash"
                this.view = true
            }
        }
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