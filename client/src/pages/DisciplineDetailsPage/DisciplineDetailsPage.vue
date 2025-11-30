<template>
  <div class="min-h-screen bg-clr1 p-6">

    <div class="max-w-4xl mx-auto">
      <div class="flex flex-col sm:flex-row gap-4 mb-8">
        <iconButton class="btn-light gap-2" @click="$router.push({name: 'DisciplinesPage'})">
          <template #icon>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
          </template>
          <template #text>
            <p>Назад</p>
          </template>
        </iconButton>

        <iconButton class="btn-main gap-2" @click="openMaterialModal">
          <template #icon>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
          </template>
          <template #text>
            <p>Добавить учебный материал</p>
          </template>
        </iconButton>

        <iconButton class="btn-main gap-2" @click="goToCreateTest">
          <template #icon>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path></svg>
          </template>
          <template #text>
            <p>Добавить тест</p>
          </template>
        </iconButton>
      </div>

      <section class="flex flex-row gap-2 items-stretch mb-10">
        <div class="w-full md:w-56 h-48 md:h-auto shrink-0 rounded-lg overflow-hidden">
          <disciplineLogo :discipline-id="currentDiscipline.id"/>
        </div>

        <div class="flex flex-col justify-start px-6 w-full cursor-default">
          <div class="mb-4">
            <h3 class="text-2xl font-bold text-gray-800 mb-2">{{ currentDiscipline.title }}</h3>
            <p class="text-gray-600 text-base line-clamp-3">
              {{ currentDiscipline.description }}
            </p>
          </div>
        </div>
      </section>

      <section class="mb-10">
        <h2 class="text-xl font-bold text-gray-800 mb-4 flex items-center gap-2 cursor-default select-none">
          <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"></path></svg>
          Учебные материалы
        </h2>
  
        <div v-if="materials.length === 0" class="text-gray-400 italic cursor-default select-none">
          Материалы пока не добавлены
        </div>
  
        <ul class="space-y-3">
          <li v-for="material in materials" :key="material.id">
            <materialItem :id="material.id" :title="material.title" :created-at="material.createdAt"/>
          </li>
        </ul>
      </section>
  
      <section>
        <h2 class="text-xl font-bold text-gray-800 mb-4 flex items-center gap-2 cursor-default select-none">
          <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"></path></svg>
          Тесты
        </h2>
  
        <div v-if="tests.length === 0" class="text-gray-400 italic cursor-default select-none">
          Тесты пока не созданы
        </div>
  
        <ul class="space-y-3">
          <li v-for="test in tests" :key="test.id">
            <testItem :id="test.id" :title="test.title" :deadline="test.deadline" :passing-score="test.passingScore" :max-score="test.maxScore"/>
          </li>
        </ul>
      </section>
    </div>
  

    <div v-if="isMaterialModalOpen" class="fixed inset-0 z-50 flex items-center justify-center">
      <div 
        class="absolute inset-0 bg-blur/20 bg-opacity-40 backdrop-blur-xs transition-opacity"
        @click="closeMaterialModal"
      ></div>

      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-md mx-4 p-6 z-10">
        <h3 class="text-lg font-bold text-gray-800 mb-4 cursor-default select-none">Добавить материал</h3>
        
        <div class="flex flex-col gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1 cursor-pointer select-none" for="inp-title">Название материала</label>
            <input 
              id="inp-title"
              v-model="newMaterial.title"
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
            <iconButton class="btn-light gap-2" @click="closeMaterialModal">
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

  </div>
</template>

<script lang="ts">
import type { IDiscipline, IMaterial, ITest } from '@/utils/types';
import { defineComponent } from 'vue';

import materialItem from './features/materialItem.vue';
import testItem from './features/testItem.vue';
import disciplineLogo from '@/widgets/PhotoTemplates/disciplineLogo.vue';

export default defineComponent({
  name: 'DisciplineDetails',
  components: {
    materialItem,
    testItem,
    disciplineLogo,
  },
  data() {
    return {
      disciplineId: '' as string,
      currentDiscipline: {} as IDiscipline,
      
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
          this.currentDiscipline = {
            id: 1,
            title: 'Высшая математика',
            description: 'Курс линейной алгебры и аналитической геометрии.',
          } as IDiscipline;
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