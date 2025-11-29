<template>
  <div class="min-h-screen bg-gray-50 p-6 font-sans">
    
    <div class="flex flex-col sm:flex-row gap-4 mb-8">
      <iconButton class="btn-main" @click="$router.push({name: 'DisciplinesPage'})">
        <template #icon>
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
        </template>
        <template #text>
          <p>Назад</p>
        </template>
      </iconButton>

      <iconButton class="btn-main" @click="openMaterialModal">
        <template #icon>
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
        </template>
        <template #text>
          <p>Добавить учебный материал</p>
        </template>
      </iconButton>

      <iconButton class="btn-main" @click="goToCreateTest">
        <template #icon>
           <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path></svg>
        </template>
        <template #text>
          <p>Добавить тест</p>
        </template>
      </iconButton>
    </div>

    <section class="mb-10">
      <h2 class="text-xl font-bold text-gray-800 mb-4 flex items-center gap-2">
        <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"></path></svg>
        Учебные материалы
      </h2>

      <div v-if="materials.length === 0" class="text-gray-400 italic">
        Материалы пока не добавлены
      </div>

      <ul class="space-y-3">
        <li 
          v-for="material in materials" 
          :key="material.id"
          class="bg-white border border-gray-200 rounded-lg p-4 flex items-center gap-4 hover:shadow-md transition-shadow cursor-pointer"
        >
          <div class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center flex-shrink-0 text-gray-500">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path></svg>
          </div>
          
          <div class="flex-grow">
            <h3 class="text-gray-800 font-medium">{{ material.title }}</h3>
            <span class="text-xs text-gray-500">Добавлено: {{ formatDate(material.createdAt) }}</span>
          </div>

          <button class="text-gray-400 hover:text-gray-600 p-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path></svg>
          </button>
        </li>
      </ul>
    </section>

    <section>
      <h2 class="text-xl font-bold text-gray-800 mb-4 flex items-center gap-2">
        <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"></path></svg>
        Тесты
      </h2>

      <div v-if="tests.length === 0" class="text-gray-400 italic">
        Тесты пока не созданы
      </div>

      <ul class="space-y-3">
        <li 
          v-for="test in tests" 
          :key="test.id"
          class="bg-white border border-gray-200 rounded-lg p-4 flex flex-col sm:flex-row sm:items-center gap-4 hover:border-gray-400 transition-colors"
        >
          <div class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center flex-shrink-0 text-gray-500">
             <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
          </div>

          <div class="flex-grow">
            <h3 class="text-gray-800 font-medium">{{ test.title }}</h3>
            <div class="text-sm text-gray-500 mt-1 flex flex-wrap gap-x-4">
              <span>Дедлайн: <span class="text-gray-700 font-semibold">{{ formatDate(test.deadline) }}</span></span>
            </div>
          </div>

          <div class="flex-shrink-0 flex items-center gap-2 bg-gray-50 px-3 py-1 rounded border border-gray-200">
            <span class="text-sm text-gray-500 font-bold">Проходной:</span>
            <span class="text-gray-800 font-medium">{{ test.passingScore }}/{{ test.maxScore }}</span>
          </div>
        </li>
      </ul>
    </section>

    <div v-if="isMaterialModalOpen" class="fixed inset-0 z-50 flex items-center justify-center">
      <div 
        class="absolute inset-0 bg-blur/20 bg-opacity-40 backdrop-blur-xs transition-opacity"
        @click="closeMaterialModal"
      ></div>

      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-md mx-4 p-6 z-10">
        <h3 class="text-lg font-bold text-gray-800 mb-4">Добавить материал</h3>
        
        <form @submit.prevent="createMaterial" class="flex flex-col gap-4">
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Название материала</label>
            <input 
              v-model="newMaterial.title"
              type="text" 
              required
              class="w-full border border-gray-300 rounded p-2 focus:ring-2 focus:ring-gray-400 focus:outline-none"
              placeholder="Введите название"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Файл</label>
            <input 
              type="file" 
              required
              @change="handleFileChange"
              class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:text-sm file:font-semibold file:bg-gray-100 file:text-gray-700 hover:file:bg-gray-200 cursor-pointer"
            />
          </div>

          <div class="flex justify-end gap-3 mt-4">
            <button 
              type="button" 
              @click="closeMaterialModal" 
              class="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
            >
              Отмена
            </button>
            <button 
              type="submit" 
              class="px-6 py-2 bg-gray-700 text-white rounded hover:bg-gray-800 transition-colors"
              :disabled="isSubmitting"
            >
              {{ isSubmitting ? 'Загрузка...' : 'Добавить' }}
            </button>
          </div>
        </form>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import type { IMaterial, ITest } from '@/utils/types';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'DisciplineDetails',
  data() {
    return {
      disciplineId: '' as string,
      
      // Данные
      materials: [] as IMaterial[],
      tests: [] as ITest[],

      // Состояние UI
      isLoading: false,
      isMaterialModalOpen: false,
      isSubmitting: false,

      // Форма нового материала
      newMaterial: {
        title: '',
        file: null as File | null
      }
    };
  },
  mounted() {
    // Получаем ID из URL
    this.disciplineId = this.$route.params.id as string;
    
    this.fetchData();
  },
  methods: {
    async fetchData() {
      this.isLoading = true;
      try {
        console.log(`Загрузка данных для дисциплины ID: ${this.disciplineId}`);
        // API: GET /disciplines/:id/materials
        // API: GET /disciplines/:id/tests
        
        // --- Mock Data ---
        setTimeout(() => {
          this.materials = [
            { id: 1, title: 'Лекция 1. Введение', fileUrl: '#', createdAt: '2023-10-01' },
            { id: 2, title: 'Методичка к лабораторной', fileUrl: '#', createdAt: '2023-10-05' },
          ];
          this.tests = [
            { id: 1, title: 'Контрольный тест №1', deadline: '2023-11-20', passingScore: 12, maxScore: 20 },
            { id: 2, title: 'Экзамен', deadline: '2023-12-25', passingScore: 50, maxScore: 100 },
          ];
          this.isLoading = false;
        }, 500);
        // -----------------

      } catch (error) {
        console.error(error);
        this.isLoading = false;
      }
    },

    // --- Навигация ---
    goToCreateTest() {
      // Редирект на конструктор тестов (можно передать ID дисциплины в query параметрах, если нужно)
      this.$router.push('/test/create');
    },

    // --- Модальное окно материалов ---
    openMaterialModal() {
      this.isMaterialModalOpen = true;
    },
    
    closeMaterialModal() {
      this.isMaterialModalOpen = false;
      this.resetMaterialForm();
    },

    resetMaterialForm() {
      this.newMaterial.title = '';
      this.newMaterial.file = null;
    },

    handleFileChange(event: Event) {
      const input = event.target as HTMLInputElement;
      if (input.files && input.files[0]) {
        this.newMaterial.file = input.files[0];
      }
    },

    async createMaterial() {
      if (!this.newMaterial.title || !this.newMaterial.file) return;

      this.isSubmitting = true;
      try {
        const formData = new FormData();
        formData.append('disciplineId', this.disciplineId);
        formData.append('title', this.newMaterial.title);
        formData.append('file', this.newMaterial.file);

        // API: POST /materials (или похожий эндпоинт)
        // await api.post('/materials', formData);
        
        console.log('Материал создан:', this.newMaterial.title);
        
        // Обновляем список (в реальном приложении добавляем ответ сервера)
        this.materials.unshift({
          id: Date.now(),
          title: this.newMaterial.title,
          fileUrl: '#',
          createdAt: new Date().toISOString()
        });

        this.closeMaterialModal();
      } catch (error) {
        console.error('Ошибка при создании материала', error);
      } finally {
        this.isSubmitting = false;
      }
    },

    // --- Утилиты ---
    formatDate(dateString: string): string {
      const options: Intl.DateTimeFormatOptions = { 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric' 
      };
      return new Date(dateString).toLocaleDateString('ru-RU', options);
    }
  }
});
</script>
```