<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="absolute inset-0 bg-gray-900/20 backdrop-blur-sm transition-opacity" @click="closeTestResultModal"></div>

    <div class="relative bg-clr1 rounded-xl shadow-2xl w-full max-w-lg mx-4 p-6 z-10 transform transition-all">
      
      <h3 class="text-xl font-medium text-gray-800 mb-4 select-none cursor-default">Результаты теста</h3>
      
      <div v-if="isLoading" class="flex justify-center items-center h-48">
        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-indigo-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span class="text-gray-600">Загрузка результатов...</span>
      </div>

      <div v-else class="flex flex-col gap-3">
        
        <div class="flex justify-between items-center text-sm font-medium text-gray-500 border-b pb-2">
          <span>Студент ({{ results.length }})</span>
          <span>Баллы / {{ maxScore }}</span>
        </div>

        <div class="max-h-96 overflow-y-auto pr-2 custom-scrollbar">
          
          <div 
            v-for="result in results" 
            :key="result.userId" 
            class="flex items-center justify-between p-2 rounded-lg transition duration-150 ease-in-out hover:bg-gray-50"
          >
            <div class="flex items-center space-x-3 min-w-0">
              <userAvatar :user-id="result.userId" :size="'md'"/>
              
              <span class="text-sm font-medium text-gray-800 truncate" :title="result.fio">
                {{ result.fio }}
              </span>
            </div>
            
            <div class="flex-shrink-0 text-right">
              <span class="text-base font-semibold" :class="{ 
                'text-green-600': result.score >= maxScore * 0.8,
                'text-yellow-600': result.score < maxScore * 0.8 && result.score >= maxScore * 0.5,
                'text-red-600': result.score < maxScore * 0.5
              }">
                {{ result.score }}
              </span>
              <span class="text-sm text-gray-500">
                / {{ maxScore }}
              </span>
            </div>
          </div>
          
        </div>
      </div>
      
      <div class="mt-6 flex justify-end">
        <iconButton class="btn-main gap-2" @click="closeTestResultModal">
          <template #text>
            <p>Закрыть</p>
          </template>
        </iconButton>
      </div>

    </div>
  </div>
</template>

<script lang="ts">
import userAvatar from '@/widgets/PhotoTemplates/userAvatar.vue';

// Мок данных для демонстрации
const MOCK_TEST_RESULTS = [
  { userId: 1, fio: 'Иванов Иван Иванович', score: 95 },
  { userId: 2, fio: 'Петрова Анна Сергеевна', score: 100 },
  { userId: 3, fio: 'Сидоров Олег Владимирович', score: 68 },
  { userId: 4, fio: 'Кузнецова Мария Николаевна', score: 81 },
  { userId: 5, fio: 'Васильев Алексей Павлович', score: 45 },
  { userId: 6, fio: 'Смирнова Елена Дмитриевна', score: 77 },
  { userId: 7, fio: 'Михайлов Георгий Андреевич', score: 92 },
  { userId: 8, fio: 'Федорова Ксения Борисовна', score: 55 },
  { userId: 9, fio: 'Абрамов Виктор Геннадьевич', score: 71 },
  { userId: 10, fio: 'Зайцева Наталья Евгеньевна', score: 88 },
];

const MAX_SCORE = 100; // Максимально возможный балл за тест

export default {
  name: 'TestResultsModal',
  components: {
    userAvatar,
  },
  emits: ['close'],
  props: {
    disciplineId: {
      type: Number,
      default: 0,
    }
  },

  // Данные
  data() {
    return {
      // Состояние UI
      isLoading: true, // Изначально true для имитации загрузки

      // Результаты теста
      results: [] as { userId: number, fio: string, score: number }[], 
      
      // Максимальный балл (можно получать с API вместе с результатами)
      maxScore: MAX_SCORE,
    }
  },

  // Жизненный цикл
  mounted() {
    // Вызов загрузки данных при монтировании
    this.loadTestResults();
  },

  // Методы
  methods: {
    closeTestResultModal() {
      this.$emit('close');
    },

    async loadTestResults() {
      this.isLoading = true;

      // TODO API GET (Закомментированный код для реального API)
      /*
      try {
        const response = await fetch(`/api/test-results?disciplineId=${this.disciplineId}`);
        if (!response.ok) {
          throw new Error('Ошибка загрузки результатов');
        }
        const data = await response.json();
        this.results = data.results;
        this.maxScore = data.maxScore;
      } catch (error) {
        console.error('Ошибка загрузки результатов теста:', error);
        // Можно показать сообщение об ошибке пользователю
        this.results = [];
      } finally {
        this.isLoading = false;
      }
      */

      // Имитация задержки API (например, 500 мс)
      await new Promise(resolve => setTimeout(resolve, 500)); 

      // Присвоение моковых данных
      this.results = MOCK_TEST_RESULTS.sort((a, b) => b.score - a.score); // Сортировка по убыванию баллов
      this.isLoading = false;
    },
  }
}
</script>

<style scoped>
/* Кастомный стиль для скроллбара, чтобы он выглядел лучше (опционально) */
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #cbd5e1; /* gray-400 */
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #94a3b8; /* gray-500 */
}
</style>