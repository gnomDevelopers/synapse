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

const ASSETS_PATH = '../../assets/images/disciplines-logo';

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
      try {
        return new URL(`${ASSETS_PATH}/discipline-logo-${this.logoIndex}.png`, import.meta.url).href;
      } catch (e) {
        console.error('Ошибка загрузки заставки:', e);
        return new URL(`${ASSETS_PATH}/discipline-logo-1.png`, import.meta.url).href;
      }
    },
  }
});
</script>