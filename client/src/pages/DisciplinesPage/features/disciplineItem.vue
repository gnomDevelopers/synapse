<template>
  <div class="flex flex-col md:flex-row bg-clr1 border border-border rounded-lg overflow-hidden shadow-sm hover:shadow-md transition-shadow duration-300 w-full max-w-4xl">
    <div class="w-full md:w-56 h-48 md:h-auto shrink-0">
      <img 
        v-if="discipline.imageUrl" 
        :src="discipline.imageUrl" 
        alt="Discipline cover" 
        class="w-full h-full object-cover"
      />
      <DisciplineLogo v-else :discipline-id="discipline.id"/>
    </div>

    <div class="flex flex-col justify-between p-6 w-full cursor-default">
      <div class="mb-4">
        <h3 class="text-xl font-bold text-gray-800 mb-2">{{ discipline.title }}</h3>
        <p class="text-gray-600 text-sm line-clamp-3">
          {{ discipline.description }}
        </p>
      </div>
      
      <div class="flex justify-end">
        <iconButton class="btn-light py-1" @click="$router.push({name: 'DisciplineDetailsPage', params: {id: discipline.id}})">
          <template #text>
            <p>Перейти</p>
          </template>
        </iconButton>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import type { IDiscipline } from '@/utils/types';
import { defineComponent, type PropType } from 'vue';
import DisciplineLogo from '@/widgets/PhotoTemplates/disciplineLogo.vue';

export default defineComponent({
  name: 'DisciplineItem',
  components: {
    DisciplineLogo,
  },
  props: {
    discipline: {
      type: Object as PropType<IDiscipline>,
      required: true
    }
  },
  emits: ['navigate'],
  methods: {
    onNavigate() {
      // Эмитим событие, чтобы родитель решил, куда перенаправлять
      this.$emit('navigate', this.discipline.id);
    }
  }
});
</script>