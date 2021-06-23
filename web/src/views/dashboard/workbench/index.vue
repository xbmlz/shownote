<template>

</template>
<script lang="ts">
  import { defineComponent, ref } from 'vue';


  export default defineComponent({
    setup() {
      const loading = ref(true);

      setTimeout(() => {
        loading.value = false;
      }, 1500);

      return {
        loading,
      };
    },
  });
</script>
