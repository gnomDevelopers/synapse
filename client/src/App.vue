<template>
  <StatusWindow />
  <RouterView></RouterView>
</template>

<script lang="ts">
import StatusWindow from './widgets/StatusWindow/components/statusWindow.vue';

export default{
  components: {
    StatusWindow
  },
};
</script>