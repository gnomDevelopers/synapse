<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="absolute inset-0 bg-blur/20 bg-opacity-40 backdrop-blur-xs transition-opacity" @click="closeAssigmentModal"></div>

    <div class="relative bg-white rounded-lg shadow-xl w-full max-w-md mx-4 p-6 z-10">
      <h3 class="text-lg font-bold text-gray-800 mb-4 cursor-default select-none">Добавить задание</h3>
      
      <div class="flex flex-col gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1 cursor-pointer select-none" for="inp-title">Название задания</label>
          <input 
            id="inp-title"
            v-model="newAssigment.title"
            type="text" 
            required
            class="w-full border border-gray-300 rounded p-2 focus:ring-2 focus:ring-gray-400 focus:outline-none"
            placeholder="Введите название"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1 cursor-pointer select-none" for="inp-file">Файл</label>
          <input 
            id="inp-file"
            type="file" 
            required
            @change="handleFileChange"
            class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:text-sm file:font-semibold file:bg-gray-100 file:text-gray-700 hover:file:bg-gray-200 cursor-pointer"
          />
        </div>

        <div class="flex justify-end gap-3 mt-4">
          <iconButton class="btn-light gap-2" @click="closeAssigmentModal">
            <template #text>
              <p>Отмена</p>
            </template>
          </iconButton>

          <iconButton class="btn-main gap-2" :class="{'btn-disabled': isSubmitting}" @click="createMaterial">
            <template #text>
              <p>{{ isSubmitting ? 'Загрузка...' : 'Добавить' }}</p>
            </template>
          </iconButton>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  emits: ['close', 'append'],
  props: {
    disciplineId: {
      type: Number,
      default: 0,
    }
  },
  data() {
    return {
      // Состояние UI
      isSubmitting: false,

      // Форма нового материала
      newAssigment: {
        title: '',
        file: null as File | null
      }
    }
  },
  methods: {
    closeAssigmentModal() {
      this.resetAssigmentForm();
      this.$emit('close');
    },

    resetAssigmentForm() {
      this.newAssigment.title = '';
      this.newAssigment.file = null;
    },

    handleFileChange(event: Event) {
      const input = event.target as HTMLInputElement;
      if (input.files && input.files[0]) {
        this.newAssigment.file = input.files[0];
      }
    },

    async createMaterial() {
      if (!this.newAssigment.title || !this.newAssigment.file) return;

      this.isSubmitting = true;
      try {
        const formData = new FormData();
        formData.append('disciplineId', this.disciplineId.toString());
        formData.append('title', this.newAssigment.title);
        formData.append('file', this.newAssigment.file);

        // API: POST /materials (или похожий эндпоинт)
        // await api.post('/materials', formData);
        
        console.log('Материал создан:', this.newAssigment.title);
        
        // Обновляем список
        this.$emit('append', {
          id: Date.now(),
          title: this.newAssigment.title,
          fileUrl: '#',
          createdAt: new Date().toISOString()
        })

        this.closeAssigmentModal();
      } catch (error) {
        console.error('Ошибка при создании задания', error);
      } finally {
        this.isSubmitting = false;
      }
    },
  }
}
</script>