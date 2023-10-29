<template>
    <div class="login-box">
       <Nuxt/>
    </div>
</template>

<script>

export default {
   head() {
       return {
           title: 'Login',
           bodyAttrs: {
               class: 'hold-transition login-page'
           }
       }
   },
}
</script>