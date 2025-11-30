<template>
  <div class="min-h-screen bg-clr0 p-6 text-gray-800">
    <div class="max-w-4xl mx-auto">

      <div class="mb-8 flex justify-between">
        <iconButton class="btn-light gap-2" @click="$router.back()">
          <template #icon>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
          </template>
          <template #text>
            <p>Назад</p>
          </template>
        </iconButton>

        <iconButton v-if="questions.length > 0" class="btn-main gap-2" @click="saveTest">
          <template #icon>
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/>
              <polyline points="17 21 17 13 7 13 7 21"/>
              <polyline points="7 3 7 8 15 8"/>
            </svg>
          </template>
          <template #text>
            <p>Сохранить тест</p>
          </template>
        </iconButton>
      </div>

      <div v-if="questions.length > 0" class="space-y-6 mb-12">
        <h2 class="text-2xl font-bold text-gray-900 border-b border-gray-200 pb-2 cursor-default select-none">
          Список вопросов ({{ questions.length }})
        </h2>

        <div v-for="(question, qIndex) in questions" :key="question.id"
          class="bg-white border border-gray-200 rounded-lg shadow-sm p-6 relative">

          <button @click="deleteQuestion(question.id)"
            class="absolute top-4 right-4 text-gray-400 hover:text-gray-900 transition-colors cursor-pointer select-none" title="Удалить вопрос">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
          </button>

          <div class="flex justify-between items-start mb-4 pr-8">
            <h3 class="text-lg font-semibold text-gray-900">{{ qIndex + 1 }}. {{ question.text }}</h3>
            <span
              class="bg-gray-100 text-gray-700 text-xs font-bold px-2 py-1 rounded border border-gray-300 whitespace-nowrap">
              {{ question.points }} баллов
            </span>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <h4 class="text-sm font-bold text-gray-900 uppercase tracking-wide mb-2 flex items-center gap-2 cursor-default select-none">
                <span class="w-2 h-2 rounded-full bg-gray-600"></span>
                Правильные ответы
              </h4>
              <ul class="list-disc list-inside text-sm text-gray-600 space-y-1">
                <li v-for="(ans, i) in question.correctAnswers" :key="'c' + i">{{ ans }}</li>
              </ul>
            </div>

            <div>
              <h4 class="text-sm font-bold text-gray-900 uppercase tracking-wide mb-2 flex items-center gap-2 cursor-default select-none">
                <span class="w-2 h-2 rounded-full border border-gray-400"></span>
                Неправильные ответы
              </h4>
              <ul class="list-disc list-inside text-sm text-gray-500 space-y-1">
                <li v-for="(ans, i) in question.incorrectAnswers" :key="'inc' + i">{{ ans }}</li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-10 mb-12 bg-gray-100 rounded border border-dashed border-gray-300">
        <p class="text-gray-500 cursor-default select-none">Вопросов пока нет. Создайте первый вопрос ниже.</p>
      </div>

      <div class="bg-white border border-gray-300 rounded-lg shadow-md overflow-hidden">
        <div class="bg-gray-50 px-6 py-4 border-b border-gray-200">
          <h2 class="text-xl font-bold text-gray-800 cursor-default select-none">Создание нового вопроса</h2>
        </div>

        <div class="p-6 space-y-6">

          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="md:col-span-3">
              <label class="block text-sm font-medium text-gray-700 mb-1 cursor-pointer select-none" for="inp-question">Текст вопроса:</label>
              <textarea id="inp-question" v-model="newQuestion.text" rows="2"
                class="w-full px-3 py-2 bg-white border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-gray-400 focus:border-gray-400 transition-shadow"
                placeholder="Введите формулировку вопроса..."></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1 cursor-pointer select-none" for="inp-score">Баллы:</label>
              <input id="inp-score" v-model.number="newQuestion.points" type="number" min="0"
                class="w-full px-3 py-2 bg-white border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-gray-400" />
            </div>
          </div>

          <hr class="border-gray-400" />

          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">

            <div class="bg-gray-50 p-4 rounded border border-gray-200">
              <div class="flex justify-between items-center mb-3">
                <label class="block text-sm font-bold text-gray-800">Правильные ответы</label>
                <button @click="addAnswerField('correct')"
                  class="text-xs flex items-center gap-1 text-gray-600 hover:text-gray-900 underline decoration-dotted cursor-pointer select-none">
                  <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
                  Добавить
                </button>
              </div>

              <div class="space-y-2">
                <div v-for="(item, index) in newQuestion.correctAnswers" :key="'new-c-' + index" class="flex gap-2">
                  <input v-model="item.value" type="text" placeholder="Верный ответ"
                    class="flex-1 px-3 py-1.5 text-sm bg-white border border-gray-300 rounded focus:outline-none focus:border-gray-500" />
                  <button @click="removeAnswerField('correct', index)"
                    class="text-gray-400 hover:text-red-800 transition-colors p-1"
                    :disabled="newQuestion.correctAnswers.length === 1"
                    :class="{ 'opacity-0 pointer-events-none': newQuestion.correctAnswers.length === 1 }">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                  </button>
                </div>
              </div>
            </div>

            <div class="bg-gray-50 p-4 rounded border border-gray-200">
              <div class="flex justify-between items-center mb-3">
                <label class="block text-sm font-bold text-gray-500">Неправильные ответы</label>
                <button @click="addAnswerField('incorrect')"
                  class="text-xs flex items-center gap-1 text-gray-500 hover:text-gray-800 underline decoration-dotted cursor-pointer select-none">
                  <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
                  Добавить
                </button>
              </div>

              <div class="space-y-2">
                <div v-for="(item, index) in newQuestion.incorrectAnswers" :key="'new-inc-' + index" class="flex gap-2">
                  <input v-model="item.value" type="text" placeholder="Ложный ответ"
                    class="flex-1 px-3 py-1.5 text-sm bg-white border border-gray-300 rounded focus:outline-none focus:border-gray-500" />
                  <button @click="removeAnswerField('incorrect', index)"
                    class="text-gray-400 hover:text-red-800 transition-colors p-1"
                    :disabled="newQuestion.incorrectAnswers.length === 1"
                    :class="{ 'opacity-0 pointer-events-none': newQuestion.incorrectAnswers.length === 1 }">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                  </button>
                </div>
              </div>
            </div>

          </div>

          <div class="pt-4 flex justify-end">
            <iconButton class="btn-main px-8 py-3 text-lg" :class="{'btn-disabled': !isValid}" @click="addQuestion">
              <template #text>
                <p>Добавить вопрос</p>
              </template>
            </iconButton>
          </div>

        </div>
      </div>

    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

// Интерфейс для готового вопроса
interface Question {
  id: number;
  text: string;
  points: number;
  correctAnswers: string[];
  incorrectAnswers: string[];
}

// Интерфейс для инпута ответа (нужен объект для стабильной работы v-for при редактировании)
interface AnswerInput {
  value: string;
}

export default defineComponent({
  name: 'TestConstructor',
  data() {
    return {
      // Список уже созданных вопросов
      questions: [] as Question[],

      // Модель формы создания нового вопроса
      newQuestion: {
        text: '',
        points: 1,
        // Инициализируем массивы с одним пустым полем для удобства
        correctAnswers: [{ value: '' }] as AnswerInput[],
        incorrectAnswers: [{ value: '' }] as AnswerInput[]
      }
    };
  },
  computed: {
    // Простая валидация: текст вопроса не пустой, и есть хотя бы по одному заполненному ответу в каждой категории
    isValid(): boolean {
      const hasText = this.newQuestion.text.trim().length > 0;
      const hasCorrect = this.newQuestion.correctAnswers.some(a => a.value.trim().length > 0);
      const hasIncorrect = this.newQuestion.incorrectAnswers.some(a => a.value.trim().length > 0);
      return hasText && hasCorrect && hasIncorrect;
    }
  },
  mounted() {
    this.fetchQuestions();
  },
  methods: {
    // --- API Методы ---
    fetchQuestions() {
      // TODO: API GET запрос для получения списка существующих вопросов теста
      // axios.get('/api/tests/1/questions').then(...)

      // Имитация данных
      this.questions = [];
    },

    // --- Логика формы ---

    // Добавление инпута для ответа
    addAnswerField(type: 'correct' | 'incorrect') {
      if (type === 'correct') {
        this.newQuestion.correctAnswers.push({ value: '' });
      } else {
        this.newQuestion.incorrectAnswers.push({ value: '' });
      }
    },

    // Удаление инпута для ответа
    removeAnswerField(type: 'correct' | 'incorrect', index: number) {
      if (type === 'correct') {
        this.newQuestion.correctAnswers.splice(index, 1);
      } else {
        this.newQuestion.incorrectAnswers.splice(index, 1);
      }
    },

    // Создание вопроса
    async addQuestion() {
      if (!this.isValid) return;

      // Формируем payload, очищая пустые строки
      const payload = {
        text: this.newQuestion.text,
        points: this.newQuestion.points,
        correctAnswers: this.newQuestion.correctAnswers
          .map(a => a.value.trim())
          .filter(v => v.length > 0),
        incorrectAnswers: this.newQuestion.incorrectAnswers
          .map(a => a.value.trim())
          .filter(v => v.length > 0)
      };

      try {
        // TODO: API POST запрос для сохранения вопроса
        // const response = await axios.post('/api/questions', payload);
        // const savedQuestion = response.data;

        // Имитация успешного ответа сервера с присвоением ID
        const savedQuestion: Question = {
          id: Date.now(), // временный ID
          ...payload
        };

        // Обновляем локальный список
        this.questions.push(savedQuestion);

        // Сброс формы
        this.resetForm();

      } catch (error) {
        console.error('Ошибка при создании вопроса', error);
      }
    },

    // Удаление вопроса из списка
    async deleteQuestion(id: number) {
      if (!confirm('Вы уверены, что хотите удалить этот вопрос?')) return;

      try {
        // TODO: API DELETE запрос
        // await axios.delete(`/api/questions/${id}`);

        this.questions = this.questions.filter(q => q.id !== id);
      } catch (error) {
        console.error('Ошибка удаления', error);
      }
    },

    async saveTest() {
      // TODO: API SAVE TEST запрос
    },

    resetForm() {
      this.newQuestion = {
        text: '',
        points: 1,
        correctAnswers: [{ value: '' }],
        incorrectAnswers: [{ value: '' }]
      };
    }
  }
});
</script>