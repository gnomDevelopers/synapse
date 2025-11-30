<template>
  <div
    class="shrink-0 h-full"
  >
    <img
      :src="logoSrc"
      alt="Discipline Logo"
      class="w-full h-full object-cover"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'DisciplineLogo',
  props: {
    disciplineId: {
      type: Number,
      required: true,
      default: 0
    },
  },
  computed: {
    logoIndex(): number {
      return (Math.abs(this.disciplineId) % 4) + 1;
    },
    logoSrc(): string {
      return `/images/disciplines-logo/discipline-logo-${this.logoIndex}.png`;
    },
  }
});
</script>
