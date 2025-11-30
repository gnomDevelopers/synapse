<template>
  <div class="h-full bg-clr1 p-6" :class="{'overflow-hidden': isMaterialModalOpen || isTestResultModalOpen}">

    <div class="max-w-4xl mx-auto" :inert="isMaterialModalOpen || isTestResultModalOpen">
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

        <iconButton class="btn-main gap-2" @click="openAssigmentModal">
          <template #icon>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
          </template>
          <template #text>
            <p>Добавить задание</p>
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
            <materialItem :id="material.id" :title="material.title" :created-at="formatDate(material.createdAt)"/>
          </li>
        </ul>
      </section>

      <section class="mb-10">
        <h2 class="text-xl font-bold text-gray-800 mb-4 flex items-center gap-2 cursor-default select-none">
          <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 30 37" xmlns="http://www.w3.org/2000/svg"><path d="M21.5 4.83333H24.8333C25.7174 4.83333 26.5652 5.18452 27.1904 5.80964C27.8155 6.43477 28.1667 7.28261 28.1667 8.16667V31.5C28.1667 32.3841 27.8155 33.2319 27.1904 33.857C26.5652 34.4821 25.7174 34.8333 24.8333 34.8333H4.83333C3.94928 34.8333 3.10143 34.4821 2.47631 33.857C1.85119 33.2319 1.5 32.3841 1.5 31.5V8.16667C1.5 7.28261 1.85119 6.43477 2.47631 5.80964C3.10143 5.18452 3.94928 4.83333 4.83333 4.83333H8.16667M9.83333 1.5H19.8333C20.7538 1.5 21.5 2.24619 21.5 3.16667V6.5C21.5 7.42047 20.7538 8.16667 19.8333 8.16667H9.83333C8.91286 8.16667 8.16667 7.42047 8.16667 6.5V3.16667C8.16667 2.24619 8.91286 1.5 9.83333 1.5Z" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></svg>
          Задания
        </h2>
  
        <div v-if="assignments.length === 0" class="text-gray-400 italic cursor-default select-none">
          Задания пока не добавлены
        </div>
  
        <ul class="space-y-3">
          <li v-for="assignment in assignments" :key="assignment.id">
            <assigmentItem :id="assignment.id" :title="assignment.title" :created-at="formatDate(assignment.createdAt)" @show-results=""/>
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
            <testItem 
              :id="test.id" 
              :title="test.title" 
              :deadline="formatDate(test.deadline)" 
              :passing-score="test.passingScore" 
              :max-score="test.maxScore" 
              @show-results="openTestResultModal(test.id)"
            />
          </li>
        </ul>
      </section>
    </div>
  
    <addMaterialMW v-if="isMaterialModalOpen" @close="closeMaterialModal" @append="(newMaterial: any) => materials.push(newMaterial)"/>

    <addAssigmentMW v-if="isAssigmentModalOpen" @close="closeAssigmentModal" @append="(newAssigment: any) => assignments.push(newAssigment)"/>

    <showTestResultsMW v-if="isTestResultModalOpen" :discipline-id="currentDiscipline.id" @close="closeTestResultModal"/>
  </div>
</template>

<script lang="ts">
import type { IDiscipline, IMaterial, ITest, IAssigment } from '@/utils/types';
import { defineComponent } from 'vue';

import materialItem from './features/materialItem.vue';
import testItem from './features/testItem.vue';
import disciplineLogo from '@/widgets/PhotoTemplates/disciplineLogo.vue';
import addMaterialMW from './widgets/addMaterialMW.vue';
import showTestResultsMW from './widgets/showTestResultsMW.vue';
import assigmentItem from './features/assigmentItem.vue';
import addAssigmentMW from './widgets/addAssigmentMW.vue';

export default defineComponent({
  name: 'DisciplineDetails',
  components: {
    materialItem,
    testItem,
    disciplineLogo,
    addMaterialMW,
    showTestResultsMW,
    assigmentItem,
    addAssigmentMW,
  },
  data() {
    return {
      disciplineId: '' as string,
      currentDiscipline: {} as IDiscipline,
      
      // Данные
      materials: [] as IMaterial[],
      assignments: [] as IAssigment[],
      tests: [] as ITest[],
      selectedTestId: null as number | null,

      // Состояние UI
      isLoading: false,
      isMaterialModalOpen: false,
      isTestResultModalOpen: false,
      isAssigmentModalOpen: false,
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
          this.assignments = [
            { id: 1, title: 'Практическая работа 1', fileUrl: '#', createdAt: '2023-10-01' },
            { id: 2, title: 'Практическая работа 2', fileUrl: '#', createdAt: '2023-10-05' },
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

    // --- Модальное окно создания учебных материалов ---
    openMaterialModal() {
      this.isTestResultModalOpen = false;
      this.isAssigmentModalOpen = false;
      this.isMaterialModalOpen = true;
    },
    closeMaterialModal() {
      this.isMaterialModalOpen = false;
    },

    // --- Модальное окно создания задания ---
    openAssigmentModal() {
      this.isMaterialModalOpen = false;
      this.isTestResultModalOpen = false;
      this.isAssigmentModalOpen = true;
    },
    closeAssigmentModal() {
      this.isAssigmentModalOpen = false;
    },

    // --- Модальное окно результатов теста ---
    openTestResultModal(testID: number) {
      this.selectedTestId = testID;
      this.isMaterialModalOpen = false;
      this.isAssigmentModalOpen = false;
      this.isTestResultModalOpen = true;
    },
    closeTestResultModal() {
      this.isTestResultModalOpen = false;
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